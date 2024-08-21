package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Judge holds the schema definition for the Judge entity.
type Judge struct {
	ent.Schema
}

// Fields of the Judge.
func (Judge) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("code").NotEmpty().Unique(),
		field.Enum("type").Values("local", "codeforces", "vjudge", "syoj", "noop").Default("local"),
		field.String("configuration"),
	}
}

// Edges of the Judge.
func (Judge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problems", Problem.Type),
	}
}

func (Judge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
