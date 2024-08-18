package healthz

import (
	"github.com/syoi-org/judge.syoi.org/transport"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(transport.AsControllerRoute(NewController)),
	fx.Provide(NewService),
)
