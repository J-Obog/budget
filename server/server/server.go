package server

import (
	resource "github.com/J-Obog/paidoff/resources"
)

type Server interface {
	Start(port int,
		accountResource *resource.AccountResource,
		authResource *resource.AuthResource,
		budgetResource *resource.BudgetResource,
		categoryResource *resource.CategoryResource,
		transactionResource *resource.TransactionResource,
	) error
}
