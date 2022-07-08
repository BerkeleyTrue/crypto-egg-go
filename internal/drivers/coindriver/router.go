package coindriver

import (
	"net/http"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddCoinRoutes(h *gin.RouterGroup, coinSrv *services.CoinService) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	h.GET("/coin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "coins"})
	})

	h.GET("/coins", func(c *gin.Context) {
		coins := coinSrv.GetAll()
		c.JSON(http.StatusOK, coins)
	})

	h.GET("/coins/sym/:sym", func(c *gin.Context) {
		sym := c.Param("sym")
		coin := coinSrv.GetBySymbol(sym)

		if coin.ID == "" {
			logger.Debug("not found")
			c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
			return
		}

		c.JSON(http.StatusOK, coin)
	})

	h.GET("/coins/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		coin := coinSrv.Get(id)

		if coin.ID == "" {
			logger.Debug("not found")
			c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
			return
		}

		c.JSON(http.StatusOK, coin)
	})

}
