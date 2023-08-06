package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type TransactionManager struct {
	store        store.TransactionStore
	uuidProvider uuid.UuidProvider
	clock        clock.Clock
}

func NewTransactionManager(
	store store.TransactionStore,
	uuidProvider uuid.UuidProvider,
	clock clock.Clock,
) *TransactionManager {
	return &TransactionManager{
		store:        store,
		uuidProvider: uuidProvider,
		clock:        clock,
	}
}

func (manager *TransactionManager) Get(id string, accountId string) (*data.Transaction, error) {
	return nil, nil
}

func (manager *TransactionManager) Create(accountId string, createReq rest.TransactionCreateBody) (data.Transaction, error) {
	return data.Transaction{}, nil
}

func (manager *TransactionManager) Update(updated *data.Transaction, updateReq rest.TransactionUpdateBody) error {
	return nil
}

func (manager *TransactionManager) Delete(id string, accountId string) error {
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
