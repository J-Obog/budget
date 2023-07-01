package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
)

type AccountAPI struct {
	accountManager *manager.AccountManager
}

func (api *AccountAPI) GetAccount(req *data.RestRequest) *data.RestResponse {
	return buildOKResponse(getAccountCtx(req))
}

func (api *AccountAPI) UpdateAccount(req *data.RestRequest) *data.RestResponse {
	account := getAccountCtx(req)
	updateReq, err := getAccountUpdateBody(req)

	if err != nil {
		return buildServerError(err)
	}

	if err := api.accountManager.Update(&account, updateReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *AccountAPI) DeleteAccount(req *data.RestRequest) *data.RestResponse {
	if err := api.accountManager.Delete(getAccountCtx(req)); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}
