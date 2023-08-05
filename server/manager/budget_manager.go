package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type BudgetManager struct {
	store        store.BudgetStore
	uuidProvider uuid.UuidProvider
	clock        clock.Clock
}

func (manager *BudgetManager) CheckCategoryNotInUse(categoryId string, accountId string) error {
	return nil
}
