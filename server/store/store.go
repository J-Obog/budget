package store

import "github.com/J-Obog/paidoff/data"

type AccountStore interface {
	Get(id string) (*data.Account, error)
	Insert(account data.Account) error
	Update(id string, update data.AccountUpdate) (bool, error)
	Delete(id string) (bool, error)
	DeleteAll() error
}

type CategoryStore interface {
	Get(id string, accountId string) (*data.Category, error)
	GetBy(filter data.CategoryFilter) (data.CategoryList, error)
	Insert(category data.Category) error
	Update(id string, update data.CategoryUpdate) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}

type BudgetStore interface {
	Get(id string, accountId string) (*data.Budget, error)
	GetBy(filter data.BudgetFilter) (data.BudgetList, error)
	Insert(budget data.Budget) error
	Update(id string, update data.BudgetUpdate) (bool, error)
	Delete(id string, accountId string) (bool, error)
	DeleteAll() error
}

type TransactionStore interface {
	Get(id string, accountId string) (*data.Transaction, error)
	GetBy(filter data.TransactionFilter) (data.TransactionList, error)
	Insert(transaction data.Transaction) error
	Update(id string, update data.TransactionUpdate) (bool, error)
	Delete(id string) (bool, error)
	DeleteAll() error
}
