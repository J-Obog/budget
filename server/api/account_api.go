package api

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
)

type AccountAPI struct {
	accountStore store.AccountStore
	clock        clock.Clock
}

func NewAccountAPI(
	accountStore store.AccountStore,
	clock clock.Clock,
) *AccountAPI {
	return &AccountAPI{
		accountStore: accountStore,
		clock:        clock,
	}
}

func (api *AccountAPI) Get(req *rest.Request) *rest.Response {
	accountId := testAccountId

	account, err := api.accountStore.Get(accountId)
	if err != nil {
		return rest.Err(err)
	}

	if account == nil {
		return rest.Err(rest.ErrInvalidAccountId)
	}

	return rest.Ok(account)
}

func (api *AccountAPI) Update(req *rest.Request) *rest.Response {
	accountId := testAccountId
	timestamp := api.clock.Now()

	body, err := rest.ParseBody[rest.AccountUpdateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	account, err := api.accountStore.Get(accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(account, body); err != nil {
		return rest.Err(err)
	}

	update := data.AccountUpdate{
		Name: body.Name,
	}

	account.Name = update.Name

	ok, err := api.accountStore.Update(accountId, update, timestamp)
	if err != nil {
		return rest.Err(rest.ErrInternalServer)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidAccountId)
	}

	return rest.Ok(account)
}

func (api *AccountAPI) Delete(req *rest.Request) *rest.Response {
	accountId := testAccountId

	ok, err := api.accountStore.Delete(accountId)
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
