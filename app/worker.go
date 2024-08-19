package app

import (
	"github.com/syoi-org/judy/config"
	"github.com/syoi-org/judy/logger"
	"go.uber.org/fx"
)

func NewWorker() *fx.App {
	return fx.New(
		config.WorkerModule,
		logger.Module,
	)
}
