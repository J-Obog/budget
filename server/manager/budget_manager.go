package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type BudgetManager struct {
	store        store.BudgetStore
	uuidProvider uuid.UuidProvider
	clock        clock.Clock
}

func NewBudgetManager(
	store store.BudgetStore,
	uuidProvider uuid.UuidProvider,
	clock clock.Clock,
) *BudgetManager {
	return &BudgetManager{
		store:        store,
		uuidProvider: uuidProvider,
		clock:        clock,
	}
}

func (manager *BudgetManager) Get(id string, accountId string) (*data.Budget, error) {
	return nil, nil
}

func (manager *BudgetManager) Create(accountId string, createReq rest.BudgetCreateBody) (data.Budget, error) {
	return data.Budget{}, nil
}

func (manager *BudgetManager) Update(updated *data.Budget, updateReq rest.BudgetUpdateBody) error {
	return nil
}

func (manager *BudgetManager) Delete(id string, accountId string) error {
	return nil
}

func (manager *BudgetManager) CheckCategoryNotInUse(categoryId string, accountId string) error {
	return nil
}

func (manager *BudgetManager) CheckCategoryNotInPeriod(categoryId string, accountId string, month int, year int) error {
	return nil
}
