package manager

import (
	"github.com/J-Obog/paidoff/clock"
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
	transaction, err := manager.store.Get(req.ResourceId, req.Account.Get().Id)
	if err != nil {
		return rest.Err(err)
	}

	if transaction.Empty() {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Ok(transaction)
}

// TODO: convert timestamps in query to dates
func (manager *TransactionManager) GetAllByRequest(req *rest.Request) *rest.Response {
	query := req.Query.(rest.TransactionQuery)
	accountId := req.Account.Get().Id
	filter := getFilterForTransactionQuery(query)

	transactions, err := manager.store.GetBy(accountId, filter)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(transactions)
}

func (manager *TransactionManager) CreateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.TransactionSetBody)
	accountId := req.Account.Get().Id

	if err := manager.validateSet(accountId, body, false); err != nil {
		return rest.Err(err)
	}

	transaction := manager.getTransactionForCreate(accountId, body)

	if err := manager.store.Insert(transaction); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *TransactionManager) UpdateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.TransactionSetBody)
	accountId := req.Account.Get().Id
	transactionId := req.ResourceId

	timestamp := manager.clock.Now()
	update := getUpdateForTransactionUpdate(body)

	if err := manager.validateSet(accountId, body, true); err != nil {
		return rest.Err(err)
	}

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
	accountId := req.Account.Get().Id

	ok, err := manager.store.Delete(transactionId, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Success()
}

func (manager *TransactionManager) validateSet(accountId string, body rest.TransactionSetBody, isUpdate bool) error {
	return nil
}

func (manager *TransactionManager) getTransactionForCreate(accountId string, body rest.TransactionSetBody) data.Transaction {
	return data.Transaction{}
}

func getFilterForTransactionQuery(q rest.TransactionQuery) data.TransactionFilter {
	return data.TransactionFilter{}
}

func getUpdateForTransactionUpdate(body rest.TransactionSetBody) data.TransactionUpdate {
	return data.TransactionUpdate{}
}
