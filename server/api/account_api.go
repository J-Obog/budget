package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type AccountAPI struct {
	accountManager *manager.AccountManager
}

func (api *AccountAPI) GetAccount(r *rest.Request) *rest.Response {
	return buildOKResponse(r.Account)
}

func (api *AccountAPI) UpdateAccount(r *rest.Request) *rest.Response {
	if errResp := validateAccountUpdate(r); errResp != nil {
		return errResp
	}

	if err := api.accountManager.Update(r.Account, r.Body.AccountUpdateBody()); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *AccountAPI) DeleteAccount(r *rest.Request) *rest.Response {
	if err := api.accountManager.Delete(*r.Account); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}
