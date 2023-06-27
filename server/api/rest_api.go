package api

import (
	"net/http"

	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
)

type RestAPI struct {
	store store.Store
	clock clock.Clock
}

func (api *RestAPI) getAccount(req *data.RestRequest, res *data.RestResponse) (data.Account, error) {
	var account data.Account

	accountId, err := ICast[string](req.Meta["account_id"])

	if err != nil {
		return account, err
	}

	dbAccount, err := api.store.GetAccount(accountId)

	if err != nil {
		return account, err
	}

	if dbAccount == nil {
		res.Status = http.StatusNotFound
		return account, nil
	}

	return *dbAccount, nil
}

func (api *RestAPI) GetAccount(req *data.RestRequest, res *data.RestResponse) error {
	account, err := api.getAccount(req, res)

	if ResponseIsError(res.Status, err) {
		return err
	}

	res.Data = account
	res.Status = http.StatusOK
	return nil
}

func (api *RestAPI) UpdateAccount(req *data.RestRequest, res *data.RestResponse) error {
	account, err := api.getAccount(req, res)

	if ResponseIsError(res.Status, err) {
		return err
	}

	updateReq, err := FromJSON[data.AccountUpdateRequest](req.Body)

	if err != nil {
		return err
	}

	account.Name = updateReq.Name
	account.UpdatedAt = api.clock.Now()

	err = api.store.UpdateAccount(account)

	if err != nil {
		return err
	}

	res.Status = http.StatusOK
	return nil
}

func (api *RestAPI) DeleteAccount(req *data.RestRequest, res *data.RestResponse) error {
	account, err := api.getAccount(req, res)

	if ResponseIsError(res.Status, err) {
		return err
	}

	account.IsDeleted = true

	err = api.store.UpdateAccount(account)
	if err != nil {
		return err
	}

	res.Status = http.StatusOK
	return nil
}
