package transport

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewHttp),
	fx.Invoke(runHttpServer),
	fx.Provide(NewGrpc),
	fx.Invoke(runGrpcServer),
	fx.Provide(NewRouter),
	fx.Provide(AsControllerRoute(NewSwaggerHandler)),
	fx.Provide(AsHandlerRoute(NewGraphqlHandler), AsHandlerRoute(NewGraphQLPlaygroundHandler)),
	fx.Provide(NewOgentServer),
)
