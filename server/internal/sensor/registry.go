package sensor

import (
	"github.com/charmingruby/densor/internal/sensor/domain/repository"
	"github.com/charmingruby/densor/internal/sensor/service"
	v1 "github.com/charmingruby/densor/internal/sensor/transport/rest/endpoint/v1"

	"github.com/gin-gonic/gin"
)

func NewService(sensorRepo repository.SensorRepository) service.SensorService {
	return service.NewSensorService(sensorRepo)
}

func NewHTTPService(router *gin.Engine) *v1.Handler {
	return v1.NewHandler(router)
}
