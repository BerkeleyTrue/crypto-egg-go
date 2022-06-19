package services

import (
	"fmt"
	"time"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
)

type CoinService struct {
	repo ports.CoinRepo
	api  ports.CoinGeckoApi
}

func CreateCoinSrv(repo ports.CoinRepo, api ports.CoinGeckoApi) *CoinService {
	return &CoinService{
		repo,
		api,
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

func (srv *CoinService) StartService(ids []string) (chan []domain.Coin, func()) {
  coinsStream := make(chan []domain.Coin, 1)
	ticker := time.NewTicker(5 * time.Second)
	quitTickerChan := make(chan struct{})

	go func() {
		var isPingOk bool = false
		for {
			select {
			case <-ticker.C:
				if isPingOk {
					coins, err := srv.api.GetCoins(ids)
					if err != nil || coins == nil {
						fmt.Println(fmt.Errorf("GetCoins err: %w", err))
						isPingOk = false
					} else {
						for _, coin := range coins {
							srv.Update(coin)
						}
					  coinsStream <- coins
					}
				} else {
					ok, err := srv.api.Ping()
					if err != nil {
						fmt.Println(fmt.Errorf("Ping err: %w", err))
					} else if !ok {
						fmt.Println("ping not ok")
					} else {
						fmt.Println("ping Ok")
						isPingOk = true
					}
				}
			case <-quitTickerChan:
				ticker.Stop()
				return
			}
		}
	}()

	return coinsStream, func() {
		quitTickerChan <- struct{}{}
		close(quitTickerChan)
	}
}
