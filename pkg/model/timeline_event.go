package model

// This file contains various types of timeline events.
// These extend ent.TimelineEvent in the GitHub API.
// Reference: https://docs.github.com/en/rest/using-the-rest-api/issue-event-types?apiVersion=2022-11-28.
// Unfortunately, I cannot recommend the reference because it does not contain
// all the fields. Not that the OpenAPI schema is any better, it doesn't even
// associate event names to corresponding data structures.

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"time"

	"github.com/gnolang/gh-sql/pkg/model/internal/jsoncatcher"
)

// TimelineEventWrapper wraps all TimelineEvents, unmarshaling JSON data
// into the appropriate type by looking at its "events" key.
type TimelineEventWrapper struct {
	TimelineEvent
}

// MarshalJSON implements [json.Marshaler].
// This will not encode a full GitHub event; but it will enrichen the
// underlying event's marshaling with the event's name.
func (ew TimelineEventWrapper) MarshalJSON() ([]byte, error) {
	switch ew := ew.TimelineEvent.(type) {
	case nil:
		return []byte("null"), nil
	case SimpleEvent:
		return json.Marshal(ew)
	}

	// Marshal underlying event and append event name.
	marshaled, err := json.Marshal(ew.TimelineEvent)
	if err != nil {
		return nil, err
	}

	if len(marshaled) == 2 {
		return []byte(`{"event":"` + ew.Name() + `"}`), nil
	}
	marshaled = append([]byte(`{"event":"`+ew.Name()+`",`), marshaled[1:]...)
	return marshaled, nil
}

var tSimpleEvent = reflect.TypeOf(SimpleEvent{})

// EventsMissedData can be set if [TimelineEventWrapper.UnmarshalJSON]
// should log data (JSON fields) missed while unmarshalling.
// A warning will be logged on stderr, and the log of the missed data will be
// written to this writer.
var EventsMissedData io.Writer

// Used as a jsoncatcher.CatchOptions.Skip function.
// These fields are caught elsewhere (ie. by the Ent resource), so if the
// marshaled version doesn't contain these fields it's OK.
func missedDataSkip(k []byte) bool {
	switch string(k) {
	case
		"id",
		"node_id",
		"url",
		"event",
		"commit_id",
		"commit_url",
		"created_at",
		"actor",
		"performed_via_github_app":
		return true
	}
	return false
}

// UnmarshalJSON implements [json.Unmarshaler].
// It expects a JSON object with field "event", identifying the name of the
// event.
func (ew *TimelineEventWrapper) UnmarshalJSON(b []byte) (resultErr error) {
	var evtName SimpleEvent
	if err := json.Unmarshal(b, &evtName); err != nil {
		return fmt.Errorf("getting event name from object: %w", err)
	}

	// use jsoncatcher to detect when unmarshaling is lossy.
	// this mostly happens because GitHub gave up on updating their API documentation.
	if EventsMissedData != nil {
		defer func() {
			if resultErr != nil {
				// skip if we already have an error.
				return
			}

			opts := jsoncatcher.CatchOptions{
				Skip:   missedDataSkip,
				Missed: EventsMissedData,
			}
			err := opts.Catch(b, ew)
			if err != nil {
				log.Printf("warning: event %s: %v\n", evtName.Event, err)
			}
		}()
	}

	// Get the expected type to decode.
	te, ok := timelineEvents[evtName.Event]
	if !ok {
		return fmt.Errorf("unknown event: %q", evtName.Event)
	}
	// Find the reflect type. If it is SimpleEvent, use the
	// SimpleEvent as the TimelineEvent and return.
	teType := reflect.TypeOf(te)
	if teType == tSimpleEvent {
		ew.TimelineEvent = evtName
		return nil
	}
	// Create new teType and unmarshal into it.
	dst := reflect.New(teType)
	if err := json.Unmarshal(b, dst.Interface()); err != nil {
		return fmt.Errorf("unmarshaling into %s: %w", teType.Name(), err)
	}

	ew.TimelineEvent = dst.Interface().(TimelineEvent)
	return nil
}

