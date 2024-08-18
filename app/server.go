package app

import (
	"github.com/syoi-org/judge.syoi.org/config"
	"github.com/syoi-org/judge.syoi.org/logger"
	"go.uber.org/fx"
)

func NewServer() *fx.App {
	return fx.New(
		config.ServerModule,
		logger.Module,
	)
}
