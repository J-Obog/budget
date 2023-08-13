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

func NewAccountAPI(accountManager *manager.AccountManager) *AccountAPI {
	return &AccountAPI{
		accountManager: accountManager,
	}
}

func (api *AccountAPI) Get(req *rest.Request) *rest.Response {
	accountId := testAccountId
	account, err := api.accountManager.Get(accountId)

	if err != nil {
		return rest.Err(err)
	}

	if account == nil {
		return rest.Err(rest.ErrInvalidAccountId)
	}

	return rest.Ok(account)
}

func (api *AccountAPI) Update(req *rest.Request) *rest.Response {
	var body rest.AccountUpdateBody

	accountId := testAccountId

	if err := req.Body.To(&body); err != nil {
		return rest.Err(err)
	}

	account, err := api.accountManager.Get(accountId)

	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(account, body); err != nil {
		return rest.Err(err)
	}

	ok, err := api.accountManager.Update(account, body)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidAccountId)
	}

	return rest.Ok(account)
}

func (api *AccountAPI) Delete(req *rest.Request) *rest.Response {
	accountId := testAccountId
	ok, err := api.accountManager.SoftDelete(accountId)

	if err != nil {
		return rest.Err(err)
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

	if err := api.checkAccountName(body.Name); err != nil {
		return err
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
