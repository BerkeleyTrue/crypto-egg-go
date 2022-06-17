package base

import (
	"net/http"

	"github.com/berkeleytrue/crypto-agg-go/internal/controllers/coin"
	"github.com/berkeleytrue/crypto-agg-go/internal/core/services"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, coinSrv *services.CoinService) {
	handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "ok"}) })

	api := handler.Group("/api")
	coin.AddCoinRoutes(api, coinSrv)
}
