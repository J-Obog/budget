package server

import (
	"github.com/J-Obog/paidoff/api"
)

type Server interface {
	Start(port int,
		accountAPI *api.AccountAPI,
		authAPI *api.AuthAPI,
		budgetAPI *api.BudgetAPI,
		categoryAPI *api.CategoryAPI,
		transactionAPI *api.TransactionAPI,
	) error
}
