package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

// healthCheck godoc
//
//	@Summary		Health Checking
//	@Description	Health Checking for API services
//	@Produce		json
//	@Success		200	{object}	Result
//	@Failure		503	{object}	Result
//	@Router			/healthz [get]
func (hc *Controller) healthCheck(ctx *gin.Context) {
	result, err := hc.Service.HealthCheck(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if result.Status == "ok" {
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusServiceUnavailable, result)
	}
}

func (hc *Controller) RoutePattern() string {
	return "/healthz"
}

func (hc *Controller) RegisterControllerRoutes(rg *gin.RouterGroup) {
	rg.GET("", hc.healthCheck)
}
