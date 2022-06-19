package services

import (
	"fmt"
	"time"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/ports"
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

func (srv *GasSrv) Update(baseFee float32, avgTime float32) domain.Gas {
	return srv.repo.Update(domain.Gas{Base: baseFee, AvgTime: avgTime})
}

func (srv *GasSrv) StartService() func() {
	ticker := time.NewTicker(5 * time.Second)
	quitTickerChan := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				baseFee, avgTime, err := srv.api.Get()

				if err != nil {
					fmt.Println(fmt.Errorf("GasSrv:api:Get: %w", err))
					return
				}

				srv.Update(baseFee, avgTime)
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
