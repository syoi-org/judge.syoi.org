package transport

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/syoi-org/judy/ent"
	"github.com/syoi-org/judy/ent/graph"
	"go.uber.org/fx"
)

type GraphQLParams struct {
	fx.In
	EntClient *ent.Client
}

type GraphQLHandler struct {
	EntClient *ent.Client
}

func NewGraphqlHandler(p GraphQLParams) *GraphQLHandler {
	return &GraphQLHandler{
		EntClient: p.EntClient,
	}
}

func (h *GraphQLHandler) HttpHandler() http.Handler {
	return handler.NewDefaultServer(graph.NewSchema(h.EntClient))
}

func (h *GraphQLHandler) RoutePattern() string {
	return "/graphql"
}

type GraphQLPlaygroundHandler struct{}

func NewGraphQLPlaygroundHandler() *GraphQLPlaygroundHandler {
	return &GraphQLPlaygroundHandler{}
}

func (h *GraphQLPlaygroundHandler) HttpHandler() http.Handler {
	return playground.Handler("GraphQL", new(GraphQLHandler).RoutePattern())
}

func (h *GraphQLPlaygroundHandler) RoutePattern() string {
	return "/graphiql"
}
