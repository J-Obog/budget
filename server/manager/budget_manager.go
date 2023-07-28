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

func (manager *BudgetManager) Get(id string, accountId string) (*data.Budget, error) {
	return manager.store.Get(id, accountId)
}

// TODO: implement, may make it get for period
func (manager *BudgetManager) GetAll() ([]data.Budget, error) {
	return []data.Budget{}, nil
}

func (manager *BudgetManager) Create(accountId string, body *rest.BudgetCreateBody) (data.Budget, error) {
	id := manager.uid.GetId()
	timestamp := manager.clock.Now()

	budget := data.Budget{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Month:      body.Month,
		Year:       body.Year,
		Projected:  body.Projected,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}

	err := manager.store.Insert(budget)

	return budget, err
}

func (manager *BudgetManager) Update(existing *data.Budget, body *rest.BudgetCreateBody) (bool, error) {
	timestamp := manager.clock.Now()

	update := data.BudgetUpdate{
		CategoryId: body.CategoryId,
		Projected:  body.Projected,
	}

	existing.CategoryId = update.CategoryId
	existing.Projected = body.Projected

	return manager.store.Update(existing.Id, existing.AccountId, update, timestamp)
}

func (manager *BudgetManager) Delete(id string, accountId string) (bool, error) {
	return manager.store.Delete(id, accountId)
}
