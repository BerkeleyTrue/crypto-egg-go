package basedriver

import (
	"net/http"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, coinSrv *services.CoinService) {
	handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "ok"}) })
}