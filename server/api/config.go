package api

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/manager"
)

type APIConfig struct {
	AccountAPI     *AccountAPI
	BudgetAPI      *BudgetAPI
	TransactionAPI *TransactionAPI
	CategoryAPI    *CategoryAPI
}

func CreateConfig(app *config.AppConfig) *APIConfig {
	managerCfg := manager.CreateConfig(app)

	return &APIConfig{
		AccountAPI:     NewAccountAPI(managerCfg.AccountManager),
		CategoryAPI:    NewCategoryAPI(managerCfg.CategoryManager, managerCfg.BudgetManager),
		TransactionAPI: NewTransactionAPI(managerCfg.TransactionManager, managerCfg.CategoryManager),
		BudgetAPI:      NewBudgetAPI(managerCfg.BudgetManager, managerCfg.TransactionManager),
	}
}
