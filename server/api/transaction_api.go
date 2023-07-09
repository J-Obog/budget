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
	transactions, err := api.transactionManager.GetByAccount(r.Account.Id, r.Query.TransactionQuery())
	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(transactions)
}

func (api *TransactionAPI) CreateTransaction(r *rest.Request) *rest.Response {
	reqBody, err := r.Body.TransactionCreateBody()
	if err != nil {
		return buildBadRequestError()
	}

	if err := api.validateCreate(&reqBody, r.Account.Id); err != nil {
		return nil
	}

	if err := api.transactionManager.Create(r.Account.Id, reqBody); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) UpdateTransaction(r *rest.Request) *rest.Response {
	transaction, errResp := api.transactionCtx(r)
	if errResp != nil {
		return errResp
	}

	reqBody, err := r.Body.TransactionUpdateBody()
	if err != nil {
		return buildBadRequestError()
	}

	if err := api.validateUpdate(&reqBody, r.Account.Id); err != nil {
		return nil
	}

	if err := api.transactionManager.Update(&transaction, reqBody); err != nil {
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

func (api *TransactionAPI) validateCreate(reqBody *rest.TransactionCreateBody, accountId string) error {
	month := reqBody.Month
	year := reqBody.Year
	day := reqBody.Day
	categoryId := reqBody.CategoryId
	description := reqBody.Description

	if err := checkDate(month, day, year); err != nil {
		return err
	}

	if categoryId != nil {
		if err := checkCategoryExists(*categoryId, accountId, api.categoryManager); err != nil {
			return err
		}
	}

	if description != nil {
		if err := checkBudgetDesciption(*description); err != nil {
			return err
		}
	}

	return nil
}

func (api *TransactionAPI) validateUpdate(reqBody *rest.TransactionUpdateBody, accountId string) error {
	return api.validateCreate(&reqBody.TransactionCreateBody, accountId)
}
