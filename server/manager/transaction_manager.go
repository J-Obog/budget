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

func (manager *TransactionManager) Get(id string) (*data.Transaction, error) {
	return manager.store.Get(id)
}

func (manager *TransactionManager) Filter(accountId string, q rest.TransactionQuery) ([]data.Transaction, error) {
	filtered := make([]data.Transaction, 0)

	transactions, err := manager.store.GetByAccount(accountId)
	if err != nil {
		return filtered, err
	}

	for _, transaction := range transactions {
		if q.CreatedBefore != nil && transaction.CreatedAt >= *q.CreatedBefore {
			continue
		}

		if q.CreatedAfter != nil && transaction.CreatedAt <= *q.CreatedAfter {
			continue
		}

		if q.AmountGte != nil && transaction.Amount < *q.AmountGte {
			continue
		}

		if q.AmountLte != nil && transaction.Amount > *q.AmountLte {
			continue
		}

		filtered = append(filtered, transaction)

	}

	return filtered, nil
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
