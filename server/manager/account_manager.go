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

func (manager *AccountManager) Update(existing *data.Account, req rest.AccountUpdateBody) error {
	existing.UpdatedAt = manager.clock.Now()
	existing.Name = req.Name

	return manager.store.Update(*existing)
}

func (manager *AccountManager) Delete(account data.Account) error {
	account.IsDeleted = true
	return manager.store.Update(account)
}
