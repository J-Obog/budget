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

func (manager *AccountManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.AccountUpdateBody)
	timestamp := manager.clock.Now()
	update := data.AccountUpdate{
		Name: &body.Name, UpdatedAt: &timestamp,
	}

	if manager.validate(res, body.Name); res.IsErr() {
		return
	}

	if _, err := manager.store.Update(req.Account.Id, update); err != nil {
		res.ErrInternal(err)
	}
}

func (manager *AccountManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	update := data.AccountUpdate{
		IsDeleted: boolPtr(true),
	}

	if _, err := manager.store.Update(req.Account.Id, update); err != nil {
		res.ErrInternal(err)
	}
}

func (manager *AccountManager) validate(res *rest.Response, name string) {
	nameLen := len(name)

	if !(nameLen >= config.LimitMinAccountNameChars && nameLen <= config.LimitMaxAccountNameChars) {
		res.ErrInvalidAccountName()
		return
	}
}
