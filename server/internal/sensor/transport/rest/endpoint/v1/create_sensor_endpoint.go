package v1

import (
	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/sensor/domain/usecase"
	"github.com/gin-gonic/gin"
)

type CreateSensorRequest struct {
	Name               string `json:"name" binding:"required"`
	Description        string `json:"description" binding:"required"`
	SensorCategoryName string `json:"sensor_category_name" binding:"required"`
	EquipmentName      string `json:"equipment_name"`
	Observation        string `json:"observation"`
	SectorName         string `json:"sector_name"`
}

func (h *Handler) createSensorEndpoint(c *gin.Context) {
	var req CreateSensorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api_rest.NewPayloadError(c, err)
		return
	}

	if err := h.sensorSvc.CreateSensorUseCase(usecase.CreateSensorInputDTO{
		Name:               req.Name,
		Description:        req.Description,
		SensorCategoryName: req.SensorCategoryName,
		EquipmentName:      req.EquipmentName,
		Observation:        req.Observation,
		SectorName:         req.SectorName,
	}); err != nil {
		if conflictErr, ok := err.(*core.ErrConflict); ok {
			api_rest.NewConflictErr(c, conflictErr)
			return
		}

		if entityErr, ok := err.(*core.ErrValidation); ok {
			api_rest.NewEntityError(c, entityErr)
			return
		}

		if notFoundErr, ok := err.(*core.ErrNotFound); ok {
			api_rest.NewResourceNotFoundError(c, notFoundErr)
			return
		}

		api_rest.NewInternalServerError(c, err)
		return
	}

	api_rest.NewCreatedResponse(c, "sensor")
}
