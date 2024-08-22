package healthz

import (
	"context"
	"time"

	"github.com/syoi-org/judy/ent/ogent"
)

type Service struct {
	startupTimestamp time.Time
}

func NewService() *Service {
	return &Service{
		startupTimestamp: time.Now(),
	}
}

func (hs *Service) HealthCheck(ctx context.Context) (*ogent.HealthCheckResult, error) {
	return &ogent.HealthCheckResult{
		Status: ogent.HealthCheckResultStatusOk,
		Uptime: time.Since(hs.startupTimestamp).String(),
	}, nil
}
