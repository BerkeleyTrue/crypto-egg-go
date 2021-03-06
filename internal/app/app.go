package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/coingecko"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/coinrepo"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/fliprepo"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/gasapi"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/gasrepo"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/basedriver"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/coindriver"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/flipdriver"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/gasdriver"
	ginInfra "github.com/berkeleytrue/crypto-egg-go/internal/infra/gin"
	"github.com/berkeleytrue/crypto-egg-go/internal/infra/httpserver"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	coinSrv := services.CreateCoinSrv(coinrepo.NewMemKVS(), coingecko.Init())
	flipSrv := services.CreateFlipSrv(fliprepo.CreateFlipRepo(), *coinSrv)
	gasSrv := services.CreateGasSrv(gasrepo.CreateMemRepo(), gasapi.CreateGasApi())

	handler := ginInfra.CreateGinHandler(cfg)
	basedriver.NewRouter(handler, coinSrv, cfg)

	api := handler.Group("/api")
	coindriver.AddCoinRoutes(api, coinSrv)
	flipdriver.AddFlipRoutes(api, flipSrv)
	gasdriver.AddGasRoutes(api, gasSrv)

	s := httpserver.New(handler, cfg.HTTP)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	coinStream, cleanupCoin := coinSrv.StartService(cfg.Coins)
	flipSrv.StartService(coinStream)
	cleanupGas := gasSrv.StartService()
	cleanup := s.Start()

	select {
	case err := <-s.Notify():
		logger.Errorf("app:Run:server: %w", err)

	case <-ctx.Done():
		logger.Info("quitting")
		stop()
		cleanup()
		cleanupCoin()
		cleanupGas()
		break
	}
}
