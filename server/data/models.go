package data

type BudgetType string

const (
	BudgetType_Income  BudgetType = "Income"
	BudgetType_Expense BudgetType = "Expense"
)

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string
	IsActivated bool  `json:"isActivated"`
	IsDeleted   bool  `json:"isDeleted"`
	CreatedAt   int64 `json:"createdAt"`
	UpdatedAt   int64 `json:"updatedAt"`
}

type Budget struct {
	Id         string  `json:"id"`
	AccountId  string  `json:"accountId"`
	CategoryId string  `json:"categoryId"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
	Projected  float64 `json:"projected"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
}

type BudgetMaterialized struct {
	Actual float64 `json:"actual"`
}

type Category struct {
	Id        string `json:"id"`
	AccountId string `json:"accountId"`
	Name      string `json:"name"`
	Color     uint   `json:"color"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt int64  `json:"createdAt"`
}

type Transaction struct {
	Id          string     `json:"id"`
	AccountId   string     `json:"accountId"`
	CategoryId  *string    `json:"categoryId"`
	Description *string    `json:"description"`
	Type        BudgetType `json:"budgetType"`
	Amount      float64    `json:"amount"`
	Month       int        `json:"month"`
	Day         int        `json:"day"`
	Year        int        `json:"year"`
	CreatedAt   int64      `json:"createdAt"`
	UpdatedAt   int64      `json:"updatedAt"`
}
