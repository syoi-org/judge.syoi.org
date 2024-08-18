package app

import (
	"github.com/syoi-org/judge.syoi.org/config"
	"github.com/syoi-org/judge.syoi.org/logger"
	"github.com/syoi-org/judge.syoi.org/transport"
	"go.uber.org/fx"
)

func NewServer() *fx.App {
	return fx.New(
		config.ServerModule,
		logger.Module,
		transport.Module,
	)
}
