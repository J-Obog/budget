package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type TransactionAPI struct {
	transactionManager *manager.TransactionManager
}

func (api *TransactionAPI) GetTransaction(req *rest.Request, res *rest.Response) {
	api.transactionManager.GetByRequest(req, res)
}

func (api *TransactionAPI) GetTransactions(req *rest.Request, res *rest.Response) {
	api.transactionManager.GetAllByRequest(req, res)
}

func (api *TransactionAPI) CreateTransaction(req *rest.Request, res *rest.Response) {
	api.transactionManager.CreateByRequest(req, res)
}

func (api *TransactionAPI) UpdateTransaction(req *rest.Request, res *rest.Response) {
	api.transactionManager.UpdateByRequest(req, res)
}

func (api *TransactionAPI) DeleteTransaction(req *rest.Request, res *rest.Response) {
	api.transactionManager.DeleteByRequest(req, res)
}
