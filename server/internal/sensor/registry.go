package sensor

import (
	"github.com/charmingruby/densor/internal/sensor/domain/repository"
	"github.com/charmingruby/densor/internal/sensor/service"
	v1 "github.com/charmingruby/densor/internal/sensor/transport/rest/endpoint/v1"

	"github.com/gin-gonic/gin"
)

func NewService(
	sensorRepo repository.SensorRepository,
	sensorCategoryRepo repository.SensorCategoryRepository,
	sectorRepo repository.SectorRepository,
	equipmentRepo repository.EquipmentRepository,
) service.SensorService {
	return service.NewSensorService(sensorRepo, sensorCategoryRepo, sectorRepo, equipmentRepo)
}

func NewHTTPService(router *gin.Engine, sensorSvc service.SensorService) *v1.Handler {
	return v1.NewHandler(router, &sensorSvc)
}
