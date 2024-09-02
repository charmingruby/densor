package v1

import (
	"github.com/charmingruby/densor/internal/sensor/domain/usecase"
	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, sensorSvc usecase.SensorService) *Handler {
	return &Handler{
		router:    router,
		sensorSvc: sensorSvc,
	}
}

type Handler struct {
	router    *gin.Engine
	sensorSvc usecase.SensorService
}

func (h *Handler) Register() {
	basePath := "/api/v1"
	v1 := h.router.Group(basePath)
	{
		v1.GET("/welcome", welcomeEndpoint)
		v1.POST("/sensors", h.createSensorEndpoint)
		v1.GET("/sensors", h.fetchSensorsEndpoint)
	}
}
