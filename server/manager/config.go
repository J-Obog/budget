package manager

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/store"
)

type ManagerConfig struct {
	BudgetManager      *BudgetManager
	CategoryManager    *CategoryManager
	TransactionManager *TransactionManager
	AccountManager     *AccountManager
}

// TODO: replace mocks with impls
func CreateConfig(app *config.AppConfig) *ManagerConfig {
	storeConfig := store.CreateConfig(app)
	queue := queue.CreateConfig(app)
	clock := new(mocks.Clock)
	uid := new(mocks.UIDProvider)

	return &ManagerConfig{
		BudgetManager:      NewBudgetManager(storeConfig.BudgetStore, clock, uid),
		TransactionManager: NewTransactionManager(storeConfig.TransactionStore, clock, uid),
		CategoryManager:    NewCategoryManager(storeConfig.CategoryStore, clock, uid, queue),
		AccountManager:     NewAccountManager(storeConfig.AccountStore, clock),
	}
}
