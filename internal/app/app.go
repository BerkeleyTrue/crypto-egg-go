package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/coinrepo"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/coin"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/http/base"
	ginInfra "github.com/berkeleytrue/crypto-egg-go/internal/infra/gin"
	"github.com/berkeleytrue/crypto-egg-go/internal/infra/httpserver"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	coinSrv := services.New(coinrepo.NewMemKVS())
	handler := gin.New()
	ginInfra.AddGinHandlers(handler)
	base.NewRouter(handler, coinSrv)

	api := handler.Group("/api")
	coin.AddCoinRoutes(api, coinSrv)

	s := httpserver.New(handler, cfg.HTTP)

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
