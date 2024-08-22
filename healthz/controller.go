package healthz

import (
	"context"

	"github.com/syoi-org/judy/ent/ogent"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

func (hc *Controller) HealthCheck(ctx context.Context) (ogent.HealthCheckRes, error) {
	result, err := hc.Service.HealthCheck(ctx)
	if err != nil {
		return nil, err
	}
	if result.Status == ogent.HealthCheckResultStatusError {
		return (*ogent.HealthCheckServiceUnavailable)(result), nil
	}
	return (*ogent.HealthCheckOK)(result), nil
}
