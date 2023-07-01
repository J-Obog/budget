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

func (manager *BudgetManager) Create(accountId string, req data.BudgetCreateRequest) error {
	now := manager.clock.Now()

	newBudget := data.Budget{
		Id:         manager.uid.GetId(),
		AccountId:  accountId,
		CategoryId: req.CategoryId,
		Name:       req.Name,
		Month:      req.Month,
		Year:       req.Year,
		Projected:  req.Projected,
		Actual:     req.Actual,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	return manager.store.Insert(newBudget)
}

func (manager *BudgetManager) Update(existing *data.Budget, req data.BudgetUpdateRequest) error {
	existing.CategoryId = req.CategoryId
	existing.Name = req.Name
	existing.Month = req.Month
	existing.Year = req.Year
	existing.Projected = req.Projected
	existing.Actual = req.Actual
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *BudgetManager) Delete(id string) error {
	return nil
}
