package store

import (
	"github.com/J-Obog/paidoff/data"
)

type AccountStore interface {
	Get(id string) (*data.Account, error)
	Insert(account data.Account) error
	Update(id string, update data.AccountUpdate, timestamp int64) (bool, error)
	SetDeleted(id string) (bool, error)
	Delete(id string) (bool, error)
	DeleteAll() error
}

type CategoryStore interface {
	Get(id string, accountId string) (*data.Category, error)
	GetByName(accountId string, name string) (*data.Category, error)
	GetAll(accountId string) ([]data.Category, error)
	Insert(category data.Category) error
	Update(id string, accountId string, update data.CategoryUpdate, timestamp int64) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}

type BudgetStore interface {
	Get(id string, accountId string) (*data.Budget, error)
	GetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error)
	GetByCategory(accountId string, categoryId string) ([]data.Budget, error)
	GetBy(accountId string, filter data.BudgetFilter) ([]data.Budget, error)
	Insert(budget data.Budget) error
	Update(id string, accountId string, update data.BudgetUpdate, timestamp int64) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}

type TransactionStore interface {
	Get(id string, accountId string) (*data.Transaction, error)
	GetBy(accountId string, filter data.TransactionFilter) ([]data.Transaction, error)
	GetByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error)
	Insert(transaction data.Transaction) error
	Update(id string, accountId string, update data.TransactionUpdate, timestamp int64) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}