var timelineEvents = map[string]TimelineEvent{
	(LabeledIssueEvent{}).Name():               LabeledIssueEvent{},
	(UnlabeledIssueEvent{}).Name():             UnlabeledIssueEvent{},
	(MilestonedIssueEvent{}).Name():            MilestonedIssueEvent{},
	(DemilestonedIssueEvent{}).Name():          DemilestonedIssueEvent{},
	(RenamedIssueEvent{}).Name():               RenamedIssueEvent{},
	(ReviewRequestedIssueEvent{}).Name():       ReviewRequestedIssueEvent{},
	(ReviewRequestRemovedIssueEvent{}).Name():  ReviewRequestRemovedIssueEvent{},
	(ReviewDismissedIssueEvent{}).Name():       ReviewDismissedIssueEvent{},
	(LockedIssueEvent{}).Name():                LockedIssueEvent{},
	(AddedToProjectIssueEvent{}).Name():        AddedToProjectIssueEvent{},
	(MovedColumnsInProjectIssueEvent{}).Name(): MovedColumnsInProjectIssueEvent{},
	(RemovedFromProjectIssueEvent{}).Name():    RemovedFromProjectIssueEvent{},
	(ConvertedNoteToIssueIssueEvent{}).Name():  ConvertedNoteToIssueIssueEvent{},
	(TimelineCommentEvent{}).Name():            TimelineCommentEvent{},
	(TimelineCrossReferencedEvent{}).Name():    TimelineCrossReferencedEvent{},
	(TimelineCommittedEvent{}).Name():          TimelineCommittedEvent{},
	(TimelineReviewedEvent{}).Name():           TimelineReviewedEvent{},
	(TimelineAssignedIssueEvent{}).Name():      TimelineAssignedIssueEvent{},
	(TimelineUnassignedIssueEvent{}).Name():    TimelineUnassignedIssueEvent{},
	(ClosedEvent{}).Name():                     ClosedEvent{},
	(ReopenedEvent{}).Name():                   ReopenedEvent{},

	"mentioned":                       SimpleEvent{},
	"merged":                          SimpleEvent{},
	"ready_for_review":                SimpleEvent{},
	"referenced":                      SimpleEvent{},
	"subscribed":                      SimpleEvent{},
	"unsubscribed":                    SimpleEvent{},
	"convert_to_draft":                SimpleEvent{},
	"head_ref_deleted":                SimpleEvent{},
	"head_ref_restored":               SimpleEvent{},
	"head_ref_force_pushed":           SimpleEvent{},
	"comment_deleted":                 SimpleEvent{},
	"automatic_base_change_succeeded": SimpleEvent{},
	"automatic_base_change_failed":    SimpleEvent{},
	"base_ref_changed":                SimpleEvent{},
	"connected":                       SimpleEvent{},
	"disconnected":                    SimpleEvent{},
	"converted_to_discussion":         SimpleEvent{},
	"deployed":                        SimpleEvent{},
	"deployment_environment_changed":  SimpleEvent{},
	"marked_as_duplicate":             SimpleEvent{},
	"pinned":                          SimpleEvent{},
	"transferred":                     SimpleEvent{},
	"unmarked_as_duplicate":           SimpleEvent{},
	"unpinned":                        SimpleEvent{},
	"user_blocked":                    SimpleEvent{},
	"auto_merge_disabled":             SimpleEvent{},

	// NOTE: these events are found through experimental evidence;
	// because as far as the GitHub documentation is concerned, they don't exist
	// (at time of writing): https://docs.github.com/en/search?query=auto_squash_enabled
	"auto_squash_enabled":   SimpleEvent{},
	"base_ref_force_pushed": SimpleEvent{},
}

// TimelineEvent is the interface for all of the timeline events.
type TimelineEvent interface {
	assertTimelineEvent()
	Name() string
}

// Helper type to make timeline events implement TimelineEvent
type timelineEvent struct{}

func (timelineEvent) assertTimelineEvent() {}

// SimpleEvent is used for the events which don't have additional information.
type SimpleEvent struct {
	timelineEvent

	Event string `json:"event"`
}

func (ev SimpleEvent) Name() string { return ev.Event }

// LabeledIssueEvent is the GitHub event for when an issue is labeled.
type LabeledIssueEvent struct {
	timelineEvent

	Label struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"label"`
}

// Name is used to associate this type with the labeled event.
func (LabeledIssueEvent) Name() string { return "labeled" }

// UnlabeledIssueEvent is the GitHub event for when an issue is labeled.
type UnlabeledIssueEvent struct {
	timelineEvent

	Label struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"label"`
}

