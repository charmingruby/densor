package v1

import (
	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/gin-gonic/gin"
)

// Welcome godoc
//
//	@Summary		Health Check
//	@Description	Health Check
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	api_rest.Response
//	@Router			/welcome [get]
func welcomeEndpoint(c *gin.Context) {
	api_rest.NewOkResponse(c, "OK!", nil)
}
