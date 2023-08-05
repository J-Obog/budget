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
	account, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, rest.ErrInvalidAccountId
	}

	return account, nil
}

func (manager *AccountManager) Update(updated *data.Account, body rest.AccountUpdateBody) error {
	return nil
}

func (manager *AccountManager) Delete(id string) error {
	ok, err := manager.store.Delete(id)
	if err != nil {
		return err
	}

	if !ok {
		return rest.ErrInvalidAccountId
	}

	return nil
}