// Name is used to associate this type with the unlabeled event.
func (UnlabeledIssueEvent) Name() string { return "unlabeled" }

// MilestonedIssueEvent is the GitHub event for when an issue is milestoned.
type MilestonedIssueEvent struct {
	timelineEvent

	Milestone struct {
		Title string `json:"title"`
	} `json:"milestone"`
}

// Name is used to associate this type with the milestone event.
func (MilestonedIssueEvent) Name() string { return "milestoned" }

// DemilestonedIssueEvent is the GitHub event for when an issue has its milestone
// removed.
type DemilestonedIssueEvent struct {
	timelineEvent

	Milestone struct {
		Title string `json:"title"`
	} `json:"milestone"`
}

// Name is used to associate this type with the demilestoned event.
func (DemilestonedIssueEvent) Name() string { return "demilestoned" }

// RenamedIssueEvent is the GitHub event for when an issue is renamed.
type RenamedIssueEvent struct {
	timelineEvent

	Rename struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"rename"`
}

// Name is used to associate this type with the renamed event.
func (RenamedIssueEvent) Name() string { return "renamed" }

// ReviewRequestedIssueEvent is the GitHub event for when an issue (PR) receives a
// new review request.
type ReviewRequestedIssueEvent struct {
	timelineEvent

	ReviewRequester struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"review_requester"`
	RequestedTeam *struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"requested_team,omitempty"`
	RequestedReviewer *struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"requested_reviewer,omitempty"`
}

// Name is used to associate this type with the review_requested event.
func (ReviewRequestedIssueEvent) Name() string { return "review_requested" }

