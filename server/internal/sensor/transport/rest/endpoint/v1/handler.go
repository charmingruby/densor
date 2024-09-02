package v1

import (
	docs "github.com/charmingruby/densor/api"
	"github.com/charmingruby/densor/internal/sensor/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHandler(router *gin.Engine, sensorSvc service.SensorService) *Handler {
	return &Handler{
		router:    router,
		sensorSvc: sensorSvc,
	}
}

type Handler struct {
	router    *gin.Engine
	sensorSvc service.SensorService
}

func (h *Handler) Register() {
	basePath := "/api/v1"
	v1 := h.router.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath
	{
		v1.GET("/welcome", welcomeEndpoint)
		v1.POST("/sensors", h.createSensorEndpoint)
		v1.GET("/sensors", h.fetchSensorsEndpoint)
	}

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
