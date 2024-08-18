package db

import "go.uber.org/fx"

var Module = fx.Module("db",
	fx.Provide(NewDbClient),
	fx.Invoke(runDbClient),
)
