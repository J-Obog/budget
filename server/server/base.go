package server

import "github.com/J-Obog/paidoff/manager"

type BaseServer struct {
	accountManager     *manager.AccountManager
	budgetManager      *manager.BudgetManager
	categoryManager    *manager.CategoryManager
	transactionManager *manager.TransactionManager
}
