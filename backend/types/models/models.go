package data

type Account struct {
	Id        string
	Name      string
	Email     string
	Password  string
	IsDeleted bool
	CreatedAt int64
	UpdatedAt int64
}

type BudgetType string

const (
	BudgetType_Income  BudgetType = "INCOME"
	BudgetType_Expense BudgetType = "EXPENSE"
)

type Budget struct {
	Id         string
	AccountId  string
	CategoryId string
	Month      int
	Year       int
	Amount     float64
	CreatedAt  int64
	UpdatedAt  int64
}

type Category struct {
	Id         string
	AccountId  string
	Name       string
	Color      uint
	BudgetType BudgetType
	UpdatedAt  int64
	CreatedAt  int64
}

type Transaction struct {
	Id          string
	AccountId   string
	CategoryId  *string
	Description *string
	Type        BudgetType
	Amount      float64
	Timestamp   int64
	CreatedAt   int64
	UpdatedAt   int64
}
