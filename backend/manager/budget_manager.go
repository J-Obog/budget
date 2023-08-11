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
	return manager.store.Get(id, accountId)
}

func (manager *BudgetManager) Create(accountId string, reqBody rest.BudgetCreateBody) (data.Budget, error) {
	timestamp := manager.clock.Now()
	uuid := manager.uuidProvider.GetUuid()

	newBudget := data.Budget{
		Id:         uuid,
		AccountId:  accountId,
		CategoryId: reqBody.CategoryId,
		Month:      reqBody.Month,
		Year:       reqBody.Year,
		Projected:  reqBody.Projected,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}

	if err := manager.store.Insert(newBudget); err != nil {
		return data.Budget{}, err
	}

	return newBudget, nil
}

func (manager *BudgetManager) Update(existing *data.Budget, body rest.BudgetUpdateBody) (bool, error) {
	existing.CategoryId = body.CategoryId
	existing.Projected = body.Projected
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *BudgetManager) Delete(id string, accountId string) (bool, error) {
	return manager.store.Delete(id, accountId)
}

func (manager *BudgetManager) CategoryIsNotUsed(categoryId string, accountId string) (bool, error) {
	budgets, err := manager.store.GetByCategory(accountId, categoryId)
	if err != nil {
		return false, err
	}

	return len(budgets) == 0, nil
}

func (manager *BudgetManager) CategoryIsUniqueForPeriod(
	categoryId string,
	accountId string,
	month int,
	year int,
) (bool, error) {
	budget, err := manager.store.GetByPeriodCategory(accountId, categoryId, month, year)
	if err != nil {
		return false, err
	}

	return budget == nil, nil
}
