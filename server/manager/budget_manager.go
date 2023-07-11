package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type BudgetManager struct {
	store store.BudgetStore
	clock clock.Clock
	uid   uid.UIDProvider
}

func (manager *BudgetManager) CategoryInPeriod(id string, accountId string, month int, year int) (bool, error) {
	budgets, err := manager.filterByPeriod(accountId, month, year)
	if err != nil {
		return false, nil
	}

	for _, budget := range budgets {
		if budget.CategoryId == id {
			return true, err
		}
	}

	return false, nil
}

func (manager *BudgetManager) Get(id string, accountId string) (*data.Budget, error) {
	budget, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}
	if budget == nil || budget.AccountId != accountId {
		return nil, nil
	}

	return budget, nil
}

func (manager *BudgetManager) Filter(accountId string, q rest.BudgetQuery) ([]data.Budget, error) {
	return manager.filterByQuery(accountId, q)
}

func (manager *BudgetManager) Create(accountId string, req rest.BudgetCreateBody) error {
	now := manager.clock.Now()

	newBudget := data.Budget{
		Id:         manager.uid.GetId(),
		AccountId:  accountId,
		CategoryId: req.CategoryId,
		Month:      req.Month,
		Year:       req.Year,
		Projected:  req.Projected,
		Actual:     0,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	return manager.store.Insert(newBudget)
}

func (manager *BudgetManager) Update(existing *data.Budget, req rest.BudgetUpdateBody) error {
	existing.CategoryId = req.CategoryId
	existing.Month = req.Month
	existing.Year = req.Year
	existing.Projected = req.Projected
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *BudgetManager) Delete(id string) error {
	return manager.store.Delete(id)
}

func (manager *BudgetManager) filterByQuery(accountId string, query rest.BudgetQuery) ([]data.Budget, error) {
	budgets, err := manager.store.GetByAccount(accountId)
	if err != nil {
		return budgets, err
	}

	filtered := filter[data.Budget](budgets, func(b *data.Budget) bool {
		if query.Month != nil && b.Month != *query.Month {
			return false
		}

		if query.Year != nil && b.Year != *query.Year {
			return false
		}
		return true
	})

	return filtered, nil
}

func (manager *BudgetManager) filterByPeriod(accountId string, month int, year int) ([]data.Budget, error) {
	budgets, err := manager.store.GetByAccount(accountId)
	if err != nil {
		return budgets, err
	}

	filtered := filter[data.Budget](budgets, func(b *data.Budget) bool {
		return b.Month == month && b.Year == year
	})

	return filtered, nil
}
