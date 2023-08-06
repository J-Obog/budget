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
		rest.Err(err)
	}

	return rest.Ok(account)
}

func (api *AccountAPI) Update(req *rest.Request) *rest.Response {
	accountId := testAccountId

	body, err := rest.ParseBody[rest.AccountUpdateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	account, err := api.accountManager.Get(accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(account, body); err != nil {
		return rest.Err(err)
	}

	if err := api.accountManager.Update(account, body); err != nil {
		return rest.Err(err)
	}

	return rest.Ok(account)
}

func (api *AccountAPI) Delete(req *rest.Request) *rest.Response {
	accountId := testAccountId

	if err := api.accountManager.Delete(accountId); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (api *AccountAPI) validateUpdate(existing *data.Account, body rest.AccountUpdateBody) error {
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
