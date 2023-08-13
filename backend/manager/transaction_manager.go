package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type TransactionManager struct {
	store        store.TransactionStore
	uuidProvider uuid.UuidProvider
	clock        clock.Clock
}

func NewTransactionManager(
	store store.TransactionStore,
	uuidProvider uuid.UuidProvider,
	clock clock.Clock,
) *TransactionManager {
	return &TransactionManager{
		store:        store,
		uuidProvider: uuidProvider,
		clock:        clock,
	}
}

func (manager *TransactionManager) Get(id string, accountId string) (*data.Transaction, error) {
	return manager.store.Get(id, accountId)
}

func (manager *TransactionManager) Create(
	accountId string,
	body rest.TransactionCreateBody,
) (data.Transaction, error) {
	timestamp := manager.clock.Now()
	uuid := manager.uuidProvider.GetUuid()

	newTransaction := data.Transaction{
		Id:         uuid,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     round(body.Amount, 2),
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}

	if err := manager.store.Insert(newTransaction); err != nil {
		return data.Transaction{}, err
	}

	return newTransaction, nil
}

func (manager *TransactionManager) Update(
	existing *data.Transaction,
	body rest.TransactionUpdateBody,
) (bool, error) {
	existing.CategoryId = body.CategoryId
	existing.Note = body.Note
	existing.Type = body.Type
	existing.Amount = round(body.Amount, 2)
	existing.Month = body.Month
	existing.Day = body.Day
	existing.Year = body.Year
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *TransactionManager) Delete(id string, accountId string) (bool, error) {
	return manager.store.Delete(id, accountId)
}

func (manager *TransactionManager) GetTotalForPeriodCategory(
	accountId string,
	categoryId string,
	month int,
	year int,
) (float64, error) {
	total := 0.00
	transactions, err := manager.store.GetByPeriodCategory(accountId, categoryId, month, year)

	if err != nil {
		return total, err
	}

	for _, transaction := range transactions {
		netMove := transaction.Amount
		if transaction.Type == data.BudgetType_Expense {
			netMove *= -1
		}

		total += netMove
	}

	return round(total, 2), nil
}
