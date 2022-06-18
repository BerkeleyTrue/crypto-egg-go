package services

import (
	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
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
	return srv.repo.Get(id)
}

func (srv *CoinService) GetBySymbol(sym string) (domain.Coin, error) {
	return srv.repo.GetBySymbol(sym)
}

func (srv *CoinService) Update(c domain.Coin) (domain.Coin, error) {
	return srv.repo.Update(c.ID, c)
}

func (srv *CoinService) GetAll() ([]domain.Coin, error) {
	return srv.repo.GetAll()
}
