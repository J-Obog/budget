package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type TransactionAPI struct {
	transactionManager *manager.TransactionManager
	categoryManager    *manager.CategoryManager
}

func (api *TransactionAPI) GetTransaction(r *rest.Request) *rest.Response {
	transaction, errRes := api.transactionCtx(r)
	if errRes != nil {
		return errRes
	}

	return buildOKResponse(transaction)
}

func (api *TransactionAPI) GetTransactions(r *rest.Request) *rest.Response {
	if errResp := api.validateGets(r); errResp != nil {
		return errResp
	}

	transactions, err := api.transactionManager.GetByAccount(r.Account.Id, r.Query.TransactionQuery())
	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(transactions)
}

func (api *TransactionAPI) CreateTransaction(r *rest.Request) *rest.Response {
	if errResp := api.validateCreate(r); errResp != nil {
		return errResp
	}

	if err := api.transactionManager.Create(r.Account.Id, r.Body.TransactionCreateBody()); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) UpdateTransaction(r *rest.Request) *rest.Response {
	transaction, errResp := api.transactionCtx(r)
	if errResp != nil {
		return errResp
	}

	if errResp := api.validateUpdate(r); errResp != nil {
		return errResp
	}

	if err := api.transactionManager.Update(&transaction, r.Body.TransactionUpdateBody()); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) DeleteTransaction(r *rest.Request) *rest.Response {
	transaction, errRes := api.transactionCtx(r)
	if errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Delete(transaction.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) validateCreate(r *rest.Request) *rest.Response {
	return nil
}

func (api *TransactionAPI) validateUpdate(r *rest.Request) *rest.Response {
	return nil
}

func (api *TransactionAPI) validateGets(r *rest.Request) *rest.Response {
	return nil
}

func (api *TransactionAPI) transactionCtx(r *rest.Request) (data.Transaction, *rest.Response) {
	transaction, err := api.transactionManager.Get(r.Params.TransactionId())

	if err != nil {
		return data.Transaction{}, buildServerError(err)
	}
	if transaction == nil || transaction.AccountId != r.Account.Id {
		return data.Transaction{}, buildBadRequestError()
	}

	return *transaction, nil
}
