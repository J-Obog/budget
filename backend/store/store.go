package store

import (
	"github.com/J-Obog/paidoff/data"
)

type AccountStore interface {
	Get(id string) (*data.Account, error)
	Insert(account data.Account) error
	Update(updated data.Account) (bool, error)
	SoftDelete(id string) (bool, error)
	Delete(id string) (bool, error)
	DeleteAll() error
}

type CategoryStore interface {
	Get(id string, accountId string) (*data.Category, error)
	GetByName(accountId string, name string) (*data.Category, error)
	GetAll(accountId string) ([]data.Category, error)
	Insert(category data.Category) error
	Update(updated data.Category) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}

type BudgetStore interface {
	Get(id string, accountId string) (*data.Budget, error)
	GetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error)
	GetByCategory(accountId string, categoryId string) ([]data.Budget, error)
	GetByPeriod(accountId string, month int, year int) ([]data.Budget, error)
	Insert(budget data.Budget) error
	Update(updated data.Budget) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}

type TransactionStore interface {
	Get(id string, accountId string) (*data.Transaction, error)
	GetByFilter(accountId string, filter data.TransactionFilter) ([]data.Transaction, error)
	GetByCategory(categoryId string, accountId string) ([]data.Transaction, error)
	GetByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error)
	Insert(transaction data.Transaction) error
	Update(updated data.Transaction) (bool, error)
	NullCategory(id string, accountId string) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}
