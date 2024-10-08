// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// JudgesColumns holds the columns for the "judges" table.
	JudgesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"local", "codeforces", "vjudge", "syoj", "noop"}, Default: "local"},
		{Name: "configuration", Type: field.TypeString},
	}
	// JudgesTable holds the schema information for the "judges" table.
	JudgesTable = &schema.Table{
		Name:       "judges",
		Columns:    JudgesColumns,
		PrimaryKey: []*schema.Column{JudgesColumns[0]},
	}
	// ProblemsColumns holds the columns for the "problems" table.
	ProblemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "judge_problems", Type: field.TypeInt},
	}
	// ProblemsTable holds the schema information for the "problems" table.
	ProblemsTable = &schema.Table{
		Name:       "problems",
		Columns:    ProblemsColumns,
		PrimaryKey: []*schema.Column{ProblemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "problems_judges_problems",
				Columns:    []*schema.Column{ProblemsColumns[5]},
				RefColumns: []*schema.Column{JudgesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SubmissionsColumns holds the columns for the "submissions" table.
	SubmissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "compiling", "judging", "finished"}, Default: "pending"},
		{Name: "verdict", Type: field.TypeEnum, Enums: []string{"OK", "TLE", "MLE", "ILE", "WA", "CE", "RE", "PE", "CRASHED", "OTHER"}, Default: "OK"},
		{Name: "test_count", Type: field.TypeInt, Default: 0},
		{Name: "problem_submissions", Type: field.TypeInt},
	}
	// SubmissionsTable holds the schema information for the "submissions" table.
	SubmissionsTable = &schema.Table{
		Name:       "submissions",
		Columns:    SubmissionsColumns,
		PrimaryKey: []*schema.Column{SubmissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "submissions_problems_submissions",
				Columns:    []*schema.Column{SubmissionsColumns[6]},
				RefColumns: []*schema.Column{ProblemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		JudgesTable,
		ProblemsTable,
		SubmissionsTable,
	}
)

func init() {
	ProblemsTable.ForeignKeys[0].RefTable = JudgesTable
	SubmissionsTable.ForeignKeys[0].RefTable = ProblemsTable
}
