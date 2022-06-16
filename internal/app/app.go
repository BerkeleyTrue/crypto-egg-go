package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/berkeleytrue/crypto-agg-go/infra/httpserver"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	handler := gin.Default()
	handler.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	s := httpserver.New(handler, &cfg.HTTP)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	cleanup := s.Start()

	select {
	case err := <-s.Notify():
		fmt.Println(fmt.Errorf("app:Run:server: %w", err))

	case <-ctx.Done():
		fmt.Println("quitting")
		stop()
		cleanup()
		break
	}
}
