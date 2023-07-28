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
	return manager.store.Get(id, accountId)
}

func (manager *TransactionManager) GetByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error) {
	return manager.store.GetByPeriodCategory(accountId, categoryId, month, year)
}

// TODO: implement
func (manager *TransactionManager) GetAll(req *rest.Request) {

}

func (manager *TransactionManager) Create(accountId string, body *rest.TransactionCreateBody) (data.Transaction, error) {
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

func (manager *TransactionManager) Update(existing *data.Transaction, body *rest.TransactionUpdateBody) (bool, error) {
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

	return manager.store.Update(existing.Id, existing.AccountId, update, timestamp)
}

func (manager *TransactionManager) Delete(id string, accountId string) (bool, error) {
	return manager.store.Delete(id, accountId)
}
