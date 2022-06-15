package app

import (
	"fmt"
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

	s.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)


	select {
	case s := <-interrupt:
		fmt.Printf("app:Run:Signal: %s\n", s.String())
	case err := <-s.Notify():
		fmt.Println(fmt.Errorf("app:Run:server: %w", err))
	}

	wg.Wait()
}
