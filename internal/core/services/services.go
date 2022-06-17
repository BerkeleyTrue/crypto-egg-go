package services

import (
	"fmt"

	"github.com/berkeleytrue/crypto-agg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-agg-go/internal/core/ports"
)

type CoinService struct {
	repo ports.CoinRepo
}

func New(repo ports.CoinRepo) *CoinService {
	return &CoinService{
		repo: repo,
	}
}

func (srv *CoinService) Get(id string) (domain.Coin, error) {
	coin, err:= srv.repo.Get(id)
	if err != nil {
    return domain.Coin{}, fmt.Errorf("Could not get coin for id %s: %w", id, err)
	}

	return coin, nil
}