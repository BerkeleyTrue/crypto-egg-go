package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/berkeleytrue/crypto-agg-go/internal/controllers/http/base"
	"github.com/berkeleytrue/crypto-agg-go/internal/core/services"
	ginInfra "github.com/berkeleytrue/crypto-agg-go/internal/infra/gin"
	"github.com/berkeleytrue/crypto-agg-go/internal/infra/httpserver"
	"github.com/berkeleytrue/crypto-agg-go/internal/repos/coinrepo"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
  coinSrv := services.New(coinrepo.NewMemKVS())
	handler := gin.New()
	ginInfra.AddGinHandlers(handler)
	base.NewRouter(handler, coinSrv)

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
