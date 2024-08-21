//go:build ignore
// +build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

func main() {
	protoext, err := entproto.NewExtension()
	if err != nil {
		log.Fatalf("creating entproto extension: %v", err)
	}
	spec := ogen.NewSpec().
		SetOpenAPI("3.0.3").
		SetInfo(
			ogen.NewInfo().
				SetTitle("Judy Judge API").
				SetVersion("1.0").
				SetDescription("This is a API for running Judy Judge."),
		)
	entoasext, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating enoas extension: %v", err)
	}
	ogentext, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(protoext, ogentext, entoasext))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
