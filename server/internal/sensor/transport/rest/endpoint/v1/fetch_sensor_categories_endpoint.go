package v1

import (
	"fmt"

	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/gin-gonic/gin"
)

func (h *Handler) fetchSensorCategoriesEndpoint(c *gin.Context) {
	categories, err := h.sensorSvc.FetchSensorCategoriesUseCase()

	if err != nil {
		api_rest.NewInternalServerError(c, err)
		return
	}

	api_rest.NewOkResponse(
		c,
		fmt.Sprintf("%d sensor categories found", len(categories)),
		categories,
	)
}
