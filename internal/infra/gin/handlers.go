package gin

import (
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/gin-gonic/gin"
)

func AddGinHandlers(h *gin.Engine, cfg *config.Config) {
	if cfg.Release != "production" {
	  fmt.Println("adding logger")
    h.Use(gin.Logger())
	}
	h.Use(gin.Recovery())
}
