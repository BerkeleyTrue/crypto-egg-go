package services

import (
	"fmt"
	"time"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
	"go.uber.org/zap"
)

type GasSrv struct {
	repo ports.GasRepo
	api  ports.GasApi
}

func CreateGasSrv(repo ports.GasRepo, api ports.GasApi) *GasSrv {
	return &GasSrv{repo, api}
}

func (srv *GasSrv) Get() domain.Gas {
	return srv.repo.Get()
}

func (srv *GasSrv) Update(baseFee float32) domain.Gas {
	return srv.repo.Update(domain.Gas{
		ID:        "ethereum",
		Base:      baseFee,
		UpdatedAt: time.Now(),
	})
}

func (srv *GasSrv) StartService() func() {
  logger := zap.NewExample().Sugar()
  defer logger.Sync()

	ticker := time.NewTicker(6 * time.Second)
	quitTickerChan := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				baseFee, err := srv.api.Get()

				if err != nil {
					fmt.Println(fmt.Errorf("GasSrv:api:Get: %w", err))
					break
				}
				// log.Printf("baseFee: %f", baseFee)

				srv.Update(baseFee)
			case <-quitTickerChan:
				ticker.Stop()
				return
			}
		}
	}()

	return func() {
		quitTickerChan <- struct{}{}
		close(quitTickerChan)
	}
}
