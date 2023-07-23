package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type AccountAPI struct {
	accountManager *manager.AccountManager
}

func (api *AccountAPI) GetAccount(req *rest.Request) *rest.Response {
	return api.accountManager.GetByRequest(req)
}

func (api *AccountAPI) UpdateAccount(req *rest.Request) *rest.Response {
	return api.accountManager.UpdateByRequest(req)
}

func (api *AccountAPI) DeleteAccount(req *rest.Request) *rest.Response {
	return api.accountManager.DeleteByRequest(req)
}
