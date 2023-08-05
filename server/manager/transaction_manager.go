package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type TransactionManager struct {
	store        store.TransactionStore
	uuidProvider uuid.UuidProvider
	clock        clock.Clock
}

func (manager *TransactionManager) Get(id string, accountId string) error {
	return nil
}

func (manager *TransactionManager) GetByPeriodCategory(
	accountId string,
	categoryId string,
	month int,
	year int,
) ([]data.Transaction, error) {
	return nil, nil
}
