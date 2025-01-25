package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gnolang/gh-sql/ent"
	"github.com/gnolang/gh-sql/ent/issue"
	"github.com/gnolang/gh-sql/ent/pullrequest"
	"github.com/gnolang/gh-sql/ent/repository"
	"github.com/gnolang/gh-sql/pkg/model"
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
		Owner model.SimpleUser `json:"owner"`
	}
	if err := httpGet(ctx, h, f.ID(), &r); err != nil {
		return nil, fmt.Errorf("fetchRepository%+v fetch: %w", f, err)
	}

	u, err := fetch(ctx, h, fetchUser{username: r.Owner.Login})
	if err != nil {
		return nil, err
	}

	err = h.DB.Repository.Create().
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
		fetchIssues(ctx, h, r.Owner.Login, r.Name)
		fetchPulls(ctx, h, r.Owner.Login, r.Name)
		return nil
	})

	return h.DB.Repository.Get(ctx, r.ID)
}

func fetchRepositories(ctx context.Context, h *hub, owner string, shouldFetch func(*ent.Repository) bool) {
	err := httpGetIterate(
		ctx, h,
		fmt.Sprintf("/users/%s/repos?per_page=100", owner),
		func(r *ent.Repository) error {
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

type issueAndEdges struct {
	ent.Issue

	User      *model.SimpleUser  `json:"user"`
	ClosedBy  *model.SimpleUser  `json:"closed_by"`
	Assignees []model.SimpleUser `json:"assignees"`
}

func (fi fetchIssue) Fetch(ctx context.Context, h *hub) (*ent.Issue, error) {
	var i issueAndEdges
	if err := httpGet(ctx, h, fi.ID(), &i); err != nil {
		return nil, fmt.Errorf("fetchIssue%+v get: %w", fi, err)
	}
	return fi.fetch(ctx, h, i)
}

func (fi fetchIssue) fetch(ctx context.Context, h *hub, i issueAndEdges) (*ent.Issue, error) {
	cr := h.DB.Issue.Create().CopyIssue(&i.Issue)

	// Note: we don't pass the resulting context, as otherwise any `fetches`
	// originating from the `fetch` would fail once the errgroup here has terminated.
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// Fetch repository and set repo ID.
		repo, err := fetch(ctx, h, fetchRepository{fi.owner, fi.repo})
		if err != nil {
			return err
		}
		cr.SetRepositoryID(repo.ID)
		return nil
	})
	var setterLock sync.Mutex
	localGetUser := func(login string, setter func(id int64) *ent.IssueCreate) {
		if login == "" {
			return
		}
		eg.Go(func() error {
			// Fetch creator
			user, err := fetch(ctx, h, fetchUser{login})
			if err != nil {
				return err
			}
			setterLock.Lock()
			defer setterLock.Unlock()
			setter(user.ID)
			return nil
		})
	}
	if i.User != nil {
		localGetUser(i.User.Login, cr.SetUserID)
	}
	if i.ClosedBy != nil {
		localGetUser(i.ClosedBy.Login, cr.SetClosedByID)
	}
	for _, assignee := range i.Assignees {
		login := assignee.Login
		localGetUser(login, func(i int64) *ent.IssueCreate { return cr.AddAssigneeIDs(i) })
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

	if i.CommentsCount > 0 {
		h.Go(func() error {
			if err := fetchIssueComments(ctx, h, fi); err != nil {
				h.warn(err)
			}
			return nil
		})
	}
	h.Go(func() error {
		// TODO: probably have to fetch both timeline AND issue events to have a complete picture.
		// why, github, why?
		if err := fetchIssueEvents(ctx, h, fi); err != nil {
			h.warn(err)
		}
		return nil
	})

	return h.DB.Issue.Get(ctx, i.ID)
}

func fetchIssues(ctx context.Context, h *hub, repoOwner, repoName string) {
	err := httpGetIterate(
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/issues?state=all&per_page=100", repoOwner, repoName),
		func(i issueAndEdges) error {
			iss, err := h.DB.Issue.Query().
				Select(issue.FieldUpdatedAt).
				Where(issue.ID(i.ID)).
				Only(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return err
			}

			if iss == nil || !iss.UpdatedAt.Equal(i.UpdatedAt) {
				// this does not incur in an additional request; we have all the data
				// we want from this request already
				fi := fetchIssue{repoOwner, repoName, i.Number}
				fn, created := setUpdatedFunc(h, fi.ID(), func() (*ent.Issue, error) {
					return fi.fetch(ctx, h, i)
				})
				if created {
					if _, err := fn(); err != nil {
						h.warn(err)
					}
				}
			}
			return nil
		})
	if err != nil {
		h.warn(fmt.Errorf("fetchIssues(%q, %q): %w", repoOwner, repoName, err))
	}
}

func fetchIssueComments(ctx context.Context, h *hub, fi fetchIssue) error {
	iss, err := fetch(ctx, h, fi)
	if err != nil {
		return err
	}
	type dstType struct {
		ent.IssueComment
		User *model.SimpleUser `json:"user"`
	}
	err = httpGetIterate(
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/issues/%d/comments?per_page=100",
			fi.owner, fi.repo, fi.issueNumber),
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
		return fmt.Errorf("fetchIssueComments(%q): %w", fi.ID(), err)
	}
	return nil
}

type fetchIssueEventType struct {
	ent.TimelineEvent
	Actor   *model.SimpleUser `json:"actor"`
	User    *model.SimpleUser `json:"user"`
	Wrapper model.TimelineEventWrapper
}

func (f *fetchIssueEventType) UnmarshalJSON(b []byte) error {
	type dstType fetchIssueEventType
	var dst dstType
	if err := json.Unmarshal(b, &dst); err != nil {
		return err
	}

	if err := f.Wrapper.UnmarshalJSON(b); err != nil {
		return err
	}

	f.TimelineEvent.Data = f.Wrapper
	return nil
}

func fetchIssueEvents(ctx context.Context, h *hub, fi fetchIssue) error {
	iss, err := fetch(ctx, h, fi)
	if err != nil {
		return err
	}
	err = httpGetIterate(
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/issues/%d/timeline?per_page=100",
			fi.owner, fi.repo, fi.issueNumber),
		func(iev *fetchIssueEventType) error {
			// Create issue event.
			cr := h.DB.TimelineEvent.Create().
				CopyTimelineEvent(&iev.TimelineEvent).
				SetIssueID(iss.ID)

			// This is a half-lie, but attempt to use the User as the Actor if
			// it is not set. This is useful in some events like `reviewed`,
			// which has user but not actor, and allows us to link many more
			// events.
			actor := iev.Actor
			if actor == nil && iev.User != nil {
				actor = iev.User
			}

			// Assign user if possible.
			if actor != nil {
				us, err := fetch(ctx, h, fetchUser{actor.Login})
				if err != nil {
					return err
				}
				cr.SetActorID(us.ID)
			}

			err := cr.
				OnConflict().UpdateNewValues().
				Exec(ctx)
			if err != nil {
				h.warn(fmt.Errorf("save event of type %v: %w", iev.Event, err))
			}
			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("fetchIssueEvents(%q): %w", fi.ID(), err)
	}
	return nil
}

// fetchPullRequest can be used to retrive an issue, knowing its repo and issue number.
type fetchPullRequest struct {
	owner      string
	repo       string
	pullNumber int64
}

var _ resource[*ent.PullRequest] = fetchPullRequest{}

func (fi fetchPullRequest) ID() string {
	return fmt.Sprintf("/repos/%s/%s/pulls/%d",
		fi.owner, fi.repo, fi.pullNumber)
}

func (fi fetchPullRequest) Fetch(ctx context.Context, h *hub) (*ent.PullRequest, error) {
	var pr pullAndEdges
	if err := httpGet(ctx, h, fi.ID(), &pr); err != nil {
		return nil, fmt.Errorf("fetchPullRequest%+v get: %w", fi, err)
	}
	return fi.fetch(ctx, h, pr)
}

type pullAndEdges struct {
	ent.PullRequest

	User               *model.SimpleUser  `json:"user"`
	Assignees          []model.SimpleUser `json:"assignees"`
	RequestedReviewers []model.SimpleUser `json:"requested_reviewers"`
}

func (fi fetchPullRequest) fetch(ctx context.Context, h *hub, pr pullAndEdges) (*ent.PullRequest, error) {
	cr := h.DB.PullRequest.Create().CopyPullRequest(&pr.PullRequest)

	eg, _ := errgroup.WithContext(ctx)
	iss, err := h.DB.Issue.Query().
		Where(
			issue.HasRepositoryWith(repository.FullName(fi.owner+"/"+fi.repo)),
			issue.Number(pr.Number),
		).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if iss == nil || iss.UpdatedAt.Before(pr.UpdatedAt) {
		// Fetch issue.
		eg.Go(func() error {
			iss, err := fetch(ctx, h, fetchIssue{fi.owner, fi.repo, fi.pullNumber})
			if err != nil {
				return err
			}
			cr.SetIssueID(iss.ID)
			return nil
		})
	} else {
		cr.SetIssueID(iss.ID)
	}
	eg.Go(func() error {
		// Fetch repository and set repo ID.
		repo, err := fetch(ctx, h, fetchRepository{fi.owner, fi.repo})
		if err != nil {
			return err
		}
		cr.SetRepositoryID(repo.ID)
		return nil
	})
	var setterLock sync.Mutex
	localGetUser := func(login string, setter func(id int64) *ent.PullRequestCreate) {
		if login == "" {
			return
		}
		eg.Go(func() error {
			// Fetch creator
			user, err := fetch(ctx, h, fetchUser{login})
			if err != nil {
				return err
			}
			setterLock.Lock()
			defer setterLock.Unlock()
			setter(user.ID)
			return nil
		})
	}
	if pr.User != nil {
		localGetUser(pr.User.Login, cr.SetUserID)
	}
	for _, assignee := range pr.Assignees {
		login := assignee.Login
		localGetUser(login, func(i int64) *ent.PullRequestCreate { return cr.AddAssigneeIDs(i) })
	}
	for _, reviewer := range pr.RequestedReviewers {
		login := reviewer.Login
		localGetUser(login, func(i int64) *ent.PullRequestCreate { return cr.AddRequestedReviewerIDs(i) })
	}

	err = eg.Wait()
	if err != nil {
		return nil, fmt.Errorf("fetchPullRequest%+v fetch deps: %w", fi, err)
	}

	err = cr.OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchPullRequest%+v save: %w", fi, err)
	}

	return h.DB.PullRequest.Get(ctx, pr.ID)
}

func fetchPulls(ctx context.Context, h *hub, repoOwner, repoName string) {
	err := httpGetIterate(
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/pulls?state=all&per_page=100", repoOwner, repoName),
		func(pr pullAndEdges) error {
			oldPR, err := h.DB.PullRequest.Query().
				Select(pullrequest.FieldUpdatedAt).
				Where(pullrequest.ID(pr.ID)).
				Only(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return err
			}

			if oldPR == nil || !oldPR.UpdatedAt.Equal(pr.UpdatedAt) {
				// this does not incur in an additional request; we have all the data
				// we want from this request already
				fi := fetchPullRequest{repoOwner, repoName, pr.Number}
				fn, created := setUpdatedFunc(h, fi.ID(), func() (*ent.PullRequest, error) {
					return fi.fetch(ctx, h, pr)
				})
				if created {
					if _, err := fn(); err != nil {
						h.warn(err)
					}
				}
			}
			return nil
		})
	if err != nil {
		h.warn(fmt.Errorf("fetchIssues(%q, %q): %w", repoOwner, repoName, err))
	}
}
