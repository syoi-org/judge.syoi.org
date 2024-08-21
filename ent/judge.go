// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/syoi-org/judy/ent/judge"
)

// Judge is the model entity for the Judge schema.
type Judge struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Type holds the value of the "type" field.
	Type judge.Type `json:"type,omitempty"`
	// Configuration holds the value of the "configuration" field.
	Configuration string `json:"configuration,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JudgeQuery when eager-loading is set.
	Edges        JudgeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// JudgeEdges holds the relations/edges for other nodes in the graph.
type JudgeEdges struct {
	// Problems holds the value of the problems edge.
	Problems []*Problem `json:"problems,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProblemsOrErr returns the Problems value or an error if the edge
// was not loaded in eager-loading.
func (e JudgeEdges) ProblemsOrErr() ([]*Problem, error) {
	if e.loadedTypes[0] {
		return e.Problems, nil
	}
	return nil, &NotLoadedError{edge: "problems"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Judge) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case judge.FieldID:
			values[i] = new(sql.NullInt64)
		case judge.FieldName, judge.FieldCode, judge.FieldType, judge.FieldConfiguration:
			values[i] = new(sql.NullString)
		case judge.FieldCreatedAt, judge.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Judge fields.
func (j *Judge) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case judge.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			j.ID = int(value.Int64)
		case judge.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				j.Name = value.String
			}
		case judge.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				j.Code = value.String
			}
		case judge.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				j.Type = judge.Type(value.String)
			}
		case judge.FieldConfiguration:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field configuration", values[i])
			} else if value.Valid {
				j.Configuration = value.String
			}
		case judge.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				j.CreatedAt = value.Time
			}
		case judge.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				j.UpdatedAt = value.Time
			}
		default:
			j.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Judge.
// This includes values selected through modifiers, order, etc.
func (j *Judge) Value(name string) (ent.Value, error) {
	return j.selectValues.Get(name)
}

// QueryProblems queries the "problems" edge of the Judge entity.
func (j *Judge) QueryProblems() *ProblemQuery {
	return NewJudgeClient(j.config).QueryProblems(j)
}

// Update returns a builder for updating this Judge.
// Note that you need to call Judge.Unwrap() before calling this method if this Judge
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Judge) Update() *JudgeUpdateOne {
	return NewJudgeClient(j.config).UpdateOne(j)
}

// Unwrap unwraps the Judge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Judge) Unwrap() *Judge {
	_tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Judge is not a transactional entity")
	}
	j.config.driver = _tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Judge) String() string {
	var builder strings.Builder
	builder.WriteString("Judge(")
	builder.WriteString(fmt.Sprintf("id=%v, ", j.ID))
	builder.WriteString("name=")
	builder.WriteString(j.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(j.Code)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", j.Type))
	builder.WriteString(", ")
	builder.WriteString("configuration=")
	builder.WriteString(j.Configuration)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(j.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(j.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Judges is a parsable slice of Judge.
type Judges []*Judge
