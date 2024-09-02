package v1

import (
	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/charmingruby/densor/internal/sensor/domain/usecase"
	"github.com/gin-gonic/gin"
)

type CreateSensorRequest struct {
	Name             string `json:"name" binding:"required"`
	Description      string `json:"description" binding:"required"`
	SensorCategoryID string `json:"sensor_category_id"`
	EquipmentID      string `json:"equipment_id" binding:"required"`
	Observation      string `json:"observation"`
	SectorID         string `json:"sector_id"`
}

func (h *Handler) createSensorEndpoint(c *gin.Context) {
	var req CreateSensorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api_rest.NewPayloadError(c, err)
		return
	}

	if err := h.sensorSvc.CreateSensorUseCase(usecase.CreateSensorInputDTO{
		Name:             req.Name,
		Description:      req.Description,
		SensorCategoryID: req.SensorCategoryID,
		EquipmentID:      req.EquipmentID,
		Observation:      req.Observation,
		SectorID:         req.SectorID,
	}); err != nil {
		api_rest.NewInternalServerError(c, err)
		return
	}

	api_rest.NewCreatedResponse(c, "sensor")
}
