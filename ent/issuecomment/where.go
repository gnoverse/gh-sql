// Code generated by ent, DO NOT EDIT.

package issuecomment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gnoverse/gh-sql/ent/predicate"
	"github.com/gnoverse/gh-sql/pkg/model"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldID, id))
}

// NodeID applies equality check predicate on the "node_id" field. It's identical to NodeIDEQ.
func NodeID(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldNodeID, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldURL, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldBody, v))
}

// HTMLURL applies equality check predicate on the "html_url" field. It's identical to HTMLURLEQ.
func HTMLURL(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldHTMLURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldUpdatedAt, v))
}

// IssueURL applies equality check predicate on the "issue_url" field. It's identical to IssueURLEQ.
func IssueURL(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldIssueURL, v))
}

// NodeIDEQ applies the EQ predicate on the "node_id" field.
func NodeIDEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldNodeID, v))
}

// NodeIDNEQ applies the NEQ predicate on the "node_id" field.
func NodeIDNEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldNodeID, v))
}

// NodeIDIn applies the In predicate on the "node_id" field.
func NodeIDIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldNodeID, vs...))
}

// NodeIDNotIn applies the NotIn predicate on the "node_id" field.
func NodeIDNotIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldNodeID, vs...))
}

// NodeIDGT applies the GT predicate on the "node_id" field.
func NodeIDGT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldNodeID, v))
}

// NodeIDGTE applies the GTE predicate on the "node_id" field.
func NodeIDGTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldNodeID, v))
}

// NodeIDLT applies the LT predicate on the "node_id" field.
func NodeIDLT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldNodeID, v))
}

// NodeIDLTE applies the LTE predicate on the "node_id" field.
func NodeIDLTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldNodeID, v))
}

// NodeIDContains applies the Contains predicate on the "node_id" field.
func NodeIDContains(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContains(FieldNodeID, v))
}

// NodeIDHasPrefix applies the HasPrefix predicate on the "node_id" field.
func NodeIDHasPrefix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasPrefix(FieldNodeID, v))
}

// NodeIDHasSuffix applies the HasSuffix predicate on the "node_id" field.
func NodeIDHasSuffix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasSuffix(FieldNodeID, v))
}

// NodeIDEqualFold applies the EqualFold predicate on the "node_id" field.
func NodeIDEqualFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEqualFold(FieldNodeID, v))
}

// NodeIDContainsFold applies the ContainsFold predicate on the "node_id" field.
func NodeIDContainsFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContainsFold(FieldNodeID, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContainsFold(FieldURL, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContainsFold(FieldBody, v))
}

// HTMLURLEQ applies the EQ predicate on the "html_url" field.
func HTMLURLEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldHTMLURL, v))
}

// HTMLURLNEQ applies the NEQ predicate on the "html_url" field.
func HTMLURLNEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldHTMLURL, v))
}

// HTMLURLIn applies the In predicate on the "html_url" field.
func HTMLURLIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldHTMLURL, vs...))
}

// HTMLURLNotIn applies the NotIn predicate on the "html_url" field.
func HTMLURLNotIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldHTMLURL, vs...))
}

// HTMLURLGT applies the GT predicate on the "html_url" field.
func HTMLURLGT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldHTMLURL, v))
}

// HTMLURLGTE applies the GTE predicate on the "html_url" field.
func HTMLURLGTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldHTMLURL, v))
}

// HTMLURLLT applies the LT predicate on the "html_url" field.
func HTMLURLLT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldHTMLURL, v))
}

// HTMLURLLTE applies the LTE predicate on the "html_url" field.
func HTMLURLLTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldHTMLURL, v))
}

// HTMLURLContains applies the Contains predicate on the "html_url" field.
func HTMLURLContains(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContains(FieldHTMLURL, v))
}

// HTMLURLHasPrefix applies the HasPrefix predicate on the "html_url" field.
func HTMLURLHasPrefix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasPrefix(FieldHTMLURL, v))
}

// HTMLURLHasSuffix applies the HasSuffix predicate on the "html_url" field.
func HTMLURLHasSuffix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasSuffix(FieldHTMLURL, v))
}

// HTMLURLEqualFold applies the EqualFold predicate on the "html_url" field.
func HTMLURLEqualFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEqualFold(FieldHTMLURL, v))
}

// HTMLURLContainsFold applies the ContainsFold predicate on the "html_url" field.
func HTMLURLContainsFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContainsFold(FieldHTMLURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtContains applies the Contains predicate on the "created_at" field.
func CreatedAtContains(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContains(FieldCreatedAt, v))
}

// CreatedAtHasPrefix applies the HasPrefix predicate on the "created_at" field.
func CreatedAtHasPrefix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasPrefix(FieldCreatedAt, v))
}

// CreatedAtHasSuffix applies the HasSuffix predicate on the "created_at" field.
func CreatedAtHasSuffix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasSuffix(FieldCreatedAt, v))
}

