// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/syoi-org/judy/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/syoi-org/judy/ent/judge"
	"github.com/syoi-org/judy/ent/problem"
	"github.com/syoi-org/judy/ent/submission"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Judge is the client for interacting with the Judge builders.
	Judge *JudgeClient
	// Problem is the client for interacting with the Problem builders.
	Problem *ProblemClient
	// Submission is the client for interacting with the Submission builders.
	Submission *SubmissionClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Judge = NewJudgeClient(c.config)
	c.Problem = NewProblemClient(c.config)
	c.Submission = NewSubmissionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Judge:      NewJudgeClient(cfg),
		Problem:    NewProblemClient(cfg),
		Submission: NewSubmissionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Judge:      NewJudgeClient(cfg),
		Problem:    NewProblemClient(cfg),
		Submission: NewSubmissionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Judge.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Judge.Use(hooks...)
	c.Problem.Use(hooks...)
	c.Submission.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Judge.Intercept(interceptors...)
	c.Problem.Intercept(interceptors...)
	c.Submission.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *JudgeMutation:
		return c.Judge.mutate(ctx, m)
	case *ProblemMutation:
		return c.Problem.mutate(ctx, m)
	case *SubmissionMutation:
		return c.Submission.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// JudgeClient is a client for the Judge schema.
type JudgeClient struct {
	config
}

