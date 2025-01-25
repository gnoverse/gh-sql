// Package model defines some types used within the ent schemas.
package model

// License is a License on the GitHub API. In the schema, this is
// nullable-license-simple.
type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
	SpdxID string `json:"spdx_id"`
}

// AuthorAssociation is an enum of possible associations of an issue or issue comment's
// author, with the repository of the issue.
type AuthorAssociation string

// Possible values of AuthorAssociation.
const (
	AssociationCollaborator         AuthorAssociation = "COLLABORATOR"
	AssociationContributor          AuthorAssociation = "CONTRIBUTOR"
	AssociationFirstTimer           AuthorAssociation = "FIRST_TIMER"
	AssociationFirstTimeContributor AuthorAssociation = "FIRST_TIME_CONTRIBUTOR"
	AssociationMannequin            AuthorAssociation = "MANNEQUIN"
	AssociationMember               AuthorAssociation = "MEMBER"
	AssociationOwner                AuthorAssociation = "OWNER"
	AssociationNone                 AuthorAssociation = "NONE"
)

var authorAssociations = [...]string{
	string(AssociationCollaborator),
	string(AssociationContributor),
	string(AssociationFirstTimer),
	string(AssociationFirstTimeContributor),
	string(AssociationMannequin),
	string(AssociationMember),
	string(AssociationOwner),
	string(AssociationNone),
}

// Values returns the possible values of an [AuthorAssociation].
func (AuthorAssociation) Values() []string {
	return authorAssociations[:]
}

// ReactionRollup associates to each reaction a number of users who clicked it.
type ReactionRollup struct {
	ThumbsUp   int    `json:"+1"`
	ThumbsDown int    `json:"-1"`
	Confused   int    `json:"confused"`
	Eyes       int    `json:"eyes"`
	Heart      int    `json:"heart"`
	Hooray     int    `json:"hooray"`
	Laugh      int    `json:"laugh"`
	Rocket     int    `json:"rocket"`
	TotalCount int    `json:"total_count"`
	URL        string `json:"url"`
}

// StateReason is the reason an issue was closed or reopened.
type StateReason string

// Possible values of [StateReason].
const (
	StateReasonCompleted  StateReason = "completed"
	StateReasonReopened   StateReason = "reopened"
	StateReasonNotPlanned StateReason = "not_planned"
)

var stateReasons = [...]string{
	string(StateReasonCompleted),
	string(StateReasonReopened),
	string(StateReasonNotPlanned),
}

// Values returns the possible values of a [StateReason].
func (StateReason) Values() []string {
	return stateReasons[:]
}

// PRBranch is the representation of a branch.
type PRBranch struct {
	Label string           `json:"label"`
	Ref   string           `json:"ref"`
	Sha   string           `json:"sha"`
	User  SimpleUser       `json:"user"`
	Repo  SimpleRepository `json:"repo"`
}

// SimpleUser is used to keep track of user information in events.
// The GitHub API actually contains more information, but for our
// purposes it redundant.
type SimpleUser struct {
	Login string `json:"login"`
	ID    int64  `json:"id"`
}

// SimpleRepository is a reduced issue struct, keeping track of important
// information but avoiding redundancy.
type SimpleRepository struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	FullName string     `json:"full_name"`
	Owner    SimpleUser `json:"owner"`
}

// SimpleIssue is a reduced issue struct, keeping track of important information
// but avoiding redundancy.
type SimpleIssue struct {
	ID         int64            `json:"id"`
	Number     int64            `json:"number"`
	Repository SimpleRepository `json:"repository"`
}
