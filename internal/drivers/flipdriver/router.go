package flipdriver

import (
	"net/http"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/gin-gonic/gin"
)

func AddFlipRoutes(h *gin.RouterGroup, flipSrv *services.FlipSrv) {
  h.GET("/flippening", func(c *gin.Context) {
    flippening := flipSrv.Get()
    c.JSON(http.StatusOK, flippening)
  })
}
