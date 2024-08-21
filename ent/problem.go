// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/syoi-org/judy/ent/judge"
	"github.com/syoi-org/judy/ent/problem"
)

// Problem is the model entity for the Problem schema.
type Problem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProblemQuery when eager-loading is set.
	Edges          ProblemEdges `json:"edges"`
	judge_problems *int
	selectValues   sql.SelectValues
}

// ProblemEdges holds the relations/edges for other nodes in the graph.
type ProblemEdges struct {
	// Submissions holds the value of the submissions edge.
	Submissions []*Submission `json:"submissions,omitempty"`
	// Judge holds the value of the judge edge.
	Judge *Judge `json:"judge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SubmissionsOrErr returns the Submissions value or an error if the edge
// was not loaded in eager-loading.
func (e ProblemEdges) SubmissionsOrErr() ([]*Submission, error) {
	if e.loadedTypes[0] {
		return e.Submissions, nil
	}
	return nil, &NotLoadedError{edge: "submissions"}
}

// JudgeOrErr returns the Judge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProblemEdges) JudgeOrErr() (*Judge, error) {
	if e.Judge != nil {
		return e.Judge, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: judge.Label}
	}
	return nil, &NotLoadedError{edge: "judge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Problem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case problem.FieldID:
			values[i] = new(sql.NullInt64)
		case problem.FieldName, problem.FieldCode:
			values[i] = new(sql.NullString)
		case problem.FieldCreatedAt, problem.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case problem.ForeignKeys[0]: // judge_problems
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Problem fields.
func (pr *Problem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case problem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case problem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case problem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case problem.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case problem.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				pr.Code = value.String
			}
		case problem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field judge_problems", value)
			} else if value.Valid {
				pr.judge_problems = new(int)
				*pr.judge_problems = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Problem.
// This includes values selected through modifiers, order, etc.
func (pr *Problem) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QuerySubmissions queries the "submissions" edge of the Problem entity.
func (pr *Problem) QuerySubmissions() *SubmissionQuery {
	return NewProblemClient(pr.config).QuerySubmissions(pr)
}

// QueryJudge queries the "judge" edge of the Problem entity.
func (pr *Problem) QueryJudge() *JudgeQuery {
	return NewProblemClient(pr.config).QueryJudge(pr)
}

// Update returns a builder for updating this Problem.
// Note that you need to call Problem.Unwrap() before calling this method if this Problem
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Problem) Update() *ProblemUpdateOne {
	return NewProblemClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Problem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Problem) Unwrap() *Problem {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Problem is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Problem) String() string {
	var builder strings.Builder
	builder.WriteString("Problem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(pr.Code)
	builder.WriteByte(')')
	return builder.String()
}

// Problems is a parsable slice of Problem.
type Problems []*Problem
