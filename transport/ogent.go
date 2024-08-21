package transport

import (
	"github.com/syoi-org/judy/ent"
	"github.com/syoi-org/judy/ent/ogent"
	"go.uber.org/fx"
)

type OgentServerParams struct {
	fx.In
	EntClient *ent.Client
}

func NewOgentServer(p OgentServerParams) (*ogent.Server, error) {

	server, err := ogent.NewServer(ogent.NewOgentHandler(p.EntClient))
	if err != nil {
		return nil, err
	}
	return server, nil
}
