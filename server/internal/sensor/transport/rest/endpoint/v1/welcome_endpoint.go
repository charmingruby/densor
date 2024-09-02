package v1

import (
	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/gin-gonic/gin"
)

func welcomeEndpoint(c *gin.Context) {
	api_rest.NewOkResponse(c, "OK!", nil)
}
