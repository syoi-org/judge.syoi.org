// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/syoi-org/judy/ent/predicate"
	"github.com/syoi-org/judy/ent/problem"
	"github.com/syoi-org/judy/ent/submission"
)

// SubmissionQuery is the builder for querying Submission entities.
type SubmissionQuery struct {
	config
	ctx         *QueryContext
	order       []submission.OrderOption
	inters      []Interceptor
	predicates  []predicate.Submission
	withProblem *ProblemQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubmissionQuery builder.
func (sq *SubmissionQuery) Where(ps ...predicate.Submission) *SubmissionQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SubmissionQuery) Limit(limit int) *SubmissionQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SubmissionQuery) Offset(offset int) *SubmissionQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SubmissionQuery) Unique(unique bool) *SubmissionQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SubmissionQuery) Order(o ...submission.OrderOption) *SubmissionQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryProblem chains the current query on the "problem" edge.
func (sq *SubmissionQuery) QueryProblem() *ProblemQuery {
	query := (&ProblemClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(submission.Table, submission.FieldID, selector),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, submission.ProblemTable, submission.ProblemColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Submission entity from the query.
// Returns a *NotFoundError when no Submission was found.
func (sq *SubmissionQuery) First(ctx context.Context) (*Submission, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{submission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubmissionQuery) FirstX(ctx context.Context) *Submission {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Submission ID from the query.
// Returns a *NotFoundError when no Submission ID was found.
func (sq *SubmissionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{submission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SubmissionQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Submission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Submission entity is found.
// Returns a *NotFoundError when no Submission entities are found.
func (sq *SubmissionQuery) Only(ctx context.Context) (*Submission, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{submission.Label}
	default:
		return nil, &NotSingularError{submission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubmissionQuery) OnlyX(ctx context.Context) *Submission {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Submission ID in the query.
// Returns a *NotSingularError when more than one Submission ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SubmissionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{submission.Label}
	default:
		err = &NotSingularError{submission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SubmissionQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Submissions.
func (sq *SubmissionQuery) All(ctx context.Context) ([]*Submission, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Submission, *SubmissionQuery]()
	return withInterceptors[[]*Submission](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubmissionQuery) AllX(ctx context.Context) []*Submission {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Submission IDs.
func (sq *SubmissionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(submission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubmissionQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubmissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SubmissionQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubmissionQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubmissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubmissionQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubmissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubmissionQuery) Clone() *SubmissionQuery {
	if sq == nil {
		return nil
	}
	return &SubmissionQuery{
		config:      sq.config,
		ctx:         sq.ctx.Clone(),
		order:       append([]submission.OrderOption{}, sq.order...),
		inters:      append([]Interceptor{}, sq.inters...),
		predicates:  append([]predicate.Submission{}, sq.predicates...),
		withProblem: sq.withProblem.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithProblem tells the query-builder to eager-load the nodes that are connected to
// the "problem" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubmissionQuery) WithProblem(opts ...func(*ProblemQuery)) *SubmissionQuery {
	query := (&ProblemClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withProblem = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Submission.Query().
//		GroupBy(submission.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SubmissionQuery) GroupBy(field string, fields ...string) *SubmissionGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubmissionGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = submission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Submission.Query().
//		Select(submission.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *SubmissionQuery) Select(fields ...string) *SubmissionSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SubmissionSelect{SubmissionQuery: sq}
	sbuild.label = submission.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubmissionSelect configured with the given aggregations.
func (sq *SubmissionQuery) Aggregate(fns ...AggregateFunc) *SubmissionSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SubmissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !submission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SubmissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Submission, error) {
	var (
		nodes       = []*Submission{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withProblem != nil,
		}
	)
	if sq.withProblem != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, submission.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Submission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Submission{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withProblem; query != nil {
		if err := sq.loadProblem(ctx, query, nodes, nil,
			func(n *Submission, e *Problem) { n.Edges.Problem = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SubmissionQuery) loadProblem(ctx context.Context, query *ProblemQuery, nodes []*Submission, init func(*Submission), assign func(*Submission, *Problem)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Submission)
	for i := range nodes {
		if nodes[i].problem_submissions == nil {
			continue
		}
		fk := *nodes[i].problem_submissions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(problem.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "problem_submissions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SubmissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubmissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(submission.Table, submission.Columns, sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, submission.FieldID)
		for i := range fields {
			if fields[i] != submission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubmissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(submission.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = submission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubmissionGroupBy is the group-by builder for Submission entities.
type SubmissionGroupBy struct {
	selector
	build *SubmissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubmissionGroupBy) Aggregate(fns ...AggregateFunc) *SubmissionGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SubmissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubmissionQuery, *SubmissionGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SubmissionGroupBy) sqlScan(ctx context.Context, root *SubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubmissionSelect is the builder for selecting fields of Submission entities.
type SubmissionSelect struct {
	*SubmissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SubmissionSelect) Aggregate(fns ...AggregateFunc) *SubmissionSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SubmissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubmissionQuery, *SubmissionSelect](ctx, ss.SubmissionQuery, ss, ss.inters, v)
}

func (ss *SubmissionSelect) sqlScan(ctx context.Context, root *SubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
