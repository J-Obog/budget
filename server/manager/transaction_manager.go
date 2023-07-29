package manager

import (
	"math"

	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type TransactionManager struct {
	store store.TransactionStore
	clock clock.Clock
	uid   uid.UIDProvider
}

func (manager *TransactionManager) Get(id string, accountId string) (*data.Transaction, error) {
	return manager.store.Get(id, accountId)
}

func (manager *TransactionManager) GetByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error) {
	return manager.store.GetByPeriodCategory(accountId, categoryId, month, year)
}

func (manager *TransactionManager) GetByQuery(accountId string, query rest.TransactionQuery) ([]data.Transaction, error) {
	filter := data.TransactionFilter{
		Before:      data.NewDate(12, 31, math.MaxInt),
		After:       data.NewDate(1, 1, math.MinInt),
		GreaterThan: 0.00,
		LessThan:    math.MaxFloat64,
	}

	if query.StartDate != nil {
		filter.After = *query.StartDate
	}

	if query.EndDate != nil {
		filter.Before = *query.EndDate
	}

	if query.MinAmount != nil {
		filter.GreaterThan = *query.MinAmount
	}

	if query.MaxAmount != nil {
		filter.LessThan = *query.MaxAmount
	}

	return manager.store.GetBy(accountId, filter)
}

func (manager *TransactionManager) Create(accountId string, body rest.TransactionCreateBody) (data.Transaction, error) {
	id := manager.uid.GetId()
	timestamp := manager.clock.Now()

	transaction := data.Transaction{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}

	err := manager.store.Insert(transaction)

	return transaction, err
}

func (manager *TransactionManager) Update(existing *data.Transaction, body rest.TransactionUpdateBody) (bool, error) {
	timestamp := manager.clock.Now()
	update := data.TransactionUpdate{
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
	}

	existing.CategoryId = update.CategoryId
	existing.Note = update.Note
	existing.Type = update.Type
	existing.Amount = update.Amount
	existing.Month = update.Month
	existing.Day = update.Day
	existing.Year = update.Year
	existing.UpdatedAt = timestamp

	return manager.store.Update(existing.Id, existing.AccountId, update, timestamp)
}

func (manager *TransactionManager) Delete(id string, accountId string) (bool, error) {
	return manager.store.Delete(id, accountId)
}
