package ports

import "github.com/berkeleytrue/crypto-egg-go/internal/core/domain"

var apiUrl = "https://api.coingecko.com/api/v3/"

type CoinRepo interface {
	Get(id string) (domain.Coin, error)
	GetBySymbol(sym string) (domain.Coin, error)
	Update(id string, coin domain.Coin) (domain.Coin, error)
	GetAll() ([]domain.Coin, error)
}

type CoinGeckoApi interface {
	Ping() (bool, error)
	GetCoins(ids []string) ([]domain.Coin, error)
}

type FlipRepo interface {
	Get() domain.Flippening
	Update(ratio float64, btcCap int64, ethCap int64)
}

type GasRepo interface {
	Get() domain.Gas
	Update(update domain.Gas) domain.Gas
}

type GasApi interface {
  Get() (float32, float32, error)
}
