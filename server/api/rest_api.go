package api

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
)

type RestAPI struct {
	store store.Store
	clock clock.Clock
}

func (api *RestAPI) GetAccount(req *data.RestRequest) data.RestResponse {
	account := getAccount(req)
	return buildOKResponse(account)
}

func (api *RestAPI) UpdateAccount(req *data.RestRequest) data.RestResponse {
	account := getAccount(req)

	if updateReq, err := FromJSON[data.AccountUpdateRequest](req.Body); err != nil {
		return buildServerError(err)
	} else {
		account.Name = updateReq.Name
		account.UpdatedAt = api.clock.Now()
	}

	if err := api.store.UpdateAccount(account); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *RestAPI) DeleteAccount(req *data.RestRequest) data.RestResponse {
	account := getAccount(req)
	account.IsDeleted = true

	if err := api.store.UpdateAccount(account); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(account)
}
