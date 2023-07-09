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
	reqBody, err := r.Body.AccountUpdateBody()
	if err != nil {
		return buildBadRequestError()
	}

	if err := api.validateUpdate(reqBody); err != nil {
		return buildBadRequestError()
	}

	if err := api.accountManager.Update(r.Account, reqBody); err != nil {
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

func (api *AccountAPI) validateUpdate(reqBody rest.AccountUpdateBody) error {
	return checkAccountName(reqBody.Name)
}
