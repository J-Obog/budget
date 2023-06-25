package store

import "github.com/J-Obog/paidoff/data"

type Store interface {
	GetAccount(id string) (*data.Account, error)
	InsertAccount(account data.Account) error
	UpdateAccount(account data.Account) error
	DeleteAccount(id string) error

	GetCategory(id string) (*data.Category, error)
	InsertCategory(account data.Category) error
	UpdateCategory(account data.Category) error
	DeleteCategory(id string) error

	GetBudget(id string) (*data.Budget, error)
	InsertBudget(account data.Budget) error
	UpdateBudget(account data.Budget) error
	DeleteBudget(id string) error

	GetTransaction(id string) (*data.Transaction, error)
	InsertTransaction(account data.Transaction) error
	UpdateTransaction(account data.Transaction) error
	DeleteTransaction(id string) error

	Flush() error
}
