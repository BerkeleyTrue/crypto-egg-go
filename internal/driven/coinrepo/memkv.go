package coinrepo

import (
	"encoding/json"
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/jinzhu/copier"
)

type MemKVS struct {
	kvs     map[string][]byte
	symToId map[string]string
}

func NewMemKVS() *MemKVS {
	return &MemKVS{
		kvs:     make(map[string][]byte),
		symToId: make(map[string]string),
	}
}

func jsonToCoin(value []byte) (domain.Coin, error) {
	coin := domain.Coin{}

	err := json.Unmarshal(value, &coin)
	if err != nil {
		return domain.Coin{}, fmt.Errorf("Failed to parse value")
	}

	return coin, nil
}

func (repo *MemKVS) Get(id string) (domain.Coin, error) {
	if value, ok := repo.kvs[id]; ok {
		coin, err := jsonToCoin(value)

		if err != nil {
			return domain.Coin{}, fmt.Errorf("Failed to parse value for id %s", id)
		}

		return coin, nil
	}

	return domain.Coin{}, nil
}

func (repo *MemKVS) GetBySymbol(sym string) (domain.Coin, error) {
  if value, ok := repo.symToId[sym]; ok {
    return repo.Get(value)
  }
  return domain.Coin{}, nil
}

func (repo *MemKVS) GetAll() ([]domain.Coin, error) {
	coins := make([]domain.Coin, len(repo.kvs))
	for _, value := range repo.kvs {
		coin, err := jsonToCoin(value)

		if err != nil {
			return nil, err
		}

		coins = append(coins, coin)
	}
	return coins, nil
}

func (repo *MemKVS) Update(id string, coin domain.Coin) (domain.Coin, error) {
	freshCoin := domain.Coin{}

	if value, ok := repo.kvs[id]; ok {
		// old coin, update freshCoin w/ old data
		err := json.Unmarshal(value, &freshCoin)

		if err != nil {
			// do we care that we can't parse old value?
			return domain.Coin{}, fmt.Errorf("Failed to parse value for id %s", id)
		}
	}

	copier.Copy(&freshCoin, &coin)

	marshelled, err := json.Marshal(freshCoin)

	if err != nil {
		return domain.Coin{}, fmt.Errorf("failed to marshel coin")
	}

	repo.kvs[id] = marshelled
	repo.symToId[coin.Symbol] = id

	return freshCoin, nil

}
