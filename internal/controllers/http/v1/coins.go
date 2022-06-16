package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addCoinRoutes(h *gin.RouterGroup) {
	h.GET("/coin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "coins"})
	})
}
