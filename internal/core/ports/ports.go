package ports

import "github.com/berkeleytrue/crypto-egg-go/internal/core/domain"

type CoinRepo interface {
	Get(id string) (domain.Coin, error)
	Update(id string, coin domain.Coin) (domain.Coin, error)
}
