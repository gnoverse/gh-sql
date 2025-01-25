// Package sync implements a system to fetch GitHub repositories concurrently.
package sync

import (
	"context"
	"log"
	"regexp"
	"strings"

	"github.com/gnolang/gh-sql/ent"
	"github.com/gnolang/gh-sql/pkg/sync/internal/synchub"
)

// Options repsent the options to run Sync. These match the flags provided on
// the command line.
type Options struct {
	DB    *ent.Client
	Token string

	DebugHTTP bool
}

// Sync performs the synchronisation of repositories to the database provided
// in the options.
func Sync(ctx context.Context, repositories []string, opts Options) error {
	h := synchub.New(ctx, synchub.Options(opts))

	// execute
	for _, repo := range repositories {
		parts := strings.SplitN(repo, "/", 2)
		if len(parts) != 2 {
			log.Printf("invalid repo syntax: %q (must be <owner>/<repo>)", repo)
		}

		if strings.IndexByte(parts[1], '*') != -1 {
			// contains wildcards
			reString := "^" + strings.ReplaceAll(regexp.QuoteMeta(parts[1]), `\*`, `.*`) + "$"
			re := regexp.MustCompile(reString)
			shouldFetch := func(r *ent.Repository) bool {
				return re.MatchString(r.Name)
			}
			fetchRepositories(ctx, h, parts[0], shouldFetch)
		} else {
			// no wildcards; fetch repo directly.
			synchub.FetchAsync(ctx, h, fetchRepository{owner: parts[0], repo: parts[1]})
		}
	}

	// wait for all goros to finish
	return h.Wait()
}
