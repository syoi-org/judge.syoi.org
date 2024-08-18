package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.WithLogger(func(logger *zap.SugaredLogger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger.Desugar()}
	}),
	fx.Decorate(RegisterLogLevelValidation),
)
