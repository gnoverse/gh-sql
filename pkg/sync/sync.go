// Package sync implements a system to fetch GitHub repositories concurrently.
package sync

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gnoverse/gh-sql/ent"
	"github.com/gnoverse/gh-sql/pkg/model"
	"github.com/gnoverse/gh-sql/pkg/sync/internal/synchub"
)

// Options repsent the options to run Sync. These match the flags provided on
// the command line.
type Options struct {
	DB           *ent.Client
	Token        string
	CreatedAfter time.Time

	DebugHTTP bool
}

// Sync performs the synchronisation of repositories to the database provided
// in the options.
func Sync(ctx context.Context, repositories []string, opts Options) error {
	h := synchub.New(ctx, synchub.Options{
		Token:     opts.Token,
		DebugHTTP: opts.DebugHTTP,
	})

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
			fetchRepositories(ctx, h, opts, parts[0], shouldFetch)
		} else {
			// no wildcards; fetch repo directly.
			synchub.FetchAsync(ctx, h, fetchRepository{
				Options: opts,
				owner:   parts[0],
				repo:    parts[1],
			})
		}
	}

	// wait for all goros to finish
	return h.Wait()
}

// fetchRepository can be used to fetch a GitHub repository, knowing its
// owner and name
type fetchRepository struct {
	Options
	owner, repo string
}

var _ synchub.Resource[*ent.Repository] = fetchRepository{}

func (f fetchRepository) ID() string { return "/repos/" + f.owner + "/" + f.repo }

func (f fetchRepository) Fetch(ctx context.Context, h *synchub.Hub) (*ent.Repository, error) {
	var r struct {
		ent.Repository
		Owner model.SimpleUser `json:"owner"`
	}
	if err := synchub.Get(ctx, h, f.ID(), &r); err != nil {
		return nil, fmt.Errorf("fetchRepository%+v fetch: %w", f, err)
	}

	u, err := synchub.Fetch(ctx, h, fetchUser{
		Options:  f.Options,
		username: r.Owner.Login,
	})
	if err != nil {
		return nil, err
	}

	err = f.DB.Repository.Create().
		CopyRepository(&r.Repository).
		SetOwnerID(u.ID).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchRepository%+v save: %w", f, err)
	}

	h.Go(func() error {
		// Pulls are fetched after issues, so that pulls already have all of the
		// PRs they link to fetched into the database.
		fetchIssues(ctx, h, f.Options, r.Owner.Login, r.Name)
		fetchPulls(ctx, h, f.Options, r.Owner.Login, r.Name)
		// TODO: fetch review comments - we can fetch these at repo level and
		// concurrently with pulls, most likely.
		// /repos/{owner}/{repo}/pulls/comments
		return nil
	})

	return f.DB.Repository.Get(ctx, r.ID)
}

func fetchRepositories(ctx context.Context, h *synchub.Hub, opts Options, owner string, shouldFetch func(*ent.Repository) bool) {
	iter := synchub.GetIterate[*ent.Repository](
		ctx, h,
		fmt.Sprintf("/users/%s/repos?per_page=100", owner))
	for r := range iter.Values {
		if shouldFetch(r) {
			synchub.FetchAsync(ctx, h, fetchRepository{
				Options: opts,
				owner:   owner,
				repo:    r.Name,
			})
		}
	}
	if iter.Err() != nil {
		h.Warn(fmt.Errorf("fetchRepositories(%q): %w", owner, iter.Err()))
	}
}

// fetchUser can be used to retrive a user, knowing its username.
type fetchUser struct {
	Options
	username string
}

var _ synchub.Resource[*ent.User] = fetchUser{}

func (fu fetchUser) ID() string { return "/users/" + fu.username }

func (fu fetchUser) Fetch(ctx context.Context, h *synchub.Hub) (*ent.User, error) {
	var u ent.User
	if err := synchub.Get(ctx, h, fu.ID(), &u); err != nil {
		return nil, fmt.Errorf("fetchUser%+v get: %w", fu, err)
	}

	err := fu.DB.User.Create().
		CopyUser(&u).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchUser%+v save: %w", fu, err)
	}

	return fu.DB.User.Get(ctx, u.ID)
}
