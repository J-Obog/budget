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

func (manager *AccountManager) UpdateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.AccountSetBody)
	accountId := req.Account.Get().Id

	if err := manager.validateSet(body); err != nil {
		return rest.Err(err)
	}

	update := data.AccountUpdate{Name: body.Name}
	timestamp := manager.clock.Now()

	if _, err := manager.store.Update(accountId, update, timestamp); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *AccountManager) DeleteByRequest(req *rest.Request) *rest.Response {
	if _, err := manager.store.UpdateDeleted(req.Account.Get().Id, true); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *AccountManager) validateSet(body rest.AccountSetBody) error {
	nameLen := len(body.Name)

	if !(nameLen >= config.LimitMinAccountNameChars && nameLen <= config.LimitMaxAccountNameChars) {
		return rest.ErrInvalidAccountName
	}

	return nil
}
