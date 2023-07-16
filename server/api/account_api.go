package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type AccountAPI struct {
	accountManager *manager.AccountManager
}

func (api *AccountAPI) GetAccount(req *rest.Request, res *rest.Response) {
	res.Ok(req.Account)
}

func (api *AccountAPI) UpdateAccount(req *rest.Request, res *rest.Response) {
	api.accountManager.UpdateByRequest(req, res)
}

func (api *AccountAPI) DeleteAccount(req *rest.Request, res *rest.Response) {
	api.accountManager.DeleteByRequest(req, res)
}
