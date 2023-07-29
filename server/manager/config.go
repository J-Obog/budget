package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type ManagerConfig struct {
	BudgetManager      *BudgetManager
	CategoryManager    *CategoryManager
	TransactionManager *TransactionManager
	AccountManager     *AccountManager
}

func CreateConfig(app *config.AppConfig) *ManagerConfig {
	storeConfig := store.CreateConfig(app)
	queue := queue.CreateConfig(app)
	clock := clock.CreateConfig(app)
	uid := uid.CreateConfig(app)

	return &ManagerConfig{
		BudgetManager:      NewBudgetManager(storeConfig.BudgetStore, clock, uid),
		TransactionManager: NewTransactionManager(storeConfig.TransactionStore, clock, uid),
		CategoryManager:    NewCategoryManager(storeConfig.CategoryStore, clock, uid, queue),
		AccountManager:     NewAccountManager(storeConfig.AccountStore, clock),
	}
}
