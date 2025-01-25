package sync

import (
	"context"
	"fmt"
	"sync"

	"github.com/gnoverse/gh-sql/ent"
	"github.com/gnoverse/gh-sql/ent/issue"
	"github.com/gnoverse/gh-sql/ent/pullrequest"
	"github.com/gnoverse/gh-sql/ent/repository"
	"github.com/gnoverse/gh-sql/pkg/model"
	"github.com/gnoverse/gh-sql/pkg/sync/internal/synchub"
	"golang.org/x/sync/errgroup"
)

// fetchPullRequest can be used to retrive an issue, knowing its repo and issue number.
type fetchPullRequest struct {
	Options
	owner      string
	repo       string
	pullNumber int64
}

var _ synchub.Resource[*ent.PullRequest] = fetchPullRequest{}

func (fi fetchPullRequest) ID() string {
	return fmt.Sprintf("/repos/%s/%s/pulls/%d",
		fi.owner, fi.repo, fi.pullNumber)
}

func (fi fetchPullRequest) Fetch(ctx context.Context, h *synchub.Hub) (*ent.PullRequest, error) {
	var pr pullAndEdges
	if err := synchub.Get(ctx, h, fi.ID(), &pr); err != nil {
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

func (fpr fetchPullRequest) fetch(ctx context.Context, h *synchub.Hub, pr pullAndEdges) (*ent.PullRequest, error) {
	cr := fpr.DB.PullRequest.Create().CopyPullRequest(&pr.PullRequest)

	eg, _ := errgroup.WithContext(ctx)
	iss, err := fpr.DB.Issue.Query().
		Where(
			issue.HasRepositoryWith(repository.FullName(fpr.owner+"/"+fpr.repo)),
			issue.Number(pr.Number),
		).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if iss == nil || iss.UpdatedAt.Before(pr.UpdatedAt) {
		// Fetch issue.
		eg.Go(func() error {
			iss, err := synchub.Fetch(ctx, h, fetchIssue{
				Options:     fpr.Options,
				owner:       fpr.owner,
				repo:        fpr.repo,
				issueNumber: fpr.pullNumber,
			})
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
		repo, err := synchub.Fetch(ctx, h, fetchRepository{
			Options: fpr.Options,
			owner:   fpr.owner,
			repo:    fpr.repo,
		})
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
			user, err := synchub.Fetch(ctx, h, fetchUser{fpr.Options, login})
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
		return nil, fmt.Errorf("fetchPullRequest%+v fetch deps: %w", fpr, err)
	}

	err = cr.OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchPullRequest%+v save: %w", fpr, err)
	}

	return fpr.DB.PullRequest.Get(ctx, pr.ID)
}

func fetchPulls(ctx context.Context, h *synchub.Hub, opts Options, repoOwner, repoName string) {
	iter := synchub.GetIterate[pullAndEdges](
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/pulls?state=all&per_page=100", repoOwner, repoName))
	for pr := range iter.Values {
		if !opts.CreatedAfter.IsZero() && pr.CreatedAt.Before(opts.CreatedAfter) {
			break
		}
		oldPR, err := opts.DB.PullRequest.Query().
			Select(pullrequest.FieldUpdatedAt).
			Where(pullrequest.ID(pr.ID)).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			h.Warn(err)
			continue
		}

		if oldPR == nil || !oldPR.UpdatedAt.Equal(pr.UpdatedAt) {
			// this does not incur in an additional request; we have all the data
			// we want from this request already
			fi := fetchPullRequest{opts, repoOwner, repoName, pr.Number}
			fn, created := synchub.SetGetter(h, fi.ID(), func() (*ent.PullRequest, error) {
				return fi.fetch(ctx, h, pr)
			})
			if created {
				if _, err := fn(); err != nil {
					h.Warn(err)
				}
			}
		}
	}
	if iter.Err() != nil {
		h.Warn(fmt.Errorf("fetchPulls(%q, %q): %w", repoOwner, repoName, iter.Err()))
	}
}
