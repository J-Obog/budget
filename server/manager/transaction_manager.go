package manager

import (
	"math"

	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type TransactionManager struct {
	store         store.TransactionStore
	categoryStore store.CategoryStore
	clock         clock.Clock
	uid           uid.UIDProvider
}

func (manager *TransactionManager) GetByRequest(req *rest.Request) *rest.Response {
	transaction, err := manager.store.Get(req.ResourceId, req.Account.Id)
	if err != nil {
		return rest.Err(err)
	}

	if transaction == nil {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Ok(transaction)
}

func (manager *TransactionManager) GetAllByRequest(req *rest.Request) *rest.Response {
	query := req.Query.(rest.TransactionQuery)
	accountId := req.Account.Id
	filter := manager.getFilterForTransactionQuery(query)

	transactions, err := manager.store.GetBy(accountId, filter)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(transactions)
}

func (manager *TransactionManager) CreateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.TransactionCreateBody)
	accountId := req.Account.Id

	if err := manager.validateCreate(accountId, body); err != nil {
		return rest.Err(err)
	}

	transaction := manager.getTransactionForCreate(accountId, body)

	if err := manager.store.Insert(transaction); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *TransactionManager) UpdateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.TransactionUpdateBody)
	accountId := req.Account.Id
	transactionId := req.ResourceId

	if err := manager.validateUpdate(transactionId, accountId, body); err != nil {
		return rest.Err(err)
	}

	timestamp := manager.clock.Now()
	update := getUpdateForTransactionUpdate(body)

	ok, err := manager.store.Update(transactionId, accountId, update, timestamp)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Success()
}

func (manager *TransactionManager) DeleteByRequest(req *rest.Request) *rest.Response {
	transactionId := req.ResourceId
	accountId := req.Account.Id

	ok, err := manager.store.Delete(transactionId, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Success()
}

// TODO: look into validation for budget type
func (manager *TransactionManager) validateUpdate(id string, accountId string, body rest.TransactionUpdateBody) error {
	existing, err := manager.store.Get(id, accountId)
	if err != nil {
		return err
	}

	if existing == nil {
		return rest.ErrInvalidTransactionId
	}

	if err := manager.checkDate(body.Month, body.Day, body.Year); err != nil {
		return err
	}

	if body.Note != nil {
		if existing.Note == nil || (*existing.Note != *body.Note) {
			if err := manager.checkNote(*body.Note); err != nil {
				return err
			}
		}
	}

	if body.CategoryId != nil {
		if existing.CategoryId == nil || (*existing.CategoryId != *body.CategoryId) {
			if err := manager.checkCategoryExists(*body.CategoryId, accountId); err != nil {
				return err
			}
		}
	}

	return nil
}

func (manager *TransactionManager) validateCreate(accountId string, body rest.TransactionCreateBody) error {
	if err := manager.checkDate(body.Month, body.Day, body.Year); err != nil {
		return err
	}

	if body.Note != nil {
		if err := manager.checkNote(*body.Note); err != nil {
			return err
		}
	}

	if body.CategoryId != nil {
		if err := manager.checkCategoryExists(*body.CategoryId, accountId); err != nil {
			return err
		}
	}

	return nil
}

func (manager *TransactionManager) checkDate(month int, day int, year int) error {
	d := data.NewDate(month, day, year)

	if ok := manager.clock.IsDateValid(d); !ok {
		return rest.ErrInvalidDate
	}

	return nil
}

func (manager *TransactionManager) checkNote(note string) error {
	if len(note) > config.LimitMaxTransactionNoteChars {
		return rest.ErrInvalidTransactionNote
	}

	return nil
}

func (manager *TransactionManager) checkCategoryExists(categoryId string, accountId string) error {
	category, err := manager.categoryStore.Get(categoryId, accountId)
	if err != nil {
		return err
	}

	if category == nil {
		return rest.ErrInvalidCategoryId
	}

	return nil
}

func (manager *TransactionManager) getTransactionForCreate(accountId string, body rest.TransactionCreateBody) data.Transaction {
	id := manager.uid.GetId()
	timestamp := manager.clock.Now()

	return data.Transaction{
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
}

// TODO: get default date bounds from config
func (manager *TransactionManager) getFilterForTransactionQuery(q rest.TransactionQuery) data.TransactionFilter {
	filter := data.TransactionFilter{
		Before:      data.NewDate(1, 1, 1902),
		After:       data.NewDate(1, 1, math.MaxInt),
		GreaterThan: math.MaxFloat64,
		LessThan:    0,
	}

	if q.CreatedBefore != nil {
		filter.Before = manager.clock.DateFromStamp(*q.CreatedBefore)
	}

	if q.CreatedAfter != nil {
		filter.After = manager.clock.DateFromStamp(*q.CreatedAfter)
	}

	if q.AmountGte != nil {
		filter.GreaterThan = *q.AmountGte
	}

	if q.AmountLte != nil {
		filter.LessThan = *q.AmountLte
	}

	return filter
}

func getUpdateForTransactionUpdate(body rest.TransactionUpdateBody) data.TransactionUpdate {
	return data.TransactionUpdate{
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
	}
}
