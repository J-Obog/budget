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

func NewBudgetManager(store store.BudgetStore, clock clock.Clock, uid uid.UIDProvider) *BudgetManager {
	return &BudgetManager{
		store: store,
		clock: clock,
		uid:   uid,
	}
}

func (manager *BudgetManager) Get(id string, accountId string) (*data.Budget, error) {
	return manager.store.Get(id, accountId)
}

func (manager *BudgetManager) GetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error) {
	return manager.store.GetByPeriodCategory(accountId, categoryId, month, year)
}

func (manager *BudgetManager) GetByCategory(accountId string, categoryId string) ([]data.Budget, error) {
	return manager.store.GetByCategory(accountId, categoryId)
}

func (manager *BudgetManager) GetByQuery(accountId string, query rest.BudgetQuery) ([]data.Budget, error) {
	filter := data.BudgetFilter{
		Month: manager.clock.CurrentMonth(),
		Year:  manager.clock.CurrentYear(),
	}

	if query.Month != nil {
		filter.Month = *query.Month
	}

	if query.Year != nil {
		filter.Year = *query.Year
	}

	return manager.store.GetBy(accountId, filter)
}

func (manager *BudgetManager) Create(accountId string, body rest.BudgetCreateBody) (data.Budget, error) {
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

func (manager *BudgetManager) Update(existing *data.Budget, body rest.BudgetUpdateBody) (bool, error) {
	timestamp := manager.clock.Now()

	update := data.BudgetUpdate{
		CategoryId: body.CategoryId,
		Projected:  body.Projected,
	}

	existing.CategoryId = update.CategoryId
	existing.Projected = body.Projected
	existing.UpdatedAt = timestamp

	return manager.store.Update(existing.Id, existing.AccountId, update, timestamp)
}

func (manager *BudgetManager) Delete(id string, accountId string) (bool, error) {
	return manager.store.Delete(id, accountId)
}
