package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/syoi-org/judy/ent"
)

type SwaggerHandler struct{}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{}
}

func (h *SwaggerHandler) RegisterControllerRoutes(rg *gin.RouterGroup) {
	rg.Any("", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/docs/swagger/index.html")
	})
	rg.StaticFileFS("/openapi.json", "openapi.json", http.FS(ent.OpenAPIFs))
	rg.Any("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/docs/openapi.json"),
	))
}

func (h *SwaggerHandler) RoutePattern() string {
	return "/docs"
}

var _ ControllerRoute = (*SwaggerHandler)(nil)
