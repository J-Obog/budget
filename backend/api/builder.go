package api

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/manager"
)

type ApiService struct {
	AccountAPI     *AccountAPI
	BudgetAPI      *BudgetAPI
	TransactionAPI *TransactionAPI
	CategoryAPI    *CategoryAPI
}

func NewApiService(app *config.AppConfig) *ApiService {
	managerSvc := manager.NewManagerService(app)

	return &ApiService{
		AccountAPI:     NewAccountAPI(managerSvc.AccountManager),
		CategoryAPI:    NewCategoryAPI(managerSvc.CategoryManager, managerSvc.BudgetManager),
		TransactionAPI: NewTransactionAPI(managerSvc.TransactionManager, managerSvc.CategoryManager),
		BudgetAPI:      NewBudgetAPI(managerSvc.BudgetManager, managerSvc.TransactionManager, managerSvc.CategoryManager),
	}
}
