package api

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type AccountAPI struct {
	accountManager *manager.AccountManager
}

func (api *AccountAPI) Get(req *rest.Request) *rest.Response {
	return rest.Ok(req.Account)
}

func (api *AccountAPI) Update(req *rest.Request) *rest.Response {
	body, err := rest.ParseBody[rest.AccountUpdateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(req.Account, body); err != nil {
		return rest.Err(err)
	}

	ok, err := api.accountManager.Update(req.Account, body)
	if err != nil {
		return rest.Err(rest.ErrInternalServer)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidAccountId)
	}

	return rest.Ok(req.Account)
}

func (api *AccountAPI) Delete(req *rest.Request) *rest.Response {
	ok, err := api.accountManager.Delete(req.Account.Id)
	if err != nil {
		return rest.Err(rest.ErrInternalServer)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidAccountId)
	}

	return rest.Success()
}

func (api *AccountAPI) validateUpdate(existing *data.Account, body rest.AccountUpdateBody) error {
	if existing == nil {
		return rest.ErrInvalidAccountId
	}

	if body.Name != existing.Name {
		if err := api.checkAccountName(body.Name); err != nil {
			return err
		}
	}

	return nil
}

func (api *AccountAPI) checkAccountName(name string) error {
	nameLen := len(name)

	if nameLen < config.LimitMinAccountNameChars || nameLen > config.LimitMaxAccountNameChars {
		return rest.ErrInvalidAccountName
	}

	return nil
}
