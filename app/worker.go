package app

import (
	"github.com/syoi-org/judge.syoi.org/config"
	"github.com/syoi-org/judge.syoi.org/logger"
	"go.uber.org/fx"
)

func NewWorker() *fx.App {
	return fx.New(
		config.WorkerModule,
		logger.Module,
	)
}
