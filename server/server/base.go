package server

import "github.com/J-Obog/paidoff/manager"

type BaseServer struct {
	accountManager     *manager.AccountManager
	budgetManager      *manager.AccountManager
	categoryManager    *manager.AccountManager
	transactionManager *manager.AccountManager
}
