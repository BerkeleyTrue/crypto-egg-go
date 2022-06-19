package services

import (
	"time"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
	"go.uber.org/zap"
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
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

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
						logger.Error("GetCoins err: %w", "err", err)
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
						logger.Error("Ping err: %w", "err", err)
					} else if !ok {
						logger.Info("ping not ok")
					} else {
						logger.Info("ping Ok")
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
