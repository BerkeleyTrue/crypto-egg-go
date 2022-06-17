package ports

import "github.com/berkeleytrue/crypto-egg-go/internal/core/domain"

var apiUrl = "https://api.coingecko.com/api/v3/"

type CoinRepo interface {
	Get(id string) (domain.Coin, error)
	Update(id string, coin domain.Coin) (domain.Coin, error)
}

type CoinGeckoApi interface {
  Ping() (bool, error)
  GetCoins(ids []string) ([]domain.Coin, error)
}
