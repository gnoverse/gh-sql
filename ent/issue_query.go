// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnolang/gh-sql/ent/issue"
	"github.com/gnolang/gh-sql/ent/issuecomment"
	"github.com/gnolang/gh-sql/ent/predicate"
	"github.com/gnolang/gh-sql/ent/repository"
	"github.com/gnolang/gh-sql/ent/timelineevent"
	"github.com/gnolang/gh-sql/ent/user"
)

// IssueQuery is the builder for querying Issue entities.
type IssueQuery struct {
	config
	ctx            *QueryContext
	order          []issue.OrderOption
	inters         []Interceptor
	predicates     []predicate.Issue
	withRepository *RepositoryQuery
	withUser       *UserQuery
	withAssignees  *UserQuery
	withComments   *IssueCommentQuery
	withTimeline   *TimelineEventQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IssueQuery builder.
func (iq *IssueQuery) Where(ps ...predicate.Issue) *IssueQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *IssueQuery) Limit(limit int) *IssueQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *IssueQuery) Offset(offset int) *IssueQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IssueQuery) Unique(unique bool) *IssueQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *IssueQuery) Order(o ...issue.OrderOption) *IssueQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryRepository chains the current query on the "repository" edge.
func (iq *IssueQuery) QueryRepository() *RepositoryQuery {
	query := (&RepositoryClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(repository.Table, repository.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, issue.RepositoryTable, issue.RepositoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (iq *IssueQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, issue.UserTable, issue.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssignees chains the current query on the "assignees" edge.
func (iq *IssueQuery) QueryAssignees() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, issue.AssigneesTable, issue.AssigneesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComments chains the current query on the "comments" edge.
func (iq *IssueQuery) QueryComments() *IssueCommentQuery {
	query := (&IssueCommentClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(issuecomment.Table, issuecomment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, issue.CommentsTable, issue.CommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTimeline chains the current query on the "timeline" edge.
func (iq *IssueQuery) QueryTimeline() *TimelineEventQuery {
	query := (&TimelineEventClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(timelineevent.Table, timelineevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, issue.TimelineTable, issue.TimelineColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Issue entity from the query.
// Returns a *NotFoundError when no Issue was found.
func (iq *IssueQuery) First(ctx context.Context) (*Issue, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{issue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IssueQuery) FirstX(ctx context.Context) *Issue {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Issue ID from the query.
// Returns a *NotFoundError when no Issue ID was found.
func (iq *IssueQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{issue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IssueQuery) FirstIDX(ctx context.Context) int64 {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Issue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Issue entity is found.
// Returns a *NotFoundError when no Issue entities are found.
func (iq *IssueQuery) Only(ctx context.Context) (*Issue, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{issue.Label}
	default:
		return nil, &NotSingularError{issue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IssueQuery) OnlyX(ctx context.Context) *Issue {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Issue ID in the query.
// Returns a *NotSingularError when more than one Issue ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *IssueQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{issue.Label}
	default:
		err = &NotSingularError{issue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IssueQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Issues.
func (iq *IssueQuery) All(ctx context.Context) ([]*Issue, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Issue, *IssueQuery]()
	return withInterceptors[[]*Issue](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *IssueQuery) AllX(ctx context.Context) []*Issue {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Issue IDs.
func (iq *IssueQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(issue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IssueQuery) IDsX(ctx context.Context) []int64 {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IssueQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*IssueQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IssueQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IssueQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IssueQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IssueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IssueQuery) Clone() *IssueQuery {
	if iq == nil {
		return nil
	}
	return &IssueQuery{
		config:         iq.config,
		ctx:            iq.ctx.Clone(),
		order:          append([]issue.OrderOption{}, iq.order...),
		inters:         append([]Interceptor{}, iq.inters...),
		predicates:     append([]predicate.Issue{}, iq.predicates...),
		withRepository: iq.withRepository.Clone(),
		withUser:       iq.withUser.Clone(),
		withAssignees:  iq.withAssignees.Clone(),
		withComments:   iq.withComments.Clone(),
		withTimeline:   iq.withTimeline.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithRepository tells the query-builder to eager-load the nodes that are connected to
// the "repository" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithRepository(opts ...func(*RepositoryQuery)) *IssueQuery {
	query := (&RepositoryClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withRepository = query
	return iq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithUser(opts ...func(*UserQuery)) *IssueQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withUser = query
	return iq
}

// WithAssignees tells the query-builder to eager-load the nodes that are connected to
// the "assignees" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithAssignees(opts ...func(*UserQuery)) *IssueQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withAssignees = query
	return iq
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithComments(opts ...func(*IssueCommentQuery)) *IssueQuery {
	query := (&IssueCommentClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withComments = query
	return iq
}

// WithTimeline tells the query-builder to eager-load the nodes that are connected to
// the "timeline" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithTimeline(opts ...func(*TimelineEventQuery)) *IssueQuery {
	query := (&TimelineEventClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withTimeline = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NodeID string `json:"node_id"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Issue.Query().
//		GroupBy(issue.FieldNodeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *IssueQuery) GroupBy(field string, fields ...string) *IssueGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IssueGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = issue.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NodeID string `json:"node_id"`
//	}
//
//	client.Issue.Query().
//		Select(issue.FieldNodeID).
//		Scan(ctx, &v)
func (iq *IssueQuery) Select(fields ...string) *IssueSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &IssueSelect{IssueQuery: iq}
	sbuild.label = issue.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IssueSelect configured with the given aggregations.
func (iq *IssueQuery) Aggregate(fns ...AggregateFunc) *IssueSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *IssueQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !issue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *IssueQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Issue, error) {
	var (
		nodes       = []*Issue{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [5]bool{
			iq.withRepository != nil,
			iq.withUser != nil,
			iq.withAssignees != nil,
			iq.withComments != nil,
			iq.withTimeline != nil,
		}
	)
	if iq.withRepository != nil || iq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, issue.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Issue).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Issue{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withRepository; query != nil {
		if err := iq.loadRepository(ctx, query, nodes, nil,
			func(n *Issue, e *Repository) { n.Edges.Repository = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withUser; query != nil {
		if err := iq.loadUser(ctx, query, nodes, nil,
			func(n *Issue, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withAssignees; query != nil {
		if err := iq.loadAssignees(ctx, query, nodes,
			func(n *Issue) { n.Edges.Assignees = []*User{} },
			func(n *Issue, e *User) { n.Edges.Assignees = append(n.Edges.Assignees, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withComments; query != nil {
		if err := iq.loadComments(ctx, query, nodes,
			func(n *Issue) { n.Edges.Comments = []*IssueComment{} },
			func(n *Issue, e *IssueComment) { n.Edges.Comments = append(n.Edges.Comments, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withTimeline; query != nil {
		if err := iq.loadTimeline(ctx, query, nodes,
			func(n *Issue) { n.Edges.Timeline = []*TimelineEvent{} },
			func(n *Issue, e *TimelineEvent) { n.Edges.Timeline = append(n.Edges.Timeline, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *IssueQuery) loadRepository(ctx context.Context, query *RepositoryQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *Repository)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Issue)
	for i := range nodes {
		if nodes[i].repository_issues == nil {
			continue
		}
		fk := *nodes[i].repository_issues
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(repository.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "repository_issues" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *IssueQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Issue)
	for i := range nodes {
		if nodes[i].user_issues_created == nil {
			continue
		}
		fk := *nodes[i].user_issues_created
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_issues_created" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *IssueQuery) loadAssignees(ctx context.Context, query *UserQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Issue)
	nids := make(map[int64]map[*Issue]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(issue.AssigneesTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(issue.AssigneesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(issue.AssigneesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(issue.AssigneesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*Issue]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "assignees" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (iq *IssueQuery) loadComments(ctx context.Context, query *IssueCommentQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *IssueComment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Issue)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.IssueComment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(issue.CommentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.issue_comments
		if fk == nil {
			return fmt.Errorf(`foreign-key "issue_comments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "issue_comments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *IssueQuery) loadTimeline(ctx context.Context, query *TimelineEventQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *TimelineEvent)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Issue)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TimelineEvent(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(issue.TimelineColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.issue_timeline
		if fk == nil {
			return fmt.Errorf(`foreign-key "issue_timeline" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "issue_timeline" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *IssueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IssueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(issue.Table, issue.Columns, sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt64))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, issue.FieldID)
		for i := range fields {
			if fields[i] != issue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IssueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(issue.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = issue.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IssueGroupBy is the group-by builder for Issue entities.
type IssueGroupBy struct {
	selector
	build *IssueQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IssueGroupBy) Aggregate(fns ...AggregateFunc) *IssueGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *IssueGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IssueQuery, *IssueGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *IssueGroupBy) sqlScan(ctx context.Context, root *IssueQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IssueSelect is the builder for selecting fields of Issue entities.
type IssueSelect struct {
	*IssueQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *IssueSelect) Aggregate(fns ...AggregateFunc) *IssueSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *IssueSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IssueQuery, *IssueSelect](ctx, is.IssueQuery, is, is.inters, v)
}

func (is *IssueSelect) sqlScan(ctx context.Context, root *IssueQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
