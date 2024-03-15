package sync

import (
	"context"
	"fmt"

	"github.com/gnolang/gh-sql/ent"
	"github.com/gnolang/gh-sql/ent/issue"
	"golang.org/x/sync/errgroup"
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

	h.DB.Repository.Create().
		SetArchiveURL("this").
		SetWatchers(11)

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

	h.Go(func() error {
		fetchIssues(ctx, h, f.owner, f.repo)
		return nil
	})

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
		h.warn(fmt.Errorf("fetchRepositories(%q): %w", owner, err))
	}
}

// -----------------------------------------------------------------------------

// fetchUser can be used to retrive a user, knowing its username.
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

// -----------------------------------------------------------------------------

// fetchIssue can be used to retrive an issue, knowing its repo and issue number.
type fetchIssue struct {
	owner       string
	repo        string
	issueNumber int64
}

var _ resource[*ent.Issue] = fetchIssue{}

func (fi fetchIssue) ID() string {
	return fmt.Sprintf("/repos/%s/%s/issues/%d",
		fi.owner, fi.repo, fi.issueNumber)
}

func (fi fetchIssue) Fetch(ctx context.Context, h *hub) (*ent.Issue, error) {
	var i struct {
		ent.Issue
		User *struct {
			Login string `json:"login"`
		} `json:"user"`
		Assignees []struct {
			Login string `json:"login"`
		} `json:"assignees"`
		ClosedBy *struct {
			Login string `json:"login"`
		} `json:"closed_by"`
	}
	if err := httpGet(ctx, h, fi.ID(), &i); err != nil {
		return nil, fmt.Errorf("fetchIssue%+v get: %w", fi, err)
	}

	cr := h.DB.Issue.Create().CopyIssue(&i.Issue)

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// Fetch repository and set repo ID.
		repo, err := fetch(egCtx, h, fetchRepository{fi.owner, fi.repo})
		if err != nil {
			return err
		}
		cr.SetRepositoryID(repo.ID)
		return nil
	})
	localGetUser := func(login string, setter func(id int64) *ent.IssueCreate) {
		if login == "" {
			return
		}
		eg.Go(func() error {
			// Fetch creator
			user, err := fetch(egCtx, h, fetchUser{login})
			if err != nil {
				return err
			}
			setter(user.ID)
			return nil
		})
	}
	if i.User != nil {
		localGetUser(i.User.Login, cr.SetUserID)
	}
	for _, assignee := range i.Assignees {
		login := assignee.Login
		localGetUser(login, func(i int64) *ent.IssueCreate { return cr.AddAssigneeIDs(i) })
	}
	if i.ClosedBy != nil {
		localGetUser(i.ClosedBy.Login, cr.SetClosedByID)
	}

	err := eg.Wait()
	if err != nil {
		return nil, fmt.Errorf("fetchIssue%+v fetch deps: %w", fi, err)
	}

	err = cr.OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchIssue%+v save: %w", fi, err)
	}

	h.Go(func() error {
		if err := fetchIssueComments(ctx, h, fi.owner, fi.repo, fi.issueNumber); err != nil {
			h.warn(err)
		}
		return nil
	})

	return h.DB.Issue.Get(ctx, i.ID)
}

func fetchIssues(ctx context.Context, h *hub, repoOwner, repoName string) {
	err := httpGetIterate(ctx, h, fmt.Sprintf("/repos/%s/%s/issues?state=all&per_page=100", repoOwner, repoName), func(i *ent.Issue) error {
		iss, err := h.DB.Issue.Query().
			Select(issue.FieldUpdatedAt).
			Where(issue.ID(i.ID)).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return err
		}
		// TODO: instead of double-fetching, we could save the issue here.
		// fields that we want to make sure that are the same in the two:
		// state_reason, body, labels, draft, assignees, active_lock_reason
		// pull_request (existance), reactions.
		if iss == nil || !iss.UpdatedAt.Equal(i.UpdatedAt) {
			fetchAsync(ctx, h, fetchIssue{repoOwner, repoName, i.Number})
		}
		return nil
	})
	if err != nil {
		h.warn(fmt.Errorf("fetchIssues(%q, %q): %w", repoOwner, repoName, err))
	}
}

func fetchIssueComments(ctx context.Context, h *hub,
	repoOwner, repoName string, issueNumber int64,
) error {
	iss, err := fetch(ctx, h, fetchIssue{repoOwner, repoName, issueNumber})
	if err != nil {
		return err
	}
	type dstType struct {
		ent.IssueComment
		User *struct {
			Login string `json:"login"`
		} `json:"user"`
	}
	err = httpGetIterate(
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/issues/%d/comments?per_page=100",
			repoOwner, repoName, issueNumber),
		func(ic *dstType) error {
			// Create issue comment
			cr := h.DB.IssueComment.Create().
				CopyIssueComment(&ic.IssueComment).
				SetIssueID(iss.ID)
			// Assign user if possible.
			if ic.User != nil && ic.User.Login != "" {
				us, err := fetch(ctx, h, fetchUser{ic.User.Login})
				if err != nil {
					return err
				}
				cr.SetUserID(us.ID)
			}
			return cr.
				OnConflict().UpdateNewValues().
				Exec(ctx)
		},
	)
	if err != nil {
		return fmt.Errorf("fetchIssueComments(%q, %q, %d): %w", repoOwner, repoName, issueNumber, err)
	}
	return nil
}

// PULL REQUESTS
