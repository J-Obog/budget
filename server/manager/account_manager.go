package manager

import (
	"github.com/J-Obog/paidoff/clock"
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
	account := req.Account
	account.UpdatedAt = manager.clock.Now()
	//account.Name = req.Name

	if err := manager.store.Update(*account); err != nil {
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
