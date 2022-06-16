package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "ok"}) })

	h := handler.Group("/api")
	addCoinRoutes(h)
}
