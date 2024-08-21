package ent

import (
	_ "ariga.io/ogent"
	_ "entgo.io/contrib/entgql"
	_ "entgo.io/contrib/entoas"
	_ "entgo.io/contrib/entproto"
	_ "entgo.io/ent/entc"
	_ "entgo.io/ent/entc/gen"
	_ "github.com/ogen-go/ogen"
)

//go:generate go run -mod=mod entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
