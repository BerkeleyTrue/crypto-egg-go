package gasdriver

import (
	"net/http"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/gin-gonic/gin"
)

func AddGasRoutes(h *gin.RouterGroup, gasSrv *services.GasSrv) {
  h.GET("/gas", func(c *gin.Context) {
    gas := gasSrv.Get()
		c.JSON(http.StatusOK, gas)
  })
}
