package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type TransactionManager struct {
	store store.TransactionStore
	clock clock.Clock
	uid   uid.UIDProvider
}

func (manager *TransactionManager) Get(id string) (*data.Transaction, error) {
	return manager.store.Get(id)
}

func (manager *TransactionManager) GetByAccount(accountId string) ([]data.Transaction, error) {
	return manager.store.GetByAccount(accountId)
}

func (manager *TransactionManager) Create(req data.TransactionCreateRequest) error {
	return nil
}

func (manager *TransactionManager) Update(existing *data.Transaction, req data.TransactionUpdateRequest) error {
	return nil
}

func (manager *TransactionManager) Delete(id string) error {
	return nil
}
