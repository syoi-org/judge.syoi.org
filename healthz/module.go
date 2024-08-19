package healthz

import (
	"github.com/syoi-org/judy/transport"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(transport.AsControllerRoute(NewController)),
	fx.Provide(NewService),
)
