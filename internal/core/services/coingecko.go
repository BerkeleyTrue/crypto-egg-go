package services

import (
	"fmt"
	"time"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
)

type CoinGeckoService struct {
	repo ports.CoinGeckoApi
}

func CreateCoinGeckoSrv(repo ports.CoinGeckoApi) *CoinGeckoService {
	return &CoinGeckoService{repo}
}

func (srv *CoinGeckoService) StartService(ids []string) (chan []domain.Coin, func()) {
  coins := make(chan []domain.Coin)

	ticker := time.NewTicker(5 * time.Second)
	quitTickerChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				ok, err := srv.repo.Ping()
				if err != nil {
					fmt.Println(fmt.Errorf("Ping err: %w", err))
				} else if !ok {
					fmt.Println("ping not ok")
				} else {
					fmt.Println("ping Ok")
				}
			case <-quitTickerChan:
				ticker.Stop()
				return
			}
		}
	}()

	return coins, func() {
		quitTickerChan <- struct{}{}
		close(quitTickerChan)
	  close(coins)
	}
}
