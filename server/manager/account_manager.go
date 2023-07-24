package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
)

type AccountManager struct {
	store store.AccountStore
	clock clock.Clock
}

func (manager *AccountManager) GetByRequest(req *rest.Request) *rest.Response {
	return rest.Ok(req.Account)
}

func (manager *AccountManager) UpdateByRequest(req *rest.Request) *rest.Response {
	accountId := req.Account.Id

	body, err := rest.ParseBody[rest.AccountUpdateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	if err := manager.validateUpdate(body, req.Account); err != nil {
		return rest.Err(err)
	}

	update := getUpdateForAccountUpdate(body)
	timestamp := manager.clock.Now()

	if _, err := manager.store.Update(accountId, update, timestamp); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *AccountManager) DeleteByRequest(req *rest.Request) *rest.Response {
	if _, err := manager.store.SetDeleted(req.Account.Id); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *AccountManager) validateUpdate(body rest.AccountUpdateBody, account *data.Account) error {
	if body.Name != account.Name {
		if err := manager.checkAccountName(body.Name); err != nil {
			return err
		}
	}

	return nil
}

func (manager *AccountManager) checkAccountName(name string) error {
	nameLen := len(name)

	if nameLen < config.LimitMinAccountNameChars || nameLen > config.LimitMaxAccountNameChars {
		return rest.ErrInvalidAccountName
	}

	return nil
}

func getUpdateForAccountUpdate(body rest.AccountUpdateBody) data.AccountUpdate {
	return data.AccountUpdate{
		Name: body.Name,
	}
}
