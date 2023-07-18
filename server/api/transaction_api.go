package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type TransactionAPI struct {
	transactionManager *manager.TransactionManager
}

func (api *TransactionAPI) GetTransaction(req *rest.Request) *rest.Response {
	return api.transactionManager.GetByRequest(req)
}

func (api *TransactionAPI) GetTransactions(req *rest.Request) *rest.Response {
	return api.transactionManager.GetAllByRequest(req)
}

func (api *TransactionAPI) CreateTransaction(req *rest.Request) *rest.Response {
	return api.transactionManager.CreateByRequest(req)
}

func (api *TransactionAPI) UpdateTransaction(req *rest.Request) *rest.Response {
	return api.transactionManager.UpdateByRequest(req)
}

func (api *TransactionAPI) DeleteTransaction(req *rest.Request) *rest.Response {
	return api.transactionManager.DeleteByRequest(req)
}
