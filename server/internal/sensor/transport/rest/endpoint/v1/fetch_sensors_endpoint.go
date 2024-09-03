package v1

import (
	"fmt"

	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/gin-gonic/gin"
)

func (h *Handler) fetchSensorsEndpoint(c *gin.Context) {
	sensors, err := h.sensorSvc.FetchSensorsUseCase()

	if err != nil {
		api_rest.NewInternalServerError(c, err)
		return
	}

	api_rest.NewOkResponse(
		c,
		fmt.Sprintf("%d sensors found", len(sensors)),
		sensors,
	)
}
