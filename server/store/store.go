package store

import (
	"github.com/J-Obog/paidoff/data"
)

type Store interface {
	AccountStore
	BudgetStore
	CategoryStore
	TransactionStore
}

type AccountStore interface {
	GetAccount(id string) (*data.Account, error)
	InsertAccount(account data.Account) error
	UpdateAccount(id string, update data.AccountUpdate, timestamp int64) (bool, error)
	SoftDeleteAccount(id string) (bool, error)
	DeleteAccount(id string) (bool, error)
	DeleteAllAccounts() error
}

type CategoryStore interface {
	GetCategory(id string, accountId string) (*data.Category, error)
	GetCategoryByName(accountId string, name string) (*data.Category, error)
	GetAllCategories(accountId string) ([]data.Category, error)
	InsertCategory(category data.Category) error
	UpdateCategory(id string, accountId string, update data.CategoryUpdate, timestamp int64) (bool, error)
	DeleteCategory(id string, accountId string) (bool, error)
	DeleteAllCategories() error
}

type BudgetStore interface {
	GetBudget(id string, accountId string) (*data.Budget, error)
	GetBudgetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error)
	GetBudgetsByCategory(accountId string, categoryId string) ([]data.Budget, error)
	GetBudgetsByFilter(accountId string, filter data.BudgetFilter) ([]data.Budget, error)
	InsertBudget(budget data.Budget) error
	UpdateBudget(id string, accountId string, update data.BudgetUpdate, timestamp int64) (bool, error)
	DeleteBudget(id string, accountId string) (bool, error)
	DeleteAllBudgets() error
}

type TransactionStore interface {
	GetTransaction(id string, accountId string) (*data.Transaction, error)
	GetTransactionsByFilter(accountId string, filter data.TransactionFilter) ([]data.Transaction, error)
	GetTransactionsByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error)
	InsertTransaction(transaction data.Transaction) error
	UpdateTransaction(id string, accountId string, update data.TransactionUpdate, timestamp int64) (bool, error)
	DeleteTransaction(id string, accountId string) (bool, error)
	DeleteAllTransactions() error
}
