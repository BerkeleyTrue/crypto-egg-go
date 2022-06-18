package services

import (
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
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
	ratio := btc.MarketCap / eth.MarketCap

	srv.repo.Update(ratio)
	return srv.repo.Get(), nil
}
