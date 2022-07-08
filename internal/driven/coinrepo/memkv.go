package coinrepo

import (
	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/jinzhu/copier"
)

type MemKVS struct {
	idToCoin map[string]domain.Coin
	symToId  map[string]string
}

func NewMemKVS() *MemKVS {
	return &MemKVS{
		idToCoin: make(map[string]domain.Coin),
		symToId:  make(map[string]string),
	}
}

func (repo *MemKVS) Get(id string) domain.Coin {
	if value, ok := repo.idToCoin[id]; ok {
		return value
	}

	return domain.Coin{}
}

func (repo *MemKVS) GetBySymbol(sym string) domain.Coin {
	if value, ok := repo.symToId[sym]; ok {
		return repo.Get(value)
	}
	return domain.Coin{}
}

func (repo *MemKVS) GetAll() []domain.Coin {
	coins := make([]domain.Coin, 0, len(repo.idToCoin))

	for _, coin := range repo.idToCoin {
		coins = append(coins, coin)
	}

	return coins
}

func (repo *MemKVS) Update(id string, coin domain.Coin) domain.Coin {
	freshCoin := domain.Coin{}

	if value, ok := repo.idToCoin[id]; ok {
		// old coin, update freshCoin w/ old data
		freshCoin = value
	}

	copier.Copy(&freshCoin, &coin)

	repo.idToCoin[id] = freshCoin
	repo.symToId[coin.Symbol] = id

	return freshCoin

}
