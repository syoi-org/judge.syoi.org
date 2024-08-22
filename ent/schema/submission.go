package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Submission holds the schema definition for the Submission entity.
type Submission struct {
	ent.Schema
}

// Fields of the Submission.
func (Submission) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("pending", "compiling", "judging", "finished").Default("pending").
			Comment("Status of the submission. Example: finished"),
		field.Enum("verdict").Values("OK", "TLE", "MLE", "ILE", "WA", "CE", "RE", "PE", "CRASHED", "OTHER").Default("OK").
			Comment("Verdict of the submission. Example: OK"),
		field.Int("test_count").Default(0).NonNegative().
			Comment("Number of test cases finished or currently judging. Example: 13"),
	}
}

// Edges of the Submission.
func (Submission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("problem", Problem.Type).Ref("submissions").Unique().Required(),
	}
}

func (Submission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		GraphQLMixin{},
	}
}
