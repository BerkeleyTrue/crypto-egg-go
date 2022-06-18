package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/services"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/coingecko"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/coinrepo"
	"github.com/berkeleytrue/crypto-egg-go/internal/driven/fliprepo"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/coin"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/flipdriver"
	"github.com/berkeleytrue/crypto-egg-go/internal/drivers/http/base"
	ginInfra "github.com/berkeleytrue/crypto-egg-go/internal/infra/gin"
	"github.com/berkeleytrue/crypto-egg-go/internal/infra/httpserver"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	coinSrv := services.New(coinrepo.NewMemKVS())
	cgSrv := services.CreateCoinGeckoSrv(coingecko.Init())
  flipSrv := services.CreateFlipSrv(fliprepo.CreateFlipRepo(), *coinSrv)

	handler := gin.New()
	ginInfra.AddGinHandlers(handler)
	base.NewRouter(handler, coinSrv)

	api := handler.Group("/api")
	coin.AddCoinRoutes(api, coinSrv)
	flipdriver.AddFlipRoutes(api, flipSrv)

	s := httpserver.New(handler, cfg.HTTP)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	coins, cgCleanup := cgSrv.StartService(cfg.Coins)
	cleanup := s.Start()

	go func() {
		for {
			select {
			case coin := <-coins:
				hasBTC := false
				hasEth := false

				for _, coin := range coin {
					if coin.Symbol == "btc" {
						hasBTC = true
					}
					if coin.Symbol == "eth" {
						hasEth = true
					}
					// fmt.Printf("updating %s\n", coin.ID)
					coinSrv.Update(coin)
				}

				if hasEth && hasBTC {
					flipSrv.Update()
				}
			}
		}
	}()

	select {
	case err := <-s.Notify():
		fmt.Println(fmt.Errorf("app:Run:server: %w", err))

	case <-ctx.Done():
		fmt.Println("quitting")
		stop()
		cleanup()
		cgCleanup()
		break
	}
}
