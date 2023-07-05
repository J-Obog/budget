package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
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

func (manager *TransactionManager) GetByAccount(accountId string, q data.TransactionQuery) ([]data.Transaction, error) {
	/*
		filter := NewFilter[data.Transaction]()
			filter.AddCheck(filterTransaction(q))
			filtered := filter.Filter(transactions)
	*/
	return manager.store.GetByAccount(accountId)
}

func (manager *TransactionManager) Create(accountId string, req data.TransactionCreateRequest) error {
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

func (manager *TransactionManager) Update(existing *data.Transaction, req data.TransactionUpdateRequest) error {
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
	return nil
}
