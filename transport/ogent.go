package transport

import (
	"github.com/syoi-org/judy/ent"
	"github.com/syoi-org/judy/ent/ogent"
	"github.com/syoi-org/judy/healthz"
	"go.uber.org/fx"
)

type ogentHandler struct {
	*ogent.OgentHandler
	*healthz.Controller
}

type OgentServerParams struct {
	fx.In
	EntClient          *ent.Client
	HealthzCountroller *healthz.Controller
}

func NewOgentServer(p OgentServerParams) (*ogent.Server, error) {
	server, err := ogent.NewServer(ogentHandler{
		OgentHandler: ogent.NewOgentHandler(p.EntClient),
		Controller:   p.HealthzCountroller,
	})
	if err != nil {
		return nil, err
	}
	return server, nil
}
