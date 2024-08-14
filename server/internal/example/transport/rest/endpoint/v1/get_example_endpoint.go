package v1

import (
	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/example/transport/rest/presenter"
	"github.com/gin-gonic/gin"
)

type GetExampleResponse struct {
	Message string                      `json:"message"`
	Data    *presenter.ExamplePresenter `json:"data"`
	Code    int                         `json:"status_code"`
}

// GetExample godoc
//
//	@Summary		Gets an example
//	@Description	Gets an example
//	@Tags			Examples
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Get Example Payload"
//	@Success		200	{object}	GetExampleResponse
//	@Failure		404	{object}	api_rest.Response
//	@Failure		500	{object}	api_rest.Response
//	@Router			/examples/{id} [get]
func (h *Handler) getExampleEndpoint(c *gin.Context) {
	exampleID := c.Param("id")

	example, err := h.exampleService.GetExampleUseCase(exampleID)
	if err != nil {
		resourceNotFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			api_rest.NewResourceNotFoundError(c, resourceNotFoundErr)
			return
		}

		api_rest.NewInternalServerError(c, err)
		return
	}

	examplePresenter := presenter.NewExamplePresenter(*example)

	api_rest.NewOkResponse(
		c,
		"example found",
		examplePresenter,
	)
}
