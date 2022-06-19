package gasrepo

import (
	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/jinzhu/copier"
)

type MemStore struct {
	symToGas map[string]domain.Gas
}

func CreateMemRepo() *MemStore {
	return &MemStore{
		symToGas: make(map[string]domain.Gas),
	}
}

func (repo *MemStore) Get() domain.Gas {
	if value, ok := repo.symToGas["eth"]; ok {
		return value
	}
	return domain.Gas{}
}

func (repo *MemStore) Update(update domain.Gas) domain.Gas {
	if gas, ok := repo.symToGas["eth"]; ok {
	  copier.Copy(&gas, &update)
	  return gas
	}

	return domain.Gas{}
}
