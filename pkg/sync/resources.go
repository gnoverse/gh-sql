package sync

import (
	"context"
	"fmt"

	"github.com/gnolang/gh-sql/ent"
)

// fetchRepository can be used to fetch a GitHub repository, knowing its
// owner and name
type fetchRepository struct {
	owner, repo string
}

var _ resource[*ent.Repository] = fetchRepository{}

func (f fetchRepository) ID() string { return "/repos/" + f.owner + "/" + f.repo }

func (f fetchRepository) Fetch(ctx context.Context, h *hub) (*ent.Repository, error) {
	var r struct {
		ent.Repository
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
	}
	if err := httpGet(ctx, h, f.ID(), &r); err != nil {
		return nil, fmt.Errorf("fetchRepository%+v fetch: %w", f, err)
	}

	err := h.DB.Repository.Create().
		CopyRepository(&r.Repository).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchRepository%+v save: %w", f, err)
	}

	u, err := fetch(ctx, h, fetchUser{username: r.Owner.Login})
	if err != nil {
		return nil, err
	}
	savedRepo, err := h.DB.Repository.UpdateOneID(r.ID).
		SetOwnerID(u.ID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("link repo %d (%s) to owner %d (%s): %w",
			r.ID, r.Name, u.ID, u.Login, err)
	}
	savedRepo.Edges.Owner = u

	return savedRepo, nil
}

func fetchRepositories(ctx context.Context, h *hub, owner string, shouldFetch func(*ent.Repository) bool) {
	err := httpGetIterate(ctx, h, fmt.Sprintf("/users/%s/repos?per_page=100", owner), func(r *ent.Repository) error {
		if shouldFetch(r) {
			fetchAsync(ctx, h, fetchRepository{owner, r.Name})
		}
		return nil
	})
	if err != nil {
		h.report(fmt.Errorf("fetchRepositories(%q): %w", owner, err))
	}
}

// -----------------------------------------------------------------------------

// fetchUser can be used to fetch a user, knowing its
type fetchUser struct {
	username string
}

var _ resource[*ent.User] = fetchUser{}

func (fu fetchUser) ID() string { return "/users/" + fu.username }

func (fu fetchUser) Fetch(ctx context.Context, h *hub) (*ent.User, error) {
	var u ent.User
	if err := httpGet(ctx, h, fu.ID(), &u); err != nil {
		return nil, fmt.Errorf("fetchUser%+v get: %w", fu, err)
	}

	err := h.DB.User.Create().
		CopyUser(&u).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchUser%+v save: %w", fu, err)
	}

	return h.DB.User.Get(ctx, u.ID)
}
