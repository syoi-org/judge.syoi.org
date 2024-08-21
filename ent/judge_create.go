// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/syoi-org/judy/ent/judge"
	"github.com/syoi-org/judy/ent/problem"
)

// JudgeCreate is the builder for creating a Judge entity.
type JudgeCreate struct {
	config
	mutation *JudgeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (jc *JudgeCreate) SetCreatedAt(t time.Time) *JudgeCreate {
	jc.mutation.SetCreatedAt(t)
	return jc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jc *JudgeCreate) SetNillableCreatedAt(t *time.Time) *JudgeCreate {
	if t != nil {
		jc.SetCreatedAt(*t)
	}
	return jc
}

// SetUpdatedAt sets the "updated_at" field.
func (jc *JudgeCreate) SetUpdatedAt(t time.Time) *JudgeCreate {
	jc.mutation.SetUpdatedAt(t)
	return jc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jc *JudgeCreate) SetNillableUpdatedAt(t *time.Time) *JudgeCreate {
	if t != nil {
		jc.SetUpdatedAt(*t)
	}
	return jc
}

// SetName sets the "name" field.
func (jc *JudgeCreate) SetName(s string) *JudgeCreate {
	jc.mutation.SetName(s)
	return jc
}

// SetCode sets the "code" field.
func (jc *JudgeCreate) SetCode(s string) *JudgeCreate {
	jc.mutation.SetCode(s)
	return jc
}

// SetType sets the "type" field.
func (jc *JudgeCreate) SetType(j judge.Type) *JudgeCreate {
	jc.mutation.SetType(j)
	return jc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (jc *JudgeCreate) SetNillableType(j *judge.Type) *JudgeCreate {
	if j != nil {
		jc.SetType(*j)
	}
	return jc
}

// SetConfiguration sets the "configuration" field.
func (jc *JudgeCreate) SetConfiguration(s string) *JudgeCreate {
	jc.mutation.SetConfiguration(s)
	return jc
}

// AddProblemIDs adds the "problems" edge to the Problem entity by IDs.
func (jc *JudgeCreate) AddProblemIDs(ids ...int) *JudgeCreate {
	jc.mutation.AddProblemIDs(ids...)
	return jc
}

// AddProblems adds the "problems" edges to the Problem entity.
func (jc *JudgeCreate) AddProblems(p ...*Problem) *JudgeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return jc.AddProblemIDs(ids...)
}

// Mutation returns the JudgeMutation object of the builder.
func (jc *JudgeCreate) Mutation() *JudgeMutation {
	return jc.mutation
}

// Save creates the Judge in the database.
func (jc *JudgeCreate) Save(ctx context.Context) (*Judge, error) {
	jc.defaults()
	return withHooks(ctx, jc.sqlSave, jc.mutation, jc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JudgeCreate) SaveX(ctx context.Context) *Judge {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jc *JudgeCreate) Exec(ctx context.Context) error {
	_, err := jc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jc *JudgeCreate) ExecX(ctx context.Context) {
	if err := jc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jc *JudgeCreate) defaults() {
	if _, ok := jc.mutation.CreatedAt(); !ok {
		v := judge.DefaultCreatedAt()
		jc.mutation.SetCreatedAt(v)
	}
	if _, ok := jc.mutation.UpdatedAt(); !ok {
		v := judge.DefaultUpdatedAt()
		jc.mutation.SetUpdatedAt(v)
	}
	if _, ok := jc.mutation.GetType(); !ok {
		v := judge.DefaultType
		jc.mutation.SetType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JudgeCreate) check() error {
	if _, ok := jc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Judge.created_at"`)}
	}
	if _, ok := jc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Judge.updated_at"`)}
	}
	if _, ok := jc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Judge.name"`)}
	}
	if v, ok := jc.mutation.Name(); ok {
		if err := judge.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Judge.name": %w`, err)}
		}
	}
	if _, ok := jc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Judge.code"`)}
	}
	if v, ok := jc.mutation.Code(); ok {
		if err := judge.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Judge.code": %w`, err)}
		}
	}
	if _, ok := jc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Judge.type"`)}
	}
	if v, ok := jc.mutation.GetType(); ok {
		if err := judge.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Judge.type": %w`, err)}
		}
	}
	if _, ok := jc.mutation.Configuration(); !ok {
		return &ValidationError{Name: "configuration", err: errors.New(`ent: missing required field "Judge.configuration"`)}
	}
	return nil
}

func (jc *JudgeCreate) sqlSave(ctx context.Context) (*Judge, error) {
	if err := jc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jc.mutation.id = &_node.ID
	jc.mutation.done = true
	return _node, nil
}

func (jc *JudgeCreate) createSpec() (*Judge, *sqlgraph.CreateSpec) {
	var (
		_node = &Judge{config: jc.config}
		_spec = sqlgraph.NewCreateSpec(judge.Table, sqlgraph.NewFieldSpec(judge.FieldID, field.TypeInt))
	)
	if value, ok := jc.mutation.CreatedAt(); ok {
		_spec.SetField(judge.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jc.mutation.UpdatedAt(); ok {
		_spec.SetField(judge.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := jc.mutation.Name(); ok {
		_spec.SetField(judge.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := jc.mutation.Code(); ok {
		_spec.SetField(judge.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := jc.mutation.GetType(); ok {
		_spec.SetField(judge.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := jc.mutation.Configuration(); ok {
		_spec.SetField(judge.FieldConfiguration, field.TypeString, value)
		_node.Configuration = value
	}
	if nodes := jc.mutation.ProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   judge.ProblemsTable,
			Columns: []string{judge.ProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JudgeCreateBulk is the builder for creating many Judge entities in bulk.
type JudgeCreateBulk struct {
	config
	err      error
	builders []*JudgeCreate
}

// Save creates the Judge entities in the database.
func (jcb *JudgeCreateBulk) Save(ctx context.Context) ([]*Judge, error) {
	if jcb.err != nil {
		return nil, jcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Judge, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JudgeMutation)
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
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jcb *JudgeCreateBulk) SaveX(ctx context.Context) []*Judge {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcb *JudgeCreateBulk) Exec(ctx context.Context) error {
	_, err := jcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcb *JudgeCreateBulk) ExecX(ctx context.Context) {
	if err := jcb.Exec(ctx); err != nil {
		panic(err)
	}
}
