package fliprepo

import "github.com/berkeleytrue/crypto-egg-go/internal/core/domain"

type flipRatio struct {
	ratio int64
}

func CreateFlipRepo() *flipRatio {
	return &flipRatio{
		ratio: 0,
	}
}

func (repo *flipRatio) Get() domain.Flippening {
	return domain.Flippening{Ratio: repo.ratio}
}

func (repo *flipRatio) Update(ratio int64) {
	repo.ratio = ratio
}
