package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gnoverse/gh-sql/ent"
	"github.com/gnoverse/gh-sql/ent/issue"
	"github.com/gnoverse/gh-sql/pkg/model"
	"github.com/gnoverse/gh-sql/pkg/sync/internal/synchub"
	"golang.org/x/sync/errgroup"
)

// fetchIssue can be used to retrive an issue, knowing its repo and issue number.
type fetchIssue struct {
	Options
	owner       string
	repo        string
	issueNumber int64
}

var _ synchub.Resource[*ent.Issue] = fetchIssue{}

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

func (fi fetchIssue) Fetch(ctx context.Context, h *synchub.Hub) (*ent.Issue, error) {
	var i issueAndEdges
	if err := synchub.Get(ctx, h, fi.ID(), &i); err != nil {
		return nil, fmt.Errorf("fetchIssue%+v get: %w", fi, err)
	}
	return fi.fetch(ctx, h, i)
}

func (fi fetchIssue) fetch(ctx context.Context, h *synchub.Hub, i issueAndEdges) (*ent.Issue, error) {
	cr := fi.DB.Issue.Create().CopyIssue(&i.Issue)

	// Note: we don't pass the resulting context, as otherwise any `fetches`
	// originating from the `fetch` would fail once the errgroup here has terminated.
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// Fetch repository and set repo ID.
		repo, err := synchub.Fetch(ctx, h, fetchRepository{
			Options: fi.Options,
			owner:   fi.owner,
			repo:    fi.repo,
		})
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
			user, err := synchub.Fetch(ctx, h, fetchUser{fi.Options, login})
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
				h.Warn(err)
			}
			return nil
		})
	}
	h.Go(func() error {
		// TODO: probably have to fetch both timeline AND issue events to have a complete picture.
		// why, github, why?
		if err := fetchIssueEvents(ctx, h, fi); err != nil {
			h.Warn(err)
		}
		return nil
	})

	return fi.DB.Issue.Get(ctx, i.ID)
}

func fetchIssues(ctx context.Context, h *synchub.Hub, opts Options, repoOwner, repoName string) {
	iter := synchub.GetIterate[issueAndEdges](
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/issues?state=all&per_page=100", repoOwner, repoName))
	for i := range iter.Values {
		if !opts.CreatedAfter.IsZero() && i.CreatedAt.Before(opts.CreatedAfter) {
			break
		}
		iss, err := opts.DB.Issue.Query().
			Select(issue.FieldUpdatedAt).
			Where(issue.ID(i.ID)).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			h.Warn(err)
			continue
		}

		if iss == nil || !iss.UpdatedAt.Equal(i.UpdatedAt) {
			// this does not incur in an additional request; we have all the data
			// we want from this request already
			fi := fetchIssue{
				Options:     opts,
				owner:       repoOwner,
				repo:        repoName,
				issueNumber: i.Number,
			}
			fn, created := synchub.SetGetter(h, fi.ID(), func() (*ent.Issue, error) {
				return fi.fetch(ctx, h, i)
			})
			if created {
				if _, err := fn(); err != nil {
					h.Warn(err)
				}
			}
		}
	}
	if iter.Err() != nil {
		h.Warn(fmt.Errorf("fetchIssues(%q, %q): %w", repoOwner, repoName, iter.Err()))
	}
}

func fetchIssueComments(ctx context.Context, h *synchub.Hub, fi fetchIssue) error {
	iss, err := synchub.Fetch(ctx, h, fi)
	if err != nil {
		return err
	}
	type dstType struct {
		ent.IssueComment
		User *model.SimpleUser `json:"user"`
	}
	iter := synchub.GetIterate[dstType](
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/issues/%d/comments?per_page=100",
			fi.owner, fi.repo, fi.issueNumber),
	)
	for ic := range iter.Values {
		// Create issue comment
		cr := fi.DB.IssueComment.Create().
			CopyIssueComment(&ic.IssueComment).
			SetIssueID(iss.ID)
		// Assign user if possible.
		if ic.User != nil && ic.User.Login != "" {
			us, err := synchub.Fetch(ctx, h, fetchUser{fi.Options, ic.User.Login})
			if err != nil {
				h.Warn(err)
				continue
			}
			cr.SetUserID(us.ID)
		}
		err = cr.
			OnConflict().UpdateNewValues().
			Exec(ctx)
		if err != nil {
			h.Warn(err)
		}
	}
	if iter.Err() != nil {
		return fmt.Errorf("fetchIssueComments(%q): %w", fi.ID(), iter.Err())
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
	*f = fetchIssueEventType(dst)

	if err := f.Wrapper.UnmarshalJSON(b); err != nil {
		return err
	}

	f.TimelineEvent.Data = f.Wrapper
	return nil
}

func fetchIssueEvents(ctx context.Context, h *synchub.Hub, fi fetchIssue) error {
	iss, err := synchub.Fetch(ctx, h, fi)
	if err != nil {
		return err
	}
	iter := synchub.GetIterate[fetchIssueEventType](
		ctx, h,
		fmt.Sprintf("/repos/%s/%s/issues/%d/timeline?per_page=100",
			fi.owner, fi.repo, fi.issueNumber),
	)
	for iev := range iter.Values {
		// Create issue event.
		cr := fi.DB.TimelineEvent.Create().
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
			us, err := synchub.Fetch(ctx, h, fetchUser{fi.Options, actor.Login})
			if err != nil {
				return err
			}
			cr.SetActorID(us.ID)
		}

		err := cr.
			OnConflict().UpdateNewValues().
			Exec(ctx)
		if err != nil {
			h.Warn(fmt.Errorf("save event of type %v: %w", iev.Event, err))
		}
		return nil
	}
	if iter.Err() != nil {
		return fmt.Errorf("fetchIssueEvents(%q): %w", fi.ID(), iter.Err())
	}
	return nil
}
