package store

import "github.com/J-Obog/paidoff/data"

type Store interface {
	GetAccount(id string) (*data.Account, error)
	InsertAccount(account data.Account) error
	UpdateAccount(account data.Account) error
	DeleteAccount(id string) error

	GetCategory(id string) (*data.Category, error)
	GetCategories(accountId string) ([]data.Category, error)
	InsertCategory(category data.Category) error
	UpdateCategory(category data.Category) error
	DeleteCategory(id string) error

	GetBudget(id string) (*data.Budget, error)
	GetBudgets(accountId string) ([]data.Budget, error)
	InsertBudget(budget data.Budget) error
	UpdateBudget(budget data.Budget) error
	DeleteBudget(id string) error

	GetTransaction(id string) (*data.Transaction, error)
	GetTransactions(accountId string) ([]data.Transaction, error)
	InsertTransaction(transaction data.Transaction) error
	UpdateTransaction(transaction data.Transaction) error
	DeleteTransaction(id string) error

	Flush() error
}
