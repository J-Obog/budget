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

func (manager *AccountManager) Get(id string) (*data.Account, error) {
	return manager.store.Get(id)
}

func (manager *AccountManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.AccountUpdateBody)
	timestamp := manager.clock.Now()
	updateAccount(body, req.Account, timestamp)

	if manager.validate(res, body.Name); res.IsErr() {
		return
	}

	if err := manager.store.Update(*req.Account); err != nil {
		res.ErrInternal(err)
	}
}

func (manager *AccountManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	account := req.Account
	account.IsDeleted = true

	if err := manager.store.Update(*account); err != nil {
		res.ErrInternal(err)
	}
}

// TODO: return account name error instead of bad request
func (manager *AccountManager) validate(res *rest.Response, name string) {
	nameLen := len(name)

	if !(nameLen >= config.LimitMinAccountNameChars && nameLen <= config.LimitMaxAccountNameChars) {
		res.ErrBadRequest()
		return
	}
}
