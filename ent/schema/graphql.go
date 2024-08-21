package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/mixin"
)

type GraphQLMixin struct {
	mixin.Schema
}

func (GraphQLMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
