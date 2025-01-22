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

func (AuthorAssociation) Values() []string {
	return authorAssociations[:]
}

// TODO: improve
type ReactionRollup map[string]any

// StateReason is the reason an issue was closed or reopened.
type StateReason string

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

func (StateReason) Values() []string {
	return stateReasons[:]
}

type PRBranch struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	// NOTE: there is also repo, user,
	// but we don't implement those for our use case.
}
