package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
)

type TransactionAPI struct {
	transactionManager *manager.TransactionManager
	categoryManager    *manager.CategoryManager
}

func (api *TransactionAPI) GetTransaction(req *data.RestRequest) *data.RestResponse {
	transaction, errRes := checkTransaction(req, api.transactionManager)
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
	createReq, err := getTransactionCreateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	if errRes := validateCategoryId(createReq.CategoryId, req, api.categoryManager); errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Create(getAccountCtx(req).Id, createReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) UpdateTransaction(req *data.RestRequest) *data.RestResponse {
	transaction, errRes := checkTransaction(req, api.transactionManager)
	if errRes != nil {
		return errRes
	}

	updateReq, err := getTransactionUpdateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	if errRes = validateCategoryId(updateReq.CategoryId, req, api.categoryManager); errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Update(&transaction, updateReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *TransactionAPI) DeleteTransaction(req *data.RestRequest) *data.RestResponse {
	transaction, errRes := checkTransaction(req, api.transactionManager)
	if errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Delete(transaction.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}
