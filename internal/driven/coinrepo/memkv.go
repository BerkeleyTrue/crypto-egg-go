package coinrepo

import (
	"encoding/json"
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/jinzhu/copier"
)

type MemKVS struct {
	kvs map[string][]byte
}

func NewMemKVS() *MemKVS {
	return &MemKVS{
		kvs: make(map[string][]byte),
	}
}

func (repo *MemKVS) Get(id string) (domain.Coin, error) {
	if value, ok := repo.kvs[id]; ok {
		coin := domain.Coin{}
		err := json.Unmarshal(value, &coin)

		if err != nil {
			return domain.Coin{}, fmt.Errorf("Failed to parse value for id %s", id)
		}

		return coin, nil
	}

	return domain.Coin{}, nil
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

	return freshCoin, nil

}
