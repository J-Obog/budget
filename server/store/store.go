package store

import "github.com/J-Obog/paidoff/data"

type AccountStore interface {
	Get(id string) (*data.Account, error)
	Insert(account data.Account) error
	Update(account data.Account) error
	Delete(id string) error
	DeleteAll() error
}

type CategoryStore interface {
	Get(id string) (*data.Category, error)
	GetByAccount(accountId string) ([]data.Category, error)
	Insert(category data.Category) error
	Update(category data.Category) error
	Delete(id string) error
	DeleteAll() error
}

type BudgetStore interface {
	Get(id string) (*data.Budget, error)
	GetByAccount(accountId string) ([]data.Budget, error)
	Insert(budget data.Budget) error
	Update(budget data.Budget) error
	Delete(id string) error
	DeleteAll() error
}

type TransactionStore interface {
	Get(id string) (*data.Transaction, error)
	GetByAccount(accountId string) ([]data.Transaction, error)
	Insert(transaction data.Transaction) error
	Update(transaction data.Transaction) error
	Delete(id string) error
	DeleteAll() error
}