// ReviewRequestRemovedIssueEvent is the GitHub event for when an issue (PR) has one of
// its existing review requests removed.
type ReviewRequestRemovedIssueEvent struct {
	timelineEvent

	ReviewRequester struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"review_requester"`
	RequestedTeam *struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"requested_team,omitempty"`
	RequestedReviewer *struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"requested_reviewer,omitempty"`
}

// Name is used to associate this type with the review_request_removed event.
func (ReviewRequestRemovedIssueEvent) Name() string { return "review_request_removed" }

// ReviewDismissedIssueEvent is the GitHub event for when a review on an issue (PR)
// is dismissed.
type ReviewDismissedIssueEvent struct {
	timelineEvent

	DismissedReview struct {
		State             string  `json:"state"`
		ReviewID          int     `json:"review_id"`
		DismissalMessage  *string `json:"dismissal_message"`
		DismissalCommitID string  `json:"dismissal_commit_id"`
	} `json:"dismissed_review"`
}

// Name is used to associate this type with the review_dismissed event.
func (ReviewDismissedIssueEvent) Name() string { return "review_dismissed" }

// LockedIssueEvent is the GitHub event for when a conversation on an issue is
// locked.
type LockedIssueEvent struct {
	timelineEvent

	LockReason *string `json:"lock_reason"`
}

// Name is used to associate this type with the locked event.
func (LockedIssueEvent) Name() string { return "locked" }

// UnlockedIssueEvent is the GitHub event for when a conversation on an issue is
// unlocked.
type UnlockedIssueEvent struct {
	timelineEvent

	LockReason *string `json:"lock_reason"`
}

// Name is used to associate this type with the locked event.
func (UnlockedIssueEvent) Name() string { return "unlocked" }

type ProjectCard struct {
	ID                 int64  `json:"id"`
	URL                string `json:"url"`
	ProjectID          int64  `json:"project_id"`
	ProjectURL         string `json:"project_url"`
	ColumnName         string `json:"column_name"`
	PreviousColumnName string `json:"previous_column_name,omitempty"`
}

type AddedToProjectIssueEvent struct {
	timelineEvent

	ProjectCard *ProjectCard `json:"project_card"`
}

func (AddedToProjectIssueEvent) Name() string { return "added_to_project" }

type MovedColumnsInProjectIssueEvent struct {
	timelineEvent

	ProjectCard *ProjectCard `json:"project_card"`
}

func (MovedColumnsInProjectIssueEvent) Name() string { return "moved_columns_in_project" }

type RemovedFromProjectIssueEvent struct {
	timelineEvent

	ProjectCard *ProjectCard `json:"project_card"`
}

func (RemovedFromProjectIssueEvent) Name() string { return "removed_from_project" }

type ConvertedNoteToIssueIssueEvent struct {
	timelineEvent

	ProjectCard *ProjectCard `json:"project_card"`
}

func (ConvertedNoteToIssueIssueEvent) Name() string { return "converted_note_to_issue" }

// SimpleUser is used to keep track of user information in events.
// The GitHub API actually contains more information, but for our
// purposes it redundant.
type SimpleUser struct {
	Login string `json:"login"`
	ID    int64  `json:"id"`
}

type TimelineCommentEvent struct {
	timelineEvent

	Body              string            `json:"body"`
	BodyText          string            `json:"body_text"`
	BodyHTML          string            `json:"body_html"`
	HTML_URL          string            `json:"html_url"`
	User              SimpleUser        `json:"user"`
	UpdatedAt         time.Time         `json:"updated_at"`
	IssueUrl          string            `json:"issue_url"`
	AuthorAssociation AuthorAssociation `json:"author_association"`
	Reactions         ReactionRollup    `json:"reactions"`
}

func (TimelineCommentEvent) Name() string { return "commented" }

// SimpleRepository is a reduced issue struct, keeping track of important information but avoiding redundancy.
type SimpleRepository struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	FullName string     `json:"full_name"`
	Owner    SimpleUser `json:"owner"`
}

// SimpleIssue is a reduced issue struct, keeping track of important information but avoiding redundancy.
type SimpleIssue struct {
	ID         int64            `json:"id"`
	Number     int64            `json:"number"`
	Repository SimpleRepository `json:"repository"`
}

type TimelineCrossReferencedEvent struct {
	timelineEvent

	UpdatedAt time.Time `json:"updated_at"`
	Source    struct {
		Type  string      `json:"type"`
		Issue SimpleIssue `json:"issue"`
	} `json:"source"`
}

func (TimelineCrossReferencedEvent) Name() string { return "cross-referenced" }

type CommitUser struct {
	Date  time.Time `json:"date"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

type VerificationInfo struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

type TimelineCommittedEvent struct {
	timelineEvent

	Sha       string     `json:"sha"`
	Author    CommitUser `json:"author"`
	Committer CommitUser `json:"committer"`
	Message   string     `json:"message"`
	Tree      struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"tree"`
	Parents []struct {
		Sha      string `json:"sha"`
		URL      string `json:"url"`
		HTML_URL string `json:"html_url"`
	} `json:"parents"`
	Verification VerificationInfo `json:"verification"`
	HtmlUrl      string           `json:"html_url"`
}

func (TimelineCommittedEvent) Name() string { return "committed" }

type TimelineReviewedEvent struct {
	timelineEvent

	User           SimpleUser `json:"user"`
	Body           string     `json:"body"`
	State          string     `json:"state"`
	HtmlUrl        string     `json:"html_url"`
	PullRequestUrl string     `json:"pull_request_url"`
	Links          struct {
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
		PullRequest struct {
			Href string `json:"href"`
		} `json:"pull_request"`
	} `json:"_links"`
	SubmittedAt       time.Time         `json:"submitted_at"`
	BodyHtml          string            `json:"body_html"`
	BodyText          string            `json:"body_text"`
	AuthorAssociation AuthorAssociation `json:"author_association"`
}

func (TimelineReviewedEvent) Name() string { return "reviewed" }

type TimelineAssignedIssueEvent struct {
	timelineEvent

	Assignee SimpleUser `json:"assignee"`
}

func (TimelineAssignedIssueEvent) Name() string { return "assigned" }

type TimelineUnassignedIssueEvent struct {
	timelineEvent

	Assignee SimpleUser `json:"assignee"`
}

func (TimelineUnassignedIssueEvent) Name() string { return "unassigned" }

type ClosedEvent struct {
	timelineEvent

	StateReason StateReason `json:"state_reason"`
}

func (ClosedEvent) Name() string { return "closed" }

type ReopenedEvent struct {
	timelineEvent

	StateReason StateReason `json:"state_reason"`
}

func (ReopenedEvent) Name() string { return "reopened" }

/*
TODO:
	"$ref": "#/components/schemas/timeline-line-commented-event"
	"$ref": "#/components/schemas/timeline-commit-commented-event"
	"$ref": "#/components/schemas/state-change-issue-event"
*/
