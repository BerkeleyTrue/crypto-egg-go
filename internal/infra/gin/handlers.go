package gin

import "github.com/gin-gonic/gin"

func AddGinHandlers(h *gin.Engine) {
	h.Use(gin.Logger())
	h.Use(gin.Recovery())
}
