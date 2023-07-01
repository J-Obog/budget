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

	transactions, err := api.transactionManager.GetByAccount(accountId)
	if err != nil {
		return buildServerError(err)
	}

	filter := NewFilter[data.Transaction]()
	filter.AddCheck(filterTransaction(q))
	filtered := filter.Filter(transactions)

	return buildOKResponse(filtered)
}

func (api *TransactionAPI) CreateTransaction(req *data.RestRequest) *data.RestResponse {
	createReq, err := getTransactionCreateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	//accountId := getAccountCtx(req).Id
	errRes := validateCategoryId(createReq.CategoryId, req, api.categoryManager)
	if errRes != nil {
		return errRes
	}

	if err := api.transactionManager.Create(createReq); err != nil {
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

	errRes = validateCategoryId(updateReq.CategoryId, req, api.categoryManager)
	if errRes != nil {
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
