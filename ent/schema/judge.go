package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
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
		field.String("name").NotEmpty().
			Comment("Name of the judge. Example: Aizu Online Judge").
			Annotations(entoas.Example("Aizu Online Judge")),
		field.String("code").NotEmpty().Unique().
			Comment("Unique codename of the judge. Example: AZOJ").
			Annotations(entoas.Example("AZOJ")),
		field.Enum("type").Values("local", "codeforces", "vjudge", "syoj", "noop").Default("local").
			Comment("Type of the judge. Example: local").
			Annotations(entoas.Example("local")),
		field.String("configuration").
			Comment("Configuration of the judge. Encoded in form urlencoded format (key1=value1&key2=value2...).").
			Annotations(entoas.Example("api_key=supersecret&api_url=https://aizu.example.com/api")),
	}
}

// Edges of the Judge.
func (Judge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problems", Problem.Type).Annotations(entgql.RelayConnection()),
	}
}

func (Judge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		GraphQLMixin{},
	}
}
