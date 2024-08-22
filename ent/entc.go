//go:build ignore
// +build ignore

package main

import (
	"encoding/json"

	"ariga.io/ogent"
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
	"go.uber.org/zap"
)

func additionalEndpoints(graph *gen.Graph, spec *ogen.Spec) error {
	spec.AddSchema("HealthCheckResult", ogen.NewSchema().
		SetDescription("Result of health check").
		AddRequiredProperties(
			ogen.NewProperty().SetName("status").SetSchema(ogen.String().SetEnum([]json.RawMessage{json.RawMessage(`"ok"`), json.RawMessage(`"error"`)})),
			ogen.NewProperty().SetName("uptime").SetSchema(ogen.String()),
		).
		AddOptionalProperties(
			ogen.NewProperty().SetName("errors").SetSchema(ogen.NewSchema()),
		),
	)
	spec.AddPathItem("/healthz", ogen.NewPathItem().
		SetDescription("Health Checking for API services").
		SetGet(ogen.NewOperation().SetOperationID("healthCheck").
			SetSummary("Health Checking").
			SetDescription("Health Checking for API services").
			AddResponse("200", ogen.NewResponse().AddContent("application/json", ogen.NewSchema().SetRef("#/components/schemas/HealthCheckResult"))).
			AddResponse("503", ogen.NewResponse().AddContent("application/json", ogen.NewSchema().SetRef("#/components/schemas/HealthCheckResult"))).
			AddResponse("500", ogen.NewResponse().SetRef("#/components/responses/500")),
		),
	)
	return nil
}

func main() {
	entgqlext, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
	)
	if err != nil {
		zap.S().Fatalf("creating entgql extension: %v", err)
	}
	protoext, err := entproto.NewExtension()
	if err != nil {
		zap.S().Fatalf("creating entproto extension: %v", err)
	}
	spec := ogen.NewSpec().
		SetOpenAPI("3.0.3").
		SetInfo(
			ogen.NewInfo().
				SetTitle("Judy Judge API").
				SetVersion("1.0").
				SetDescription("This is a API for running Judy Judge."),
		)
	entoasext, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(additionalEndpoints),
	)
	if err != nil {
		zap.S().Fatalf("creating enoas extension: %v", err)
	}
	ogentext, err := ogent.NewExtension(spec)
	if err != nil {
		zap.S().Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(
		entgqlext,
		protoext,
		ogentext,
		entoasext,
	))
	if err != nil {
		zap.S().Fatalf("running ent codegen: %v", err)
	}
}
