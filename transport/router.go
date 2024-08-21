package transport

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/syoi-org/judy/ent/ogent"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Logger           *zap.SugaredLogger `optional:"true"`
	ControllerRoutes []ControllerRoute  `group:"controllerRoutes"`
	HandlerRoutes    []HandlerRoute     `group:"handlerRoutes"`
	OgentServer      *ogent.Server
}

type Result struct {
	fx.Out
	Http *gin.Engine
}

func NewRouter(p Params) Result {
	gin.SetMode(gin.ReleaseMode)

	server := gin.New()
	if p.Logger != nil {
		server.Use(ginzap.Ginzap(p.Logger.Desugar(), time.RFC3339, true))
	}
	server.Use(gin.Recovery())

	for _, route := range p.ControllerRoutes {
		if p.Logger != nil {
			p.Logger.Infow("registering controller route", "pattern", route.RoutePattern())
		}
		route.RegisterControllerRoutes(
			server.Group(route.RoutePattern()),
		)
	}

	for _, route := range p.HandlerRoutes {
		if p.Logger != nil {
			p.Logger.Infow("registering handler route", "pattern", route.RoutePattern())
		}
		server.Any(route.RoutePattern(), gin.WrapH(route.HttpHandler()))
	}

	server.NoRoute(gin.WrapH(p.OgentServer))

	return Result{
		Http: server,
	}
}

func (r *Result) GetHttpRouter() *gin.Engine {
	return r.Http
}

type ControllerRoute interface {
	RegisterControllerRoutes(rg *gin.RouterGroup)
	RoutePattern() string
}

func AsControllerRoute(controller any) any {
	return fx.Annotate(
		controller,
		fx.As(new(ControllerRoute)),
		fx.ResultTags(`group:"controllerRoutes"`),
	)
}

type HandlerRoute interface {
	HttpHandler() http.Handler
	RoutePattern() string
}

func AsHandlerRoute(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(HandlerRoute)),
		fx.ResultTags(`group:"handlerRoutes"`),
	)
}
