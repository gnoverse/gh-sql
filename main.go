// Command gh-sql is the main entrypoint for the gh-sql tool.
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gnolang/gh-sql/ent"
	"github.com/gnolang/gh-sql/ent/migrate"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"

	entsql "entgo.io/ent/dialect/sql"
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

func run() error {
	// Parse configuration
	var (
		globalFS = ff.NewFlagSet("gh-sql")
		entDebug = globalFS.BoolLong("debug.ent", "use ent's debug client (very verbose)")
		_        = globalFS.StringLong("config", "", "config file (optional)")
		ctx      = &execContext{}
	)
	cmd := &ff.Command{
		Name:      "gh-sql",
		Usage:     "gh-sql SUBCOMMAND ...",
		ShortHelp: "a tool to locally store information about GitHub repository, query them and serve them.",
		LongHelp:  "", // TODO
		Flags:     globalFS,
		Subcommands: []*ff.Command{
			newSyncCmd(globalFS, ctx),

			// gh-sql query       # simple readonly query interface/CLI to get info from the database, json/csv/msgpack output.
			// gh-sql serve       # runs API server similar to GitHub API.
		},
		Exec: func(context.Context, []string) error { return ff.ErrHelp },
	}

	err := cmd.Parse(os.Args[1:],
		ff.WithEnvVarPrefix("GHSQL_"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, ffhelp.Command(cmd))
		return err
	}

	abs, err := filepath.Abs("./database.sqlite")
	if err != nil {
		return err
	}

	// TODO: this does not work concurrently; consider what to do (mattn/sqlite3?)
	ctx.sqlDB, err = sql.Open("sqlite", fmt.Sprintf("file://%s?_pragma=foreign_keys(1)&_txlock=exclusive", abs))
	if err != nil {
		return fmt.Errorf("failed opening connection to postgres: %v", err)
	}
	defer ctx.sqlDB.Close()

	entDrv := entsql.OpenDB("sqlite3", ctx.sqlDB)
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
