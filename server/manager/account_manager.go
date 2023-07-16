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

type accountValidateCommon struct {
	name string
}

func (manager *AccountManager) Get(id string) (*data.Account, error) {
	return manager.store.Get(id)
}

func (manager *AccountManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.AccountUpdateBody)

	validateCom := accountValidateCommon{name: body.Name}
	manager.validate(res, validateCom)
	if res.IsErr() {
		return
	}

	now := manager.clock.Now()
	updateAccount(body, req.Account, now)

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
func (manager *AccountManager) validate(res *rest.Response, validateCom accountValidateCommon) {
	nameLen := len(validateCom.name)

	if !(nameLen >= config.LimitMinAccountNameChars && nameLen <= config.LimitMaxAccountNameChars) {
		res.ErrBadRequest()
		return
	}
}
