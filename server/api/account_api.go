package api

import (
	"net/http"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
)

type AccountAPI struct {
	accountStore store.AccountStore
}

func (api *AccountAPI) GetAccount(req *data.RestRequest, res *data.RestResponse) error {
	err, accountId := getCurrentAccountId(req)

	if err != nil {
		return err
	}

	err, account := api.accountStore.Get(accountId)

	if err != nil {
		return err
	}

	if account != nil {
		res.Data = account
		return nil
	}

	res.Status = http.StatusNotFound
	return nil
}

func (api *AccountAPI) UpdateAccount(req *data.RestRequest, res *data.RestResponse) error {
	return nil
}

func (api *AccountAPI) DeleteAccount(req *data.RestRequest, res *data.RestResponse) error {
	return nil
}
