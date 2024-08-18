package healthz

import (
	"context"
	"time"
)

type ResultStatus string

var (
	StatusOK    ResultStatus = "ok"
	StatusError ResultStatus = "error"
)

type Result struct {
	Status  string        `json:"status" example:"ok" enums:"ok,error"`
	Uptime  string        `json:"uptime"`
	Details ResultDetails `json:"details"`
}

type ResultDetails struct {
	Db ResultDetail `json:"db"`
}

type ResultDetail struct {
	Status ResultStatus `json:"status" example:"ok" enums:"ok,error"`
	Error  string       `json:"error,omitempty"`
}

type Service struct {
	startupTimestamp time.Time
}

func NewService() *Service {
	return &Service{
		startupTimestamp: time.Now(),
	}
}

func (hs *Service) HealthCheck(ctx context.Context) (*Result, error) {
	return &Result{
		Status: "ok",
		Uptime: time.Since(hs.startupTimestamp).String(),
		Details: ResultDetails{
			Db: ResultDetail{
				Status: "ok",
			},
		},
	}, nil
}