// CreatedAtEqualFold applies the EqualFold predicate on the "created_at" field.
func CreatedAtEqualFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEqualFold(FieldCreatedAt, v))
}

// CreatedAtContainsFold applies the ContainsFold predicate on the "created_at" field.
func CreatedAtContainsFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContainsFold(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtContains applies the Contains predicate on the "updated_at" field.
func UpdatedAtContains(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContains(FieldUpdatedAt, v))
}

// UpdatedAtHasPrefix applies the HasPrefix predicate on the "updated_at" field.
func UpdatedAtHasPrefix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasPrefix(FieldUpdatedAt, v))
}

// UpdatedAtHasSuffix applies the HasSuffix predicate on the "updated_at" field.
func UpdatedAtHasSuffix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasSuffix(FieldUpdatedAt, v))
}

// UpdatedAtEqualFold applies the EqualFold predicate on the "updated_at" field.
func UpdatedAtEqualFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEqualFold(FieldUpdatedAt, v))
}

// UpdatedAtContainsFold applies the ContainsFold predicate on the "updated_at" field.
func UpdatedAtContainsFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContainsFold(FieldUpdatedAt, v))
}

// IssueURLEQ applies the EQ predicate on the "issue_url" field.
func IssueURLEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEQ(FieldIssueURL, v))
}

// IssueURLNEQ applies the NEQ predicate on the "issue_url" field.
func IssueURLNEQ(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNEQ(FieldIssueURL, v))
}

// IssueURLIn applies the In predicate on the "issue_url" field.
func IssueURLIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldIn(FieldIssueURL, vs...))
}

// IssueURLNotIn applies the NotIn predicate on the "issue_url" field.
func IssueURLNotIn(vs ...string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldNotIn(FieldIssueURL, vs...))
}

// IssueURLGT applies the GT predicate on the "issue_url" field.
func IssueURLGT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGT(FieldIssueURL, v))
}

// IssueURLGTE applies the GTE predicate on the "issue_url" field.
func IssueURLGTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldGTE(FieldIssueURL, v))
}

// IssueURLLT applies the LT predicate on the "issue_url" field.
func IssueURLLT(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLT(FieldIssueURL, v))
}

// IssueURLLTE applies the LTE predicate on the "issue_url" field.
func IssueURLLTE(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldLTE(FieldIssueURL, v))
}

// IssueURLContains applies the Contains predicate on the "issue_url" field.
func IssueURLContains(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContains(FieldIssueURL, v))
}

// IssueURLHasPrefix applies the HasPrefix predicate on the "issue_url" field.
func IssueURLHasPrefix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasPrefix(FieldIssueURL, v))
}

// IssueURLHasSuffix applies the HasSuffix predicate on the "issue_url" field.
func IssueURLHasSuffix(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldHasSuffix(FieldIssueURL, v))
}

// IssueURLEqualFold applies the EqualFold predicate on the "issue_url" field.
func IssueURLEqualFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldEqualFold(FieldIssueURL, v))
}

// IssueURLContainsFold applies the ContainsFold predicate on the "issue_url" field.
func IssueURLContainsFold(v string) predicate.IssueComment {
	return predicate.IssueComment(sql.FieldContainsFold(FieldIssueURL, v))
}

// AuthorAssociationEQ applies the EQ predicate on the "author_association" field.
func AuthorAssociationEQ(v model.AuthorAssociation) predicate.IssueComment {
	vc := v
	return predicate.IssueComment(sql.FieldEQ(FieldAuthorAssociation, vc))
}

// AuthorAssociationNEQ applies the NEQ predicate on the "author_association" field.
func AuthorAssociationNEQ(v model.AuthorAssociation) predicate.IssueComment {
	vc := v
	return predicate.IssueComment(sql.FieldNEQ(FieldAuthorAssociation, vc))
}

// AuthorAssociationIn applies the In predicate on the "author_association" field.
func AuthorAssociationIn(vs ...model.AuthorAssociation) predicate.IssueComment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IssueComment(sql.FieldIn(FieldAuthorAssociation, v...))
}

// AuthorAssociationNotIn applies the NotIn predicate on the "author_association" field.
func AuthorAssociationNotIn(vs ...model.AuthorAssociation) predicate.IssueComment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IssueComment(sql.FieldNotIn(FieldAuthorAssociation, v...))
}

// HasIssue applies the HasEdge predicate on the "issue" edge.
func HasIssue() predicate.IssueComment {
	return predicate.IssueComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, IssueTable, IssueColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIssueWith applies the HasEdge predicate on the "issue" edge with a given conditions (other predicates).
func HasIssueWith(preds ...predicate.Issue) predicate.IssueComment {
	return predicate.IssueComment(func(s *sql.Selector) {
		step := newIssueStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.IssueComment {
	return predicate.IssueComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.IssueComment {
	return predicate.IssueComment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IssueComment) predicate.IssueComment {
	return predicate.IssueComment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IssueComment) predicate.IssueComment {
	return predicate.IssueComment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.IssueComment) predicate.IssueComment {
	return predicate.IssueComment(sql.NotPredicates(p))
}
