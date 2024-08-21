package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("code").NotEmpty().Unique(),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submissions", Submission.Type),
		edge.From("judge", Judge.Type).Ref("problems").Unique().Required(),
	}
}

func (Problem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		GraphQLMixin{},
	}
}
