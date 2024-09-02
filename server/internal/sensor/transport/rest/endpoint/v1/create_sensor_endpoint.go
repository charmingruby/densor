package v1

import "github.com/gin-gonic/gin"

type CreateSensorRequest struct {
	Name             string `json:"name" binding:"required"`
	Description      string `json:"description" binding:"required"`
	SensorCategoryID string `json:"sensor_category_id"`
	EquipmentID      string `json:"equipment_id" binding:"required"`
	Observation      string `json:"observation"`
	SectorID         string `json:"sector_id"`
}

// Create Sensor godoc
//
//	@Summary		Creates an sensor
//	@Description	Creates an sensor
//	@Tags			Sensors
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateSensorRequest	true	"Create Sensor Payload"
//	@Success		201		{object}	api_rest.Response
//	@Failure		400		{object}	api_rest.Response
//	@Failure		500		{object}	api_rest.Response
//	@Router			/sensors [post]
func (h *Handler) createSensorEndpoint(c *gin.Context) {}
