package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type TransactionManager struct {
	store           store.TransactionStore
	categoryManager *CategoryManager
	clock           clock.Clock
	uid             uid.UIDProvider
}

func (manager *TransactionManager) Get(id string) (*data.Transaction, error) {
	transaction, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (manager *TransactionManager) GetByRequest(req *rest.Request, res *rest.Response) {
	transaction := manager.getTransactionByAccount(res, req.ResourceId, req.Account.Id)

	if res.IsErr() {
		return
	}

	res.Ok(transaction)
}

func (manager *TransactionManager) GetAllByRequest(req *rest.Request, res *rest.Response) {
	query := req.Query.(rest.TransactionQuery)

	transactions, err := manager.store.GetByAccount(req.Account.Id)
	if err != nil {
		res.ErrInternal(err)
		return
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

	res.Ok(filtered)
}

func (manager *TransactionManager) CreateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.TransactionCreateBody)
	timestamp := manager.clock.Now()
	id := manager.uid.GetId()
	newTransaction := newTransaction(body, id, req.Account.Id, timestamp)

	if manager.validate(res, body.Description, body.Month, body.Day, body.Year, *body.CategoryId, req.Account.Id); res.IsErr() {
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

	transaction := manager.getTransactionByAccount(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	if manager.validate(res, body.Description, body.Month, body.Day, body.Year, *body.CategoryId, req.Account.Id); res.IsErr() {
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
	manager.getTransactionByAccount(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	if err := manager.store.Delete(req.ResourceId); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *TransactionManager) getTransactionByAccount(res *rest.Response, id string, accountId string) *data.Transaction {
	transaction, err := manager.Get(id)

	if err != nil {
		res.ErrInternal(err)
		return nil
	}

	if transaction == nil || transaction.AccountId != accountId {
		res.ErrTransactionNotFound()
		return nil
	}

	return transaction
}

// TODO: Check date is valid
// TODO: Check if description is valid
func (manager *TransactionManager) validate(res *rest.Response, description *string, month int, day int, year int, categoryId string, accountId string) {
	ok, err := manager.categoryManager.Exists(categoryId, accountId)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if !ok {
		res.ErrCategoryNotFound()
		return
	}
}
