package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/validation"
)

type TransactionAPI struct {
	transactionManager *manager.TransactionManager
	categoryManager    *manager.CategoryManager
}

func (api *TransactionAPI) getTransactionCtx(req *data.RestRequest) (data.Transaction, *data.RestResponse) {
	transaction, err := api.transactionManager.Get(req.UrlParams["transactionId"].(string))

	if err != nil {
		return data.Transaction{}, buildServerError(err)
	}
	if transaction == nil || transaction.AccountId != getAccountCtx(req).Id {
		return data.Transaction{}, buildBadRequestError()
	}

	return *transaction, nil
}

func (api *TransactionAPI) GetTransaction(req *data.RestRequest) *data.RestResponse {
	transaction, errRes := api.getTransactionCtx(req)
	if errRes != nil {
		return errRes
	}

	return buildOKResponse(transaction)
}

func (api *TransactionAPI) GetTransactions(req *data.RestRequest) *data.RestResponse {
	q, err := getTransactionGetQuery(req)

	if err != nil {
		return buildServerError(err)
	}

	accountId := getAccountCtx(req).Id

	transactions, err := api.transactionManager.GetByAccount(accountId, q)
	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(transactions)
}

func (api *TransactionAPI) CreateTransaction(req *data.RestRequest) *data.RestResponse {
	accountId := getAccountCtx(req).Id

	if err := validation.ValidateTransactionCreateReq(req.Body); err != nil {
		return buildBadRequestError()
	}

	createReq, err := getTransactionCreateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	if errRes := checkCategoryExists(createReq.CategoryId, accountId, api.categoryManager); errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Create(accountId, createReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) UpdateTransaction(req *data.RestRequest) *data.RestResponse {
	if err := validation.ValidateTransactionUpdateReq(req.Body); err != nil {
		return buildBadRequestError()
	}

	transaction, errRes := api.getTransactionCtx(req)
	if errRes != nil {
		return errRes
	}

	updateReq, err := getTransactionUpdateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	if errRes := checkCategoryExists(updateReq.CategoryId, transaction.AccountId, api.categoryManager); errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Update(&transaction, updateReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) DeleteTransaction(req *data.RestRequest) *data.RestResponse {
	transaction, errRes := api.getTransactionCtx(req)
	if errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Delete(transaction.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}
