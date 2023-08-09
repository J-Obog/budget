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

func NewAccountManager(store store.AccountStore, clock clock.Clock) *AccountManager {
	return &AccountManager{
		store: store,
		clock: clock,
	}
}

func (manager *AccountManager) Get(id string) (*data.Account, error) {
	account, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (manager *AccountManager) Update(updated *data.Account, updateReq rest.AccountUpdateBody) error {
	return nil
}

func (manager *AccountManager) DeleteForRequest(id string) error {
	ok, err := manager.store.Delete(id)
	if err != nil {
		return err
	}

	if !ok {
		return rest.ErrInvalidAccountId
	}

	return nil
}
