// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnoverse/gh-sql/ent/issue"
	"github.com/gnoverse/gh-sql/ent/timelineevent"
	"github.com/gnoverse/gh-sql/ent/user"
	"github.com/gnoverse/gh-sql/pkg/model"
)

// TimelineEventCreate is the builder for creating a TimelineEvent entity.
type TimelineEventCreate struct {
	config
	mutation *TimelineEventMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetURL sets the "url" field.
func (tec *TimelineEventCreate) SetURL(s string) *TimelineEventCreate {
	tec.mutation.SetURL(s)
	return tec
}

// SetEvent sets the "event" field.
func (tec *TimelineEventCreate) SetEvent(s string) *TimelineEventCreate {
	tec.mutation.SetEvent(s)
	return tec
}

// SetCommitID sets the "commit_id" field.
func (tec *TimelineEventCreate) SetCommitID(s string) *TimelineEventCreate {
	tec.mutation.SetCommitID(s)
	return tec
}

// SetNillableCommitID sets the "commit_id" field if the given value is not nil.
func (tec *TimelineEventCreate) SetNillableCommitID(s *string) *TimelineEventCreate {
	if s != nil {
		tec.SetCommitID(*s)
	}
	return tec
}

// SetCommitURL sets the "commit_url" field.
func (tec *TimelineEventCreate) SetCommitURL(s string) *TimelineEventCreate {
	tec.mutation.SetCommitURL(s)
	return tec
}

// SetNillableCommitURL sets the "commit_url" field if the given value is not nil.
func (tec *TimelineEventCreate) SetNillableCommitURL(s *string) *TimelineEventCreate {
	if s != nil {
		tec.SetCommitURL(*s)
	}
	return tec
}

// SetCreatedAt sets the "created_at" field.
func (tec *TimelineEventCreate) SetCreatedAt(t time.Time) *TimelineEventCreate {
	tec.mutation.SetCreatedAt(t)
	return tec
}

// SetData sets the "data" field.
func (tec *TimelineEventCreate) SetData(mew model.TimelineEventWrapper) *TimelineEventCreate {
	tec.mutation.SetData(mew)
	return tec
}

// SetID sets the "id" field.
func (tec *TimelineEventCreate) SetID(s string) *TimelineEventCreate {
	tec.mutation.SetID(s)
	return tec
}

// SetActorID sets the "actor" edge to the User entity by ID.
func (tec *TimelineEventCreate) SetActorID(id int64) *TimelineEventCreate {
	tec.mutation.SetActorID(id)
	return tec
}

// SetNillableActorID sets the "actor" edge to the User entity by ID if the given value is not nil.
func (tec *TimelineEventCreate) SetNillableActorID(id *int64) *TimelineEventCreate {
	if id != nil {
		tec = tec.SetActorID(*id)
	}
	return tec
}

// SetActor sets the "actor" edge to the User entity.
func (tec *TimelineEventCreate) SetActor(u *User) *TimelineEventCreate {
	return tec.SetActorID(u.ID)
}

// SetIssueID sets the "issue" edge to the Issue entity by ID.
func (tec *TimelineEventCreate) SetIssueID(id int64) *TimelineEventCreate {
	tec.mutation.SetIssueID(id)
	return tec
}

// SetNillableIssueID sets the "issue" edge to the Issue entity by ID if the given value is not nil.
func (tec *TimelineEventCreate) SetNillableIssueID(id *int64) *TimelineEventCreate {
	if id != nil {
		tec = tec.SetIssueID(*id)
	}
	return tec
}

// SetIssue sets the "issue" edge to the Issue entity.
func (tec *TimelineEventCreate) SetIssue(i *Issue) *TimelineEventCreate {
	return tec.SetIssueID(i.ID)
}

// Mutation returns the TimelineEventMutation object of the builder.
func (tec *TimelineEventCreate) Mutation() *TimelineEventMutation {
	return tec.mutation
}

// Save creates the TimelineEvent in the database.
func (tec *TimelineEventCreate) Save(ctx context.Context) (*TimelineEvent, error) {
	return withHooks(ctx, tec.sqlSave, tec.mutation, tec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tec *TimelineEventCreate) SaveX(ctx context.Context) *TimelineEvent {
	v, err := tec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tec *TimelineEventCreate) Exec(ctx context.Context) error {
	_, err := tec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tec *TimelineEventCreate) ExecX(ctx context.Context) {
	if err := tec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tec *TimelineEventCreate) check() error {
	if _, ok := tec.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "TimelineEvent.url"`)}
	}
	if _, ok := tec.mutation.Event(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required field "TimelineEvent.event"`)}
	}
	if _, ok := tec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TimelineEvent.created_at"`)}
	}
	if _, ok := tec.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "TimelineEvent.data"`)}
	}
	return nil
}

func (tec *TimelineEventCreate) sqlSave(ctx context.Context) (*TimelineEvent, error) {
	if err := tec.check(); err != nil {
		return nil, err
	}
	_node, _spec := tec.createSpec()
	if err := sqlgraph.CreateNode(ctx, tec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TimelineEvent.ID type: %T", _spec.ID.Value)
		}
	}
	tec.mutation.id = &_node.ID
	tec.mutation.done = true
	return _node, nil
}

func (tec *TimelineEventCreate) createSpec() (*TimelineEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &TimelineEvent{config: tec.config}
		_spec = sqlgraph.NewCreateSpec(timelineevent.Table, sqlgraph.NewFieldSpec(timelineevent.FieldID, field.TypeString))
	)
	_spec.OnConflict = tec.conflict
	if id, ok := tec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tec.mutation.URL(); ok {
		_spec.SetField(timelineevent.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := tec.mutation.Event(); ok {
		_spec.SetField(timelineevent.FieldEvent, field.TypeString, value)
		_node.Event = value
	}
	if value, ok := tec.mutation.CommitID(); ok {
		_spec.SetField(timelineevent.FieldCommitID, field.TypeString, value)
		_node.CommitID = &value
	}
	if value, ok := tec.mutation.CommitURL(); ok {
		_spec.SetField(timelineevent.FieldCommitURL, field.TypeString, value)
		_node.CommitURL = &value
	}
	if value, ok := tec.mutation.CreatedAt(); ok {
		_spec.SetField(timelineevent.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tec.mutation.Data(); ok {
		_spec.SetField(timelineevent.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if nodes := tec.mutation.ActorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timelineevent.ActorTable,
			Columns: []string{timelineevent.ActorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.timeline_event_actor = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tec.mutation.IssueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timelineevent.IssueTable,
			Columns: []string{timelineevent.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.issue_timeline = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TimelineEvent.Create().
//		SetURL(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TimelineEventUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (tec *TimelineEventCreate) OnConflict(opts ...sql.ConflictOption) *TimelineEventUpsertOne {
	tec.conflict = opts
	return &TimelineEventUpsertOne{
		create: tec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TimelineEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tec *TimelineEventCreate) OnConflictColumns(columns ...string) *TimelineEventUpsertOne {
	tec.conflict = append(tec.conflict, sql.ConflictColumns(columns...))
	return &TimelineEventUpsertOne{
		create: tec,
	}
}

type (
	// TimelineEventUpsertOne is the builder for "upsert"-ing
	//  one TimelineEvent node.
	TimelineEventUpsertOne struct {
		create *TimelineEventCreate
	}

	// TimelineEventUpsert is the "OnConflict" setter.
	TimelineEventUpsert struct {
		*sql.UpdateSet
	}
)

// SetURL sets the "url" field.
func (u *TimelineEventUpsert) SetURL(v string) *TimelineEventUpsert {
	u.Set(timelineevent.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *TimelineEventUpsert) UpdateURL() *TimelineEventUpsert {
	u.SetExcluded(timelineevent.FieldURL)
	return u
}

// SetEvent sets the "event" field.
func (u *TimelineEventUpsert) SetEvent(v string) *TimelineEventUpsert {
	u.Set(timelineevent.FieldEvent, v)
	return u
}

// UpdateEvent sets the "event" field to the value that was provided on create.
func (u *TimelineEventUpsert) UpdateEvent() *TimelineEventUpsert {
	u.SetExcluded(timelineevent.FieldEvent)
	return u
}

// SetCommitID sets the "commit_id" field.
func (u *TimelineEventUpsert) SetCommitID(v string) *TimelineEventUpsert {
	u.Set(timelineevent.FieldCommitID, v)
	return u
}

// UpdateCommitID sets the "commit_id" field to the value that was provided on create.
func (u *TimelineEventUpsert) UpdateCommitID() *TimelineEventUpsert {
	u.SetExcluded(timelineevent.FieldCommitID)
	return u
}

// ClearCommitID clears the value of the "commit_id" field.
func (u *TimelineEventUpsert) ClearCommitID() *TimelineEventUpsert {
	u.SetNull(timelineevent.FieldCommitID)
	return u
}

// SetCommitURL sets the "commit_url" field.
func (u *TimelineEventUpsert) SetCommitURL(v string) *TimelineEventUpsert {
	u.Set(timelineevent.FieldCommitURL, v)
	return u
}

// UpdateCommitURL sets the "commit_url" field to the value that was provided on create.
func (u *TimelineEventUpsert) UpdateCommitURL() *TimelineEventUpsert {
	u.SetExcluded(timelineevent.FieldCommitURL)
	return u
}

// ClearCommitURL clears the value of the "commit_url" field.
func (u *TimelineEventUpsert) ClearCommitURL() *TimelineEventUpsert {
	u.SetNull(timelineevent.FieldCommitURL)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TimelineEventUpsert) SetCreatedAt(v time.Time) *TimelineEventUpsert {
	u.Set(timelineevent.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TimelineEventUpsert) UpdateCreatedAt() *TimelineEventUpsert {
	u.SetExcluded(timelineevent.FieldCreatedAt)
	return u
}

// SetData sets the "data" field.
func (u *TimelineEventUpsert) SetData(v model.TimelineEventWrapper) *TimelineEventUpsert {
	u.Set(timelineevent.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *TimelineEventUpsert) UpdateData() *TimelineEventUpsert {
	u.SetExcluded(timelineevent.FieldData)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TimelineEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(timelineevent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TimelineEventUpsertOne) UpdateNewValues() *TimelineEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(timelineevent.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TimelineEvent.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TimelineEventUpsertOne) Ignore() *TimelineEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TimelineEventUpsertOne) DoNothing() *TimelineEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TimelineEventCreate.OnConflict
// documentation for more info.
func (u *TimelineEventUpsertOne) Update(set func(*TimelineEventUpsert)) *TimelineEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TimelineEventUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *TimelineEventUpsertOne) SetURL(v string) *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *TimelineEventUpsertOne) UpdateURL() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateURL()
	})
}

// SetEvent sets the "event" field.
func (u *TimelineEventUpsertOne) SetEvent(v string) *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetEvent(v)
	})
}

// UpdateEvent sets the "event" field to the value that was provided on create.
func (u *TimelineEventUpsertOne) UpdateEvent() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateEvent()
	})
}

// SetCommitID sets the "commit_id" field.
func (u *TimelineEventUpsertOne) SetCommitID(v string) *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetCommitID(v)
	})
}

// UpdateCommitID sets the "commit_id" field to the value that was provided on create.
func (u *TimelineEventUpsertOne) UpdateCommitID() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateCommitID()
	})
}

// ClearCommitID clears the value of the "commit_id" field.
func (u *TimelineEventUpsertOne) ClearCommitID() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.ClearCommitID()
	})
}

// SetCommitURL sets the "commit_url" field.
func (u *TimelineEventUpsertOne) SetCommitURL(v string) *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetCommitURL(v)
	})
}

// UpdateCommitURL sets the "commit_url" field to the value that was provided on create.
func (u *TimelineEventUpsertOne) UpdateCommitURL() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateCommitURL()
	})
}

// ClearCommitURL clears the value of the "commit_url" field.
func (u *TimelineEventUpsertOne) ClearCommitURL() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.ClearCommitURL()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TimelineEventUpsertOne) SetCreatedAt(v time.Time) *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TimelineEventUpsertOne) UpdateCreatedAt() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetData sets the "data" field.
func (u *TimelineEventUpsertOne) SetData(v model.TimelineEventWrapper) *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *TimelineEventUpsertOne) UpdateData() *TimelineEventUpsertOne {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *TimelineEventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TimelineEventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TimelineEventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TimelineEventUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TimelineEventUpsertOne.ID is not supported by MySQL driver. Use TimelineEventUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TimelineEventUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TimelineEventCreateBulk is the builder for creating many TimelineEvent entities in bulk.
type TimelineEventCreateBulk struct {
	config
	err      error
	builders []*TimelineEventCreate
	conflict []sql.ConflictOption
}

// Save creates the TimelineEvent entities in the database.
func (tecb *TimelineEventCreateBulk) Save(ctx context.Context) ([]*TimelineEvent, error) {
	if tecb.err != nil {
		return nil, tecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tecb.builders))
	nodes := make([]*TimelineEvent, len(tecb.builders))
	mutators := make([]Mutator, len(tecb.builders))
	for i := range tecb.builders {
		func(i int, root context.Context) {
			builder := tecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimelineEventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tecb *TimelineEventCreateBulk) SaveX(ctx context.Context) []*TimelineEvent {
	v, err := tecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tecb *TimelineEventCreateBulk) Exec(ctx context.Context) error {
	_, err := tecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tecb *TimelineEventCreateBulk) ExecX(ctx context.Context) {
	if err := tecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TimelineEvent.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TimelineEventUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (tecb *TimelineEventCreateBulk) OnConflict(opts ...sql.ConflictOption) *TimelineEventUpsertBulk {
	tecb.conflict = opts
	return &TimelineEventUpsertBulk{
		create: tecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TimelineEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tecb *TimelineEventCreateBulk) OnConflictColumns(columns ...string) *TimelineEventUpsertBulk {
	tecb.conflict = append(tecb.conflict, sql.ConflictColumns(columns...))
	return &TimelineEventUpsertBulk{
		create: tecb,
	}
}

// TimelineEventUpsertBulk is the builder for "upsert"-ing
// a bulk of TimelineEvent nodes.
type TimelineEventUpsertBulk struct {
	create *TimelineEventCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TimelineEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(timelineevent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TimelineEventUpsertBulk) UpdateNewValues() *TimelineEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(timelineevent.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TimelineEvent.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TimelineEventUpsertBulk) Ignore() *TimelineEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TimelineEventUpsertBulk) DoNothing() *TimelineEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TimelineEventCreateBulk.OnConflict
// documentation for more info.
func (u *TimelineEventUpsertBulk) Update(set func(*TimelineEventUpsert)) *TimelineEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TimelineEventUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *TimelineEventUpsertBulk) SetURL(v string) *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *TimelineEventUpsertBulk) UpdateURL() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateURL()
	})
}

// SetEvent sets the "event" field.
func (u *TimelineEventUpsertBulk) SetEvent(v string) *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetEvent(v)
	})
}

// UpdateEvent sets the "event" field to the value that was provided on create.
func (u *TimelineEventUpsertBulk) UpdateEvent() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateEvent()
	})
}

// SetCommitID sets the "commit_id" field.
func (u *TimelineEventUpsertBulk) SetCommitID(v string) *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetCommitID(v)
	})
}

// UpdateCommitID sets the "commit_id" field to the value that was provided on create.
func (u *TimelineEventUpsertBulk) UpdateCommitID() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateCommitID()
	})
}

// ClearCommitID clears the value of the "commit_id" field.
func (u *TimelineEventUpsertBulk) ClearCommitID() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.ClearCommitID()
	})
}

// SetCommitURL sets the "commit_url" field.
func (u *TimelineEventUpsertBulk) SetCommitURL(v string) *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetCommitURL(v)
	})
}

// UpdateCommitURL sets the "commit_url" field to the value that was provided on create.
func (u *TimelineEventUpsertBulk) UpdateCommitURL() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateCommitURL()
	})
}

// ClearCommitURL clears the value of the "commit_url" field.
func (u *TimelineEventUpsertBulk) ClearCommitURL() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.ClearCommitURL()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TimelineEventUpsertBulk) SetCreatedAt(v time.Time) *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TimelineEventUpsertBulk) UpdateCreatedAt() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetData sets the "data" field.
func (u *TimelineEventUpsertBulk) SetData(v model.TimelineEventWrapper) *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *TimelineEventUpsertBulk) UpdateData() *TimelineEventUpsertBulk {
	return u.Update(func(s *TimelineEventUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *TimelineEventUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TimelineEventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TimelineEventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TimelineEventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
