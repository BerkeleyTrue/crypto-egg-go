package basedriver

import (
	"net/http"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, coinSrv *services.CoinService, cfg *config.Config) {
	handler.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hash": cfg.Hash, "user": cfg.User, "time": cfg.Time})
	})

	handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "ok"}) })
}
