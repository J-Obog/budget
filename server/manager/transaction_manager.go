package manager

import (
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

func (manager *TransactionManager) GetByRequest(req *rest.Request, res *rest.Response) {
	transaction := manager.getTransaction(res, req.ResourceId, req.Account.Id)

	if res.IsErr() {
		return
	}

	res.Ok(transaction)
}

// TODO: convert timestamps in query to dates
func (manager *TransactionManager) GetAllByRequest(req *rest.Request, res *rest.Response) {
	query := req.Query.(rest.TransactionQuery)

	filter := data.TransactionFilter{
		GreaterThan: query.AmountGte,
		LessThan:    query.AmountLte,
	}

	transactions, err := manager.store.GetBy(filter)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(transactions)
}

func (manager *TransactionManager) CreateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.TransactionCreateBody)
	timestamp := manager.clock.Now()
	id := manager.uid.GetId()
	newTransaction := newTransaction(body, id, req.Account.Id, timestamp)

	if manager.validate(res, body.Note, body.Month, body.Day, body.Year, *body.CategoryId, req.Account.Id); res.IsErr() {
		return
	}

	if err := manager.store.Insert(newTransaction); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *TransactionManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	now := manager.clock.Now()
	body := req.Body.(rest.TransactionUpdateBody)

	transaction := manager.getTransaction(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	if manager.validate(res, body.Note, body.Month, body.Day, body.Year, *body.CategoryId, req.Account.Id); res.IsErr() {
		return
	}

	updateTransaction(body, transaction, now)

	if err := manager.store.Update(*transaction); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *TransactionManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	manager.getTransaction(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	if err := manager.store.Delete(req.ResourceId); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *TransactionManager) getTransaction(res *rest.Response, id string, accountId string) *data.Transaction {
	transaction, err := manager.store.Get(id, accountId)

	if err != nil {
		res.ErrInternal(err)
		return nil
	}

	if transaction == nil {
		res.ErrTransactionNotFound()
		return nil
	}

	return transaction
}

func (manager *TransactionManager) validate(res *rest.Response, note *string, month int, day int, year int, categoryId string, accountId string) {
	if note != nil {
		noteLen := len(*note)
		if noteLen > config.LimitMaxTransactionNoteChars {
			res.ErrInvalidTransactionNote()
			return
		}
	}

	if ok := isDateValid(month, day, year); !ok {
		res.ErrInvalidDate()
		return
	}

	category, err := manager.categoryStore.Get(categoryId, accountId)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if category == nil {
		res.ErrCategoryNotFound()
		return
	}
}
