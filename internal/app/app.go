package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/berkeleytrue/crypto-agg-go/infra/httpserver"
	"github.com/berkeleytrue/crypto-agg-go/internal/controllers/http/v1"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	handler := gin.New()
	v1.NewRouter(handler)

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
