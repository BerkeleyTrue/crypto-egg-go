package app

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/berkeleytrue/crypto-agg-go/infra/httpserver"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config, wg *sync.WaitGroup) {
	handler := gin.Default()
	handler.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

  s := httpserver.New(handler, &cfg.HTTP, wg)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

  httpserver.Start(s)
  wg.Wait()
}
