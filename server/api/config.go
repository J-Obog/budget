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
	managerConfig := manager.CreateConfig(app)

	return &APIConfig{}
}
