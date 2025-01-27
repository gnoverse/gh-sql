// Command gh-sql is the main entrypoint for the gh-sql tool.
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gnoverse/gh-sql/ent"
	"github.com/gnoverse/gh-sql/ent/migrate"
	"github.com/gnoverse/gh-sql/pkg/model"
	"github.com/gnoverse/gh-sql/pkg/rest"
	"github.com/gnoverse/gh-sql/pkg/sync"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

func main() {
	if err := run(); err != nil {
		if !errors.Is(err, ff.ErrHelp) {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		os.Exit(1)
	}
}

// TODO: ideally this should go in the help for dsn, see https://github.com/peterbourgon/ff/issues/140
const longHelp = `gh-sql is a tool to synchronize data from the GitHub REST API into a
SQL database.

CONFIGURATION

All the options of the program can be passed through flags, environment variables,
or a configuration file (and the -config flag), in this order of precedence.

- Environment variables match the flag name, by prepending "GHSQL_" to the name
  of the flag and replacing hyphens and periods with an underscore. For instance,
  --debug.http is available as GHSQL_DEBUG_HTTP.
- The config file is a simple key/value association, using ff's "PlainParser".
  For more information:
  <https://pkg.go.dev/github.com/peterbourgon/ff/v4@v4.0.0-alpha.4#PlainParser>

DATABASE SUPPORT

gh-sql currently supports postgres and sqlite as SQL databases. The database is
selected with the --dsn flag, which accepts an URL with information for
connecting to the specified database.

- sqlite://<FILE>
	<FILE> may be an absolute or relative path.
- postgres://[<USER>[:<PASS>]@][<HOST>][/<DB_NAME>]
	a postgres connection string
		see: <https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING-URIS>
	postgres environment variables are also recognized
		see: <https://pkg.go.dev/github.com/jackc/pgx/v5@v5.7.2/pgconn#ParseConfig>

CREDITS

Contribute: <https://github.com/gnoverse/gh-sql>
Report an issue: <https://github.com/gnoverse/gh-sql/issues/new/choose>

Morgan Bazalgette <morgan@howl.moe>
Licensed under the Apache License 2.0`

const (
	sqliteSchema   = "sqlite://"
	postgresSchema = "postgres://"
)

func run() error {
	// Parse configuration
	var (
		globalFS = ff.NewFlagSet("gh-sql")
		dsn      = globalFS.StringLong("dsn", "sqlite://database.sqlite", "data source name URL, see long help")
		_        = globalFS.StringLong("config", "", "config file (optional)")
		entDebug = globalFS.BoolLong("debug.ent", "use ent's debug client (very verbose)")
		ctx      = &execContext{}
	)
	_ = dsn
	cmd := &ff.Command{
		Name:      "gh-sql",
		Usage:     "gh-sql SUBCOMMAND ...",
		ShortHelp: "a tool to locally store information about GitHub repository, query them and serve them.",
		LongHelp:  longHelp,
		Flags:     globalFS,
		Subcommands: []*ff.Command{
			newSyncCmd(globalFS, ctx),
			newServeCmd(globalFS, ctx),

			// gh-sql query       # simple readonly query interface/CLI to get info from the database, json/csv/msgpack output.
		},
		Exec: func(context.Context, []string) error { return ff.ErrHelp },
	}

	err := cmd.Parse(os.Args[1:],
		ff.WithEnvVarPrefix("GHSQL"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, ffhelp.Command(cmd))
		return err
	}

	var entDrv *entsql.Driver
	switch {
	case strings.HasPrefix(*dsn, sqliteSchema):
		fname := (*dsn)[len(sqliteSchema):]
		abs, err := filepath.Abs(fname)
		if err != nil {
			return fmt.Errorf("could not make file path absolute: %w", err)
		}
		ctx.sqlDB, err = sql.Open("sqlite", fmt.Sprintf("file://%s?_pragma=foreign_keys(1)&_txlock=exclusive", abs))
		if err != nil {
			return fmt.Errorf("failed opening connection to postgres: %v", err)
		}
		defer ctx.sqlDB.Close()
		// sync has mostly writes to the database, and modernc.org/sqlite
		// does not support concurrent writes. Set MaxOpenConns=1 to avoid errors.
		ctx.sqlDB.SetMaxOpenConns(1)

		entDrv = entsql.OpenDB(dialect.SQLite, ctx.sqlDB)
	case strings.HasPrefix(*dsn, postgresSchema):
		ctx.sqlDB, err = sql.Open("pgx", *dsn)
		if err != nil {
			return fmt.Errorf("failed opening connection to postgres: %v", err)
		}
		defer ctx.sqlDB.Close()

		entDrv = entsql.OpenDB(dialect.Postgres, ctx.sqlDB)
	default:
		return fmt.Errorf("unsupported dsn URL: %q", *dsn)
	}
	// TODO: test out postgres

	ctx.db = ent.NewClient(ent.Driver(entDrv))
	if *entDebug {
		ctx.db = ctx.db.Debug()
	}

	// Run the auto migration tool.
	if err := ctx.db.Schema.Create(context.Background(),
		// Allows ent to drop indexes if we remove constraints from the schema.
		migrate.WithDropIndex(true)); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	err = cmd.Run(context.Background())
	if errors.Is(err, ff.ErrHelp) {
		fmt.Fprintln(os.Stderr, ffhelp.Command(cmd))
		return err
	}
	return nil
}

type execContext struct {
	sqlDB *sql.DB
	db    *ent.Client
}

func newSyncCmd(fs *ff.FlagSet, ec *execContext) *ff.Command {
	var (
		fset = ff.NewFlagSet("gh-sql sync").SetParent(fs)

		token        = fset.StringLong("token", "", "github personal access token to use (heavily suggested)")
		createdAfter = fset.StringLong("created-after", "", "only inspect issues and PRs created after this date (YYYY-MM-DD)")

		debugHTTP        = fset.BoolLong("debug.http", "log http requests")
		debugJSONCatcher = fset.StringLong(
			"debug.jsoncatcher",
			filepath.Join(os.TempDir(), "jsoncatcher.log"),
			"use jsoncatcher to find incorrectly-parsed events. (enabled by default)")
	)
	return &ff.Command{
		Name:      "sync",
		Usage:     "gh-sql sync [FLAGS]... REPOS...",
		ShortHelp: "synchronize the current state with GitHub",
		LongHelp:  "Repositories must be specified with the syntax <owner>/<repo>.",
		Flags:     fset,
		Exec: func(ctx context.Context, args []string) error {
			if *debugJSONCatcher != "" {
				model.EventsMissedData = &onDemandFile{fileName: *debugJSONCatcher}
			}

			var ca time.Time
			if *createdAfter != "" {
				var err error
				ca, err = time.Parse(time.DateOnly, *createdAfter)
				if err != nil {
					return err
				}
			}

			return sync.Sync(ctx, args, sync.Options{
				DB:           ec.db,
				Token:        *token,
				CreatedAfter: ca,
				DebugHTTP:    *debugHTTP,
			})
		},
	}
}

type onDemandFile struct {
	*os.File
	err      error
	fileName string
}

func (w *onDemandFile) Write(p []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	if w.File == nil {
		w.File, w.err = os.OpenFile(w.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if w.err != nil {
			return 0, w.err
		}
	}
	return w.File.Write(p)
}

func newServeCmd(fs *ff.FlagSet, ec *execContext) *ff.Command {
	var (
		fset = ff.NewFlagSet("gh-sql serve").SetParent(fs)
		addr = fset.StringLong("addr", ":31415", "listening address for web server")
	)
	return &ff.Command{
		Name:      "serve",
		Usage:     "gh-sql serve [--addr :31345] [FLAGS...]",
		ShortHelp: "listen for requests on the given address, mimicking the GitHub REST API",
		Flags:     fset,
		Exec: func(ctx context.Context, args []string) error {
			if len(args) > 0 {
				return ff.ErrHelp
			}
			srv := rest.Handler(rest.Options{
				DB: ec.db,
			})
			log.Printf("Listening on %s", *addr)
			return http.ListenAndServe(*addr, srv)
		},
	}
}
