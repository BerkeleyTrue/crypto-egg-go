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
	coins := make(chan []domain.Coin, 1)

	ticker := time.NewTicker(5 * time.Second)
	quitTickerChan := make(chan struct{})

	go func() {
		var isPingOk bool = false
		for {
			select {
			case <-ticker.C:
				if isPingOk {
					res, err := srv.repo.GetCoins(ids)
					if err != nil || res == nil {
						fmt.Println(fmt.Errorf("GetCoins err: %w", err))
						isPingOk = false
					} else {
						coins <- res
					}
				} else {
					ok, err := srv.repo.Ping()
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

	return coins, func() {
		quitTickerChan <- struct{}{}
		close(quitTickerChan)
		close(coins)
	}
}
