package app

import (
	"github.com/syoi-org/judy/config"
	"github.com/syoi-org/judy/db"
	"github.com/syoi-org/judy/healthz"
	"github.com/syoi-org/judy/internal/broker"
	"github.com/syoi-org/judy/logger"
	"github.com/syoi-org/judy/transport"
	"go.uber.org/fx"
)

func NewServer() *fx.App {
	return fx.New(
		config.ServerModule,
		logger.Module,
		transport.Module,
		healthz.Module,
		db.Module,
		broker.Module,
	)
}
