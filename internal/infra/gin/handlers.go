package gin

import (
	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/gin-gonic/gin"
)

func CreateGinHandler(cfg *config.Config) *gin.Engine {
	if cfg.Release == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

  h := gin.New()

	if cfg.Release != "production" {
		h.Use(gin.Logger())
	}

	h.Use(gin.Recovery())

	return h
}
