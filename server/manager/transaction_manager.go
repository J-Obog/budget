package manager

import (
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
	transaction, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}
	if transaction == nil || transaction.AccountId != accountId {
		return nil, nil
	}

	return transaction, nil
}

func (manager *TransactionManager) Filter(accountId string, q rest.TransactionQuery) ([]data.Transaction, error) {
	return manager.filterByQuery(accountId, q)
}

func (manager *TransactionManager) Create(accountId string, req rest.TransactionCreateBody) error {
	now := manager.clock.Now()

	newTransaction := data.Transaction{
		Id:          manager.uid.GetId(),
		AccountId:   accountId,
		CategoryId:  req.CategoryId,
		Description: req.Description,
		Amount:      req.Amount,
		Month:       req.Month,
		Day:         req.Day,
		Year:        req.Year,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return manager.store.Insert(newTransaction)
}

func (manager *TransactionManager) Update(existing *data.Transaction, req rest.TransactionUpdateBody) error {
	existing.CategoryId = req.CategoryId
	existing.Description = req.Description
	existing.Amount = req.Amount
	existing.Month = req.Month
	existing.Day = req.Day
	existing.Year = req.Year
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *TransactionManager) Delete(id string) error {
	return manager.store.Delete(id)
}

func (manager *TransactionManager) filterByQuery(accountId string, query rest.TransactionQuery) ([]data.Transaction, error) {
	transactions, err := manager.store.GetByAccount(accountId)
	if err != nil {
		return transactions, err
	}

	filtered := filter[data.Transaction](transactions, func(t *data.Transaction) bool {
		if query.CreatedBefore != nil && t.CreatedAt >= *query.CreatedBefore {
			return false
		}

		if query.CreatedAfter != nil && t.CreatedAt <= *query.CreatedAfter {
			return false
		}

		if query.AmountGte != nil && t.Amount < *query.AmountGte {
			return false
		}

		if query.AmountLte != nil && t.Amount > *query.AmountLte {
			return false
		}
		return true

	})

	return filtered, nil
}
