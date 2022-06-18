package fliprepo

import (
	"time"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
)

type flipRatio struct {
	ratio    float64
	ethCap   int64
	btcCap   int64
	init     bool
	updateAt time.Time
}

func CreateFlipRepo() *flipRatio {
	return &flipRatio{
		ratio:  0,
		btcCap: 0,
		ethCap: 0,
		init:   false,
	}
}

func (repo *flipRatio) Get() domain.Flippening {
	return domain.Flippening{
		Ratio:     repo.ratio,
		EthCap:    repo.ethCap,
		BtcCap:    repo.btcCap,
		Init:      repo.init,
		UpdatedAt: repo.updateAt,
	}
}

func (repo *flipRatio) Update(ratio float64, ethCap int64, btcCap int64) {
	repo.init = true
	repo.ratio = ratio
	repo.ethCap = ethCap
	repo.btcCap = btcCap
	repo.updateAt = time.Now()
}
