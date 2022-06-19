package services

import (
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
	"go.uber.org/zap"
)

type FlipSrv struct {
	repo    ports.FlipRepo
	coinSrv CoinService
}

func CreateFlipSrv(repo ports.FlipRepo, coinSrv CoinService) *FlipSrv {
	return &FlipSrv{
		repo,
		coinSrv,
	}
}

func (srv *FlipSrv) Get() domain.Flippening {
	return srv.repo.Get()
}

func (srv *FlipSrv) Update() (domain.Flippening, error) {
	btc, err := srv.coinSrv.Get("bitcoin")
	if err != nil {
		return domain.Flippening{}, fmt.Errorf("Couldn't fetch bitcoin market cap: %w", err)
	}

	if btc.MarketCap == 0 {
		return domain.Flippening{}, fmt.Errorf("No bitcoin market cap returned")
	}

	eth, err := srv.coinSrv.Get("ethereum")

	if err != nil {
		return domain.Flippening{}, fmt.Errorf("Couldn't fetch ethereum market cap: %w", err)
	}

	if eth.MarketCap == 0 {
		return domain.Flippening{}, fmt.Errorf("No ethereum market cap returned")
	}
	var ratio float64 = float64(eth.MarketCap) / float64(btc.MarketCap)

	srv.repo.Update(ratio, eth.MarketCap, btc.MarketCap)
	return srv.repo.Get(), nil
}

func (srv *FlipSrv) StartService(coinsStream chan []domain.Coin) func() {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	go func() {
		select {
		case coins := <-coinsStream:
			logger.Info("coin updated")
			hasBtc := false
			hasEth := false

			for _, coin := range coins {
				if coin.Symbol == "btc" {
					hasBtc = true
				}
				if coin.Symbol == "eth" {
					hasEth = true
				}
			}

			if hasBtc && hasEth {
				srv.Update()
			} else {
				logger.Info("No btc or eth found")
			}
		}
	}()
	return func() {}
}
