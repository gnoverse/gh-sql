// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnolang/gh-sql/ent/issue"
	"github.com/gnolang/gh-sql/ent/issuecomment"
	"github.com/gnolang/gh-sql/ent/predicate"
	"github.com/gnolang/gh-sql/ent/user"
	"github.com/gnolang/gh-sql/pkg/model"
)

// IssueCommentUpdate is the builder for updating IssueComment entities.
type IssueCommentUpdate struct {
	config
	hooks    []Hook
	mutation *IssueCommentMutation
}

// Where appends a list predicates to the IssueCommentUpdate builder.
func (icu *IssueCommentUpdate) Where(ps ...predicate.IssueComment) *IssueCommentUpdate {
	icu.mutation.Where(ps...)
	return icu
}

// SetNodeID sets the "node_id" field.
func (icu *IssueCommentUpdate) SetNodeID(s string) *IssueCommentUpdate {
	icu.mutation.SetNodeID(s)
	return icu
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableNodeID(s *string) *IssueCommentUpdate {
	if s != nil {
		icu.SetNodeID(*s)
	}
	return icu
}

// SetURL sets the "url" field.
func (icu *IssueCommentUpdate) SetURL(s string) *IssueCommentUpdate {
	icu.mutation.SetURL(s)
	return icu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableURL(s *string) *IssueCommentUpdate {
	if s != nil {
		icu.SetURL(*s)
	}
	return icu
}

// SetBody sets the "body" field.
func (icu *IssueCommentUpdate) SetBody(s string) *IssueCommentUpdate {
	icu.mutation.SetBody(s)
	return icu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableBody(s *string) *IssueCommentUpdate {
	if s != nil {
		icu.SetBody(*s)
	}
	return icu
}

// SetHTMLURL sets the "html_url" field.
func (icu *IssueCommentUpdate) SetHTMLURL(s string) *IssueCommentUpdate {
	icu.mutation.SetHTMLURL(s)
	return icu
}

// SetNillableHTMLURL sets the "html_url" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableHTMLURL(s *string) *IssueCommentUpdate {
	if s != nil {
		icu.SetHTMLURL(*s)
	}
	return icu
}

// SetCreatedAt sets the "created_at" field.
func (icu *IssueCommentUpdate) SetCreatedAt(s string) *IssueCommentUpdate {
	icu.mutation.SetCreatedAt(s)
	return icu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableCreatedAt(s *string) *IssueCommentUpdate {
	if s != nil {
		icu.SetCreatedAt(*s)
	}
	return icu
}

// SetUpdatedAt sets the "updated_at" field.
func (icu *IssueCommentUpdate) SetUpdatedAt(s string) *IssueCommentUpdate {
	icu.mutation.SetUpdatedAt(s)
	return icu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableUpdatedAt(s *string) *IssueCommentUpdate {
	if s != nil {
		icu.SetUpdatedAt(*s)
	}
	return icu
}

// SetIssueURL sets the "issue_url" field.
func (icu *IssueCommentUpdate) SetIssueURL(s string) *IssueCommentUpdate {
	icu.mutation.SetIssueURL(s)
	return icu
}

// SetNillableIssueURL sets the "issue_url" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableIssueURL(s *string) *IssueCommentUpdate {
	if s != nil {
		icu.SetIssueURL(*s)
	}
	return icu
}

// SetAuthorAssociation sets the "author_association" field.
func (icu *IssueCommentUpdate) SetAuthorAssociation(ma model.AuthorAssociation) *IssueCommentUpdate {
	icu.mutation.SetAuthorAssociation(ma)
	return icu
}

// SetNillableAuthorAssociation sets the "author_association" field if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableAuthorAssociation(ma *model.AuthorAssociation) *IssueCommentUpdate {
	if ma != nil {
		icu.SetAuthorAssociation(*ma)
	}
	return icu
}

// SetReactions sets the "reactions" field.
func (icu *IssueCommentUpdate) SetReactions(m map[string]interface{}) *IssueCommentUpdate {
	icu.mutation.SetReactions(m)
	return icu
}

// SetIssueID sets the "issue" edge to the Issue entity by ID.
func (icu *IssueCommentUpdate) SetIssueID(id int64) *IssueCommentUpdate {
	icu.mutation.SetIssueID(id)
	return icu
}

// SetNillableIssueID sets the "issue" edge to the Issue entity by ID if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableIssueID(id *int64) *IssueCommentUpdate {
	if id != nil {
		icu = icu.SetIssueID(*id)
	}
	return icu
}

// SetIssue sets the "issue" edge to the Issue entity.
func (icu *IssueCommentUpdate) SetIssue(i *Issue) *IssueCommentUpdate {
	return icu.SetIssueID(i.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (icu *IssueCommentUpdate) SetUserID(id int64) *IssueCommentUpdate {
	icu.mutation.SetUserID(id)
	return icu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (icu *IssueCommentUpdate) SetNillableUserID(id *int64) *IssueCommentUpdate {
	if id != nil {
		icu = icu.SetUserID(*id)
	}
	return icu
}

// SetUser sets the "user" edge to the User entity.
func (icu *IssueCommentUpdate) SetUser(u *User) *IssueCommentUpdate {
	return icu.SetUserID(u.ID)
}

// Mutation returns the IssueCommentMutation object of the builder.
func (icu *IssueCommentUpdate) Mutation() *IssueCommentMutation {
	return icu.mutation
}

// ClearIssue clears the "issue" edge to the Issue entity.
func (icu *IssueCommentUpdate) ClearIssue() *IssueCommentUpdate {
	icu.mutation.ClearIssue()
	return icu
}

// ClearUser clears the "user" edge to the User entity.
func (icu *IssueCommentUpdate) ClearUser() *IssueCommentUpdate {
	icu.mutation.ClearUser()
	return icu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (icu *IssueCommentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, icu.sqlSave, icu.mutation, icu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (icu *IssueCommentUpdate) SaveX(ctx context.Context) int {
	affected, err := icu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (icu *IssueCommentUpdate) Exec(ctx context.Context) error {
	_, err := icu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icu *IssueCommentUpdate) ExecX(ctx context.Context) {
	if err := icu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icu *IssueCommentUpdate) check() error {
	if v, ok := icu.mutation.AuthorAssociation(); ok {
		if err := issuecomment.AuthorAssociationValidator(v); err != nil {
			return &ValidationError{Name: "author_association", err: fmt.Errorf(`ent: validator failed for field "IssueComment.author_association": %w`, err)}
		}
	}
	return nil
}

func (icu *IssueCommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := icu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(issuecomment.Table, issuecomment.Columns, sqlgraph.NewFieldSpec(issuecomment.FieldID, field.TypeInt64))
	if ps := icu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icu.mutation.NodeID(); ok {
		_spec.SetField(issuecomment.FieldNodeID, field.TypeString, value)
	}
	if value, ok := icu.mutation.URL(); ok {
		_spec.SetField(issuecomment.FieldURL, field.TypeString, value)
	}
	if value, ok := icu.mutation.Body(); ok {
		_spec.SetField(issuecomment.FieldBody, field.TypeString, value)
	}
	if value, ok := icu.mutation.HTMLURL(); ok {
		_spec.SetField(issuecomment.FieldHTMLURL, field.TypeString, value)
	}
	if value, ok := icu.mutation.CreatedAt(); ok {
		_spec.SetField(issuecomment.FieldCreatedAt, field.TypeString, value)
	}
	if value, ok := icu.mutation.UpdatedAt(); ok {
		_spec.SetField(issuecomment.FieldUpdatedAt, field.TypeString, value)
	}
	if value, ok := icu.mutation.IssueURL(); ok {
		_spec.SetField(issuecomment.FieldIssueURL, field.TypeString, value)
	}
	if value, ok := icu.mutation.AuthorAssociation(); ok {
		_spec.SetField(issuecomment.FieldAuthorAssociation, field.TypeEnum, value)
	}
	if value, ok := icu.mutation.Reactions(); ok {
		_spec.SetField(issuecomment.FieldReactions, field.TypeJSON, value)
	}
	if icu.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.IssueTable,
			Columns: []string{issuecomment.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icu.mutation.IssueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.IssueTable,
			Columns: []string{issuecomment.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if icu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.UserTable,
			Columns: []string{issuecomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.UserTable,
			Columns: []string{issuecomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, icu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issuecomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	icu.mutation.done = true
	return n, nil
}

// IssueCommentUpdateOne is the builder for updating a single IssueComment entity.
type IssueCommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IssueCommentMutation
}

// SetNodeID sets the "node_id" field.
func (icuo *IssueCommentUpdateOne) SetNodeID(s string) *IssueCommentUpdateOne {
	icuo.mutation.SetNodeID(s)
	return icuo
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableNodeID(s *string) *IssueCommentUpdateOne {
	if s != nil {
		icuo.SetNodeID(*s)
	}
	return icuo
}

// SetURL sets the "url" field.
func (icuo *IssueCommentUpdateOne) SetURL(s string) *IssueCommentUpdateOne {
	icuo.mutation.SetURL(s)
	return icuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableURL(s *string) *IssueCommentUpdateOne {
	if s != nil {
		icuo.SetURL(*s)
	}
	return icuo
}

// SetBody sets the "body" field.
func (icuo *IssueCommentUpdateOne) SetBody(s string) *IssueCommentUpdateOne {
	icuo.mutation.SetBody(s)
	return icuo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableBody(s *string) *IssueCommentUpdateOne {
	if s != nil {
		icuo.SetBody(*s)
	}
	return icuo
}

// SetHTMLURL sets the "html_url" field.
func (icuo *IssueCommentUpdateOne) SetHTMLURL(s string) *IssueCommentUpdateOne {
	icuo.mutation.SetHTMLURL(s)
	return icuo
}

// SetNillableHTMLURL sets the "html_url" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableHTMLURL(s *string) *IssueCommentUpdateOne {
	if s != nil {
		icuo.SetHTMLURL(*s)
	}
	return icuo
}

// SetCreatedAt sets the "created_at" field.
func (icuo *IssueCommentUpdateOne) SetCreatedAt(s string) *IssueCommentUpdateOne {
	icuo.mutation.SetCreatedAt(s)
	return icuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableCreatedAt(s *string) *IssueCommentUpdateOne {
	if s != nil {
		icuo.SetCreatedAt(*s)
	}
	return icuo
}

// SetUpdatedAt sets the "updated_at" field.
func (icuo *IssueCommentUpdateOne) SetUpdatedAt(s string) *IssueCommentUpdateOne {
	icuo.mutation.SetUpdatedAt(s)
	return icuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableUpdatedAt(s *string) *IssueCommentUpdateOne {
	if s != nil {
		icuo.SetUpdatedAt(*s)
	}
	return icuo
}

// SetIssueURL sets the "issue_url" field.
func (icuo *IssueCommentUpdateOne) SetIssueURL(s string) *IssueCommentUpdateOne {
	icuo.mutation.SetIssueURL(s)
	return icuo
}

// SetNillableIssueURL sets the "issue_url" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableIssueURL(s *string) *IssueCommentUpdateOne {
	if s != nil {
		icuo.SetIssueURL(*s)
	}
	return icuo
}

// SetAuthorAssociation sets the "author_association" field.
func (icuo *IssueCommentUpdateOne) SetAuthorAssociation(ma model.AuthorAssociation) *IssueCommentUpdateOne {
	icuo.mutation.SetAuthorAssociation(ma)
	return icuo
}

// SetNillableAuthorAssociation sets the "author_association" field if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableAuthorAssociation(ma *model.AuthorAssociation) *IssueCommentUpdateOne {
	if ma != nil {
		icuo.SetAuthorAssociation(*ma)
	}
	return icuo
}

// SetReactions sets the "reactions" field.
func (icuo *IssueCommentUpdateOne) SetReactions(m map[string]interface{}) *IssueCommentUpdateOne {
	icuo.mutation.SetReactions(m)
	return icuo
}

// SetIssueID sets the "issue" edge to the Issue entity by ID.
func (icuo *IssueCommentUpdateOne) SetIssueID(id int64) *IssueCommentUpdateOne {
	icuo.mutation.SetIssueID(id)
	return icuo
}

// SetNillableIssueID sets the "issue" edge to the Issue entity by ID if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableIssueID(id *int64) *IssueCommentUpdateOne {
	if id != nil {
		icuo = icuo.SetIssueID(*id)
	}
	return icuo
}

// SetIssue sets the "issue" edge to the Issue entity.
func (icuo *IssueCommentUpdateOne) SetIssue(i *Issue) *IssueCommentUpdateOne {
	return icuo.SetIssueID(i.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (icuo *IssueCommentUpdateOne) SetUserID(id int64) *IssueCommentUpdateOne {
	icuo.mutation.SetUserID(id)
	return icuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (icuo *IssueCommentUpdateOne) SetNillableUserID(id *int64) *IssueCommentUpdateOne {
	if id != nil {
		icuo = icuo.SetUserID(*id)
	}
	return icuo
}

// SetUser sets the "user" edge to the User entity.
func (icuo *IssueCommentUpdateOne) SetUser(u *User) *IssueCommentUpdateOne {
	return icuo.SetUserID(u.ID)
}

// Mutation returns the IssueCommentMutation object of the builder.
func (icuo *IssueCommentUpdateOne) Mutation() *IssueCommentMutation {
	return icuo.mutation
}

// ClearIssue clears the "issue" edge to the Issue entity.
func (icuo *IssueCommentUpdateOne) ClearIssue() *IssueCommentUpdateOne {
	icuo.mutation.ClearIssue()
	return icuo
}

// ClearUser clears the "user" edge to the User entity.
func (icuo *IssueCommentUpdateOne) ClearUser() *IssueCommentUpdateOne {
	icuo.mutation.ClearUser()
	return icuo
}

// Where appends a list predicates to the IssueCommentUpdate builder.
func (icuo *IssueCommentUpdateOne) Where(ps ...predicate.IssueComment) *IssueCommentUpdateOne {
	icuo.mutation.Where(ps...)
	return icuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (icuo *IssueCommentUpdateOne) Select(field string, fields ...string) *IssueCommentUpdateOne {
	icuo.fields = append([]string{field}, fields...)
	return icuo
}

// Save executes the query and returns the updated IssueComment entity.
func (icuo *IssueCommentUpdateOne) Save(ctx context.Context) (*IssueComment, error) {
	return withHooks(ctx, icuo.sqlSave, icuo.mutation, icuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (icuo *IssueCommentUpdateOne) SaveX(ctx context.Context) *IssueComment {
	node, err := icuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (icuo *IssueCommentUpdateOne) Exec(ctx context.Context) error {
	_, err := icuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icuo *IssueCommentUpdateOne) ExecX(ctx context.Context) {
	if err := icuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icuo *IssueCommentUpdateOne) check() error {
	if v, ok := icuo.mutation.AuthorAssociation(); ok {
		if err := issuecomment.AuthorAssociationValidator(v); err != nil {
			return &ValidationError{Name: "author_association", err: fmt.Errorf(`ent: validator failed for field "IssueComment.author_association": %w`, err)}
		}
	}
	return nil
}

func (icuo *IssueCommentUpdateOne) sqlSave(ctx context.Context) (_node *IssueComment, err error) {
	if err := icuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(issuecomment.Table, issuecomment.Columns, sqlgraph.NewFieldSpec(issuecomment.FieldID, field.TypeInt64))
	id, ok := icuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IssueComment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := icuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, issuecomment.FieldID)
		for _, f := range fields {
			if !issuecomment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != issuecomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := icuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icuo.mutation.NodeID(); ok {
		_spec.SetField(issuecomment.FieldNodeID, field.TypeString, value)
	}
	if value, ok := icuo.mutation.URL(); ok {
		_spec.SetField(issuecomment.FieldURL, field.TypeString, value)
	}
	if value, ok := icuo.mutation.Body(); ok {
		_spec.SetField(issuecomment.FieldBody, field.TypeString, value)
	}
	if value, ok := icuo.mutation.HTMLURL(); ok {
		_spec.SetField(issuecomment.FieldHTMLURL, field.TypeString, value)
	}
	if value, ok := icuo.mutation.CreatedAt(); ok {
		_spec.SetField(issuecomment.FieldCreatedAt, field.TypeString, value)
	}
	if value, ok := icuo.mutation.UpdatedAt(); ok {
		_spec.SetField(issuecomment.FieldUpdatedAt, field.TypeString, value)
	}
	if value, ok := icuo.mutation.IssueURL(); ok {
		_spec.SetField(issuecomment.FieldIssueURL, field.TypeString, value)
	}
	if value, ok := icuo.mutation.AuthorAssociation(); ok {
		_spec.SetField(issuecomment.FieldAuthorAssociation, field.TypeEnum, value)
	}
	if value, ok := icuo.mutation.Reactions(); ok {
		_spec.SetField(issuecomment.FieldReactions, field.TypeJSON, value)
	}
	if icuo.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.IssueTable,
			Columns: []string{issuecomment.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icuo.mutation.IssueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.IssueTable,
			Columns: []string{issuecomment.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if icuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.UserTable,
			Columns: []string{issuecomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issuecomment.UserTable,
			Columns: []string{issuecomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IssueComment{config: icuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, icuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issuecomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	icuo.mutation.done = true
	return _node, nil
}
