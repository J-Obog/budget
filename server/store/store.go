package store

import (
	"github.com/J-Obog/paidoff/data"
)

type Store interface {
	AccountStore2
	BudgetStore2
	CategoryStore2
	TransactionStore2
}

type AccountStore interface {
	Get(id string) (*data.Account, error)
	Insert(account data.Account) error
	Update(id string, update data.AccountUpdate, timestamp int64) (bool, error)
	SetDeleted(id string) (bool, error)
	Delete(id string) (bool, error)
	DeleteAll() error
}

type AccountStore2 interface {
	GetAccount(id string) (*data.Account, error)
	InsertAccount(account data.Account) error
	UpdateAccount(id string, update data.AccountUpdate, timestamp int64) (bool, error)
	SoftDeleteAccount(id string) (bool, error)
	DeleteAccount(id string) (bool, error)
	DeleteAllAccounts() error
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

type CategoryStore2 interface {
	GetCategory(id string, accountId string) (*data.Category, error)
	GetCategoryByName(accountId string, name string) (*data.Category, error)
	GetAllCategories(accountId string) ([]data.Category, error)
	InsertCategory(category data.Category) error
	UpdateCategory(id string, accountId string, update data.CategoryUpdate, timestamp int64) (bool, error)
	DeleteCategory(id string, accountId string) (bool, error)
	DeleteAllCategories() error
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

type BudgetStore2 interface {
	GetBudget(id string, accountId string) (*data.Budget, error)
	GetBudgetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error)
	GetBudgetByCategory(accountId string, categoryId string) ([]data.Budget, error)
	GetBudgetByFilter(accountId string, filter data.BudgetFilter) ([]data.Budget, error)
	InsertBudget(budget data.Budget) error
	UpdateBudget(id string, accountId string, update data.BudgetUpdate, timestamp int64) (bool, error)
	DeleteBudget(id string, accountId string) (bool, error)
	DeleteAllBudgets() error
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

type TransactionStore2 interface {
	GetTransaction(id string, accountId string) (*data.Transaction, error)
	GetTransactionByFilter(accountId string, filter data.TransactionFilter) ([]data.Transaction, error)
	GetTransactionByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error)
	InsertTransaction(transaction data.Transaction) error
	UpdateTransaction(id string, accountId string, update data.TransactionUpdate, timestamp int64) (bool, error)
	DeleteTransaction(id string, accountId string) (bool, error)
	DeleteAllTransactions() error
}
