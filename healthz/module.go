package healthz

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewController),
	fx.Provide(NewService),
)