// NewJudgeClient returns a client for the Judge from the given config.
func NewJudgeClient(c config) *JudgeClient {
	return &JudgeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `judge.Hooks(f(g(h())))`.
func (c *JudgeClient) Use(hooks ...Hook) {
	c.hooks.Judge = append(c.hooks.Judge, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `judge.Intercept(f(g(h())))`.
func (c *JudgeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Judge = append(c.inters.Judge, interceptors...)
}

// Create returns a builder for creating a Judge entity.
func (c *JudgeClient) Create() *JudgeCreate {
	mutation := newJudgeMutation(c.config, OpCreate)
	return &JudgeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Judge entities.
func (c *JudgeClient) CreateBulk(builders ...*JudgeCreate) *JudgeCreateBulk {
	return &JudgeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *JudgeClient) MapCreateBulk(slice any, setFunc func(*JudgeCreate, int)) *JudgeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &JudgeCreateBulk{err: fmt.Errorf("calling to JudgeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*JudgeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &JudgeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Judge.
func (c *JudgeClient) Update() *JudgeUpdate {
	mutation := newJudgeMutation(c.config, OpUpdate)
	return &JudgeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *JudgeClient) UpdateOne(j *Judge) *JudgeUpdateOne {
	mutation := newJudgeMutation(c.config, OpUpdateOne, withJudge(j))
	return &JudgeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *JudgeClient) UpdateOneID(id int) *JudgeUpdateOne {
	mutation := newJudgeMutation(c.config, OpUpdateOne, withJudgeID(id))
	return &JudgeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Judge.
func (c *JudgeClient) Delete() *JudgeDelete {
	mutation := newJudgeMutation(c.config, OpDelete)
	return &JudgeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *JudgeClient) DeleteOne(j *Judge) *JudgeDeleteOne {
	return c.DeleteOneID(j.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *JudgeClient) DeleteOneID(id int) *JudgeDeleteOne {
	builder := c.Delete().Where(judge.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &JudgeDeleteOne{builder}
}

// Query returns a query builder for Judge.
func (c *JudgeClient) Query() *JudgeQuery {
	return &JudgeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeJudge},
		inters: c.Interceptors(),
	}
}

// Get returns a Judge entity by its id.
func (c *JudgeClient) Get(ctx context.Context, id int) (*Judge, error) {
	return c.Query().Where(judge.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JudgeClient) GetX(ctx context.Context, id int) *Judge {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProblems queries the problems edge of a Judge.
func (c *JudgeClient) QueryProblems(j *Judge) *ProblemQuery {
	query := (&ProblemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := j.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(judge.Table, judge.FieldID, id),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, judge.ProblemsTable, judge.ProblemsColumn),
		)
		fromV = sqlgraph.Neighbors(j.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *JudgeClient) Hooks() []Hook {
	return c.hooks.Judge
}

// Interceptors returns the client interceptors.
func (c *JudgeClient) Interceptors() []Interceptor {
	return c.inters.Judge
}

func (c *JudgeClient) mutate(ctx context.Context, m *JudgeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&JudgeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&JudgeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&JudgeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&JudgeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Judge mutation op: %q", m.Op())
	}
}

// ProblemClient is a client for the Problem schema.
type ProblemClient struct {
	config
}

// NewProblemClient returns a client for the Problem from the given config.
func NewProblemClient(c config) *ProblemClient {
	return &ProblemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `problem.Hooks(f(g(h())))`.
func (c *ProblemClient) Use(hooks ...Hook) {
	c.hooks.Problem = append(c.hooks.Problem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `problem.Intercept(f(g(h())))`.
func (c *ProblemClient) Intercept(interceptors ...Interceptor) {
	c.inters.Problem = append(c.inters.Problem, interceptors...)
}

// Create returns a builder for creating a Problem entity.
func (c *ProblemClient) Create() *ProblemCreate {
	mutation := newProblemMutation(c.config, OpCreate)
	return &ProblemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Problem entities.
func (c *ProblemClient) CreateBulk(builders ...*ProblemCreate) *ProblemCreateBulk {
	return &ProblemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ProblemClient) MapCreateBulk(slice any, setFunc func(*ProblemCreate, int)) *ProblemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ProblemCreateBulk{err: fmt.Errorf("calling to ProblemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ProblemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ProblemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Problem.
func (c *ProblemClient) Update() *ProblemUpdate {
	mutation := newProblemMutation(c.config, OpUpdate)
	return &ProblemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProblemClient) UpdateOne(pr *Problem) *ProblemUpdateOne {
	mutation := newProblemMutation(c.config, OpUpdateOne, withProblem(pr))
	return &ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProblemClient) UpdateOneID(id int) *ProblemUpdateOne {
	mutation := newProblemMutation(c.config, OpUpdateOne, withProblemID(id))
	return &ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Problem.
func (c *ProblemClient) Delete() *ProblemDelete {
	mutation := newProblemMutation(c.config, OpDelete)
	return &ProblemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProblemClient) DeleteOne(pr *Problem) *ProblemDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProblemClient) DeleteOneID(id int) *ProblemDeleteOne {
	builder := c.Delete().Where(problem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProblemDeleteOne{builder}
}

// Query returns a query builder for Problem.
func (c *ProblemClient) Query() *ProblemQuery {
	return &ProblemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProblem},
		inters: c.Interceptors(),
	}
}

// Get returns a Problem entity by its id.
func (c *ProblemClient) Get(ctx context.Context, id int) (*Problem, error) {
	return c.Query().Where(problem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProblemClient) GetX(ctx context.Context, id int) *Problem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySubmissions queries the submissions edge of a Problem.
func (c *ProblemClient) QuerySubmissions(pr *Problem) *SubmissionQuery {
	query := (&SubmissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, id),
			sqlgraph.To(submission.Table, submission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, problem.SubmissionsTable, problem.SubmissionsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryJudge queries the judge edge of a Problem.
func (c *ProblemClient) QueryJudge(pr *Problem) *JudgeQuery {
	query := (&JudgeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, id),
			sqlgraph.To(judge.Table, judge.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.JudgeTable, problem.JudgeColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProblemClient) Hooks() []Hook {
	return c.hooks.Problem
}

// Interceptors returns the client interceptors.
func (c *ProblemClient) Interceptors() []Interceptor {
	return c.inters.Problem
}

func (c *ProblemClient) mutate(ctx context.Context, m *ProblemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProblemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProblemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProblemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Problem mutation op: %q", m.Op())
	}
}

// SubmissionClient is a client for the Submission schema.
type SubmissionClient struct {
	config
}

// NewSubmissionClient returns a client for the Submission from the given config.
func NewSubmissionClient(c config) *SubmissionClient {
	return &SubmissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `submission.Hooks(f(g(h())))`.
func (c *SubmissionClient) Use(hooks ...Hook) {
	c.hooks.Submission = append(c.hooks.Submission, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `submission.Intercept(f(g(h())))`.
func (c *SubmissionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Submission = append(c.inters.Submission, interceptors...)
}

// Create returns a builder for creating a Submission entity.
func (c *SubmissionClient) Create() *SubmissionCreate {
	mutation := newSubmissionMutation(c.config, OpCreate)
	return &SubmissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Submission entities.
func (c *SubmissionClient) CreateBulk(builders ...*SubmissionCreate) *SubmissionCreateBulk {
	return &SubmissionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SubmissionClient) MapCreateBulk(slice any, setFunc func(*SubmissionCreate, int)) *SubmissionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SubmissionCreateBulk{err: fmt.Errorf("calling to SubmissionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SubmissionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SubmissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Submission.
func (c *SubmissionClient) Update() *SubmissionUpdate {
	mutation := newSubmissionMutation(c.config, OpUpdate)
	return &SubmissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubmissionClient) UpdateOne(s *Submission) *SubmissionUpdateOne {
	mutation := newSubmissionMutation(c.config, OpUpdateOne, withSubmission(s))
	return &SubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubmissionClient) UpdateOneID(id int) *SubmissionUpdateOne {
	mutation := newSubmissionMutation(c.config, OpUpdateOne, withSubmissionID(id))
	return &SubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Submission.
func (c *SubmissionClient) Delete() *SubmissionDelete {
	mutation := newSubmissionMutation(c.config, OpDelete)
	return &SubmissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SubmissionClient) DeleteOne(s *Submission) *SubmissionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SubmissionClient) DeleteOneID(id int) *SubmissionDeleteOne {
	builder := c.Delete().Where(submission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubmissionDeleteOne{builder}
}

// Query returns a query builder for Submission.
func (c *SubmissionClient) Query() *SubmissionQuery {
	return &SubmissionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSubmission},
		inters: c.Interceptors(),
	}
}

// Get returns a Submission entity by its id.
func (c *SubmissionClient) Get(ctx context.Context, id int) (*Submission, error) {
	return c.Query().Where(submission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubmissionClient) GetX(ctx context.Context, id int) *Submission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProblem queries the problem edge of a Submission.
func (c *SubmissionClient) QueryProblem(s *Submission) *ProblemQuery {
	query := (&ProblemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(submission.Table, submission.FieldID, id),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, submission.ProblemTable, submission.ProblemColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubmissionClient) Hooks() []Hook {
	return c.hooks.Submission
}

// Interceptors returns the client interceptors.
func (c *SubmissionClient) Interceptors() []Interceptor {
	return c.inters.Submission
}

func (c *SubmissionClient) mutate(ctx context.Context, m *SubmissionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SubmissionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SubmissionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SubmissionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Submission mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Judge, Problem, Submission []ent.Hook
	}
	inters struct {
		Judge, Problem, Submission []ent.Interceptor
	}
)
