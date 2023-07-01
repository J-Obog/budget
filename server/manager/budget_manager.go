package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type BudgetManager struct {
	store store.BudgetStore
	clock clock.Clock
	uid   uid.UIDProvider
}

func (manager *BudgetManager) Get(id string) (*data.Budget, error) {
	return manager.store.Get(id)
}

func (manager *BudgetManager) GetByAccount(accountId string) ([]data.Budget, error) {
	return manager.store.GetByAccount(accountId)
}

func (manager *BudgetManager) Create(req data.BudgetCreateRequest) error {
	return nil
}

func (manager *BudgetManager) Update(existing *data.Budget, req data.BudgetUpdateRequest) error {
	return nil
}

func (manager *BudgetManager) Delete(id string) error {
	return nil
}
