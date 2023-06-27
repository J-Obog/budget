package data

type BudgetType uint

const (
	BudgetType_INCOME  BudgetType = 0
	BudgetType_EXPENSE BudgetType = 1
)

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string
	IsActivated bool  `json:"isActivated"`
	CreatedAt   int64 `json:"createdAt"`
	UpdatedAt   int64 `json:"updatedAt"`
}

type Budget struct {
	Id         string     `json:"id"`
	AccountId  string     `json:"accountId"`
	CategoryId *string    `json:"categoryId"`
	Name       string     `json:"name"`
	Type       BudgetType `json:"type"`
	Month      int        `json:"month"`
	Year       int        `json:"year"`
	Projected  float64    `json:"projected"`
	Actual     float64    `json:"actual"`
	CreatedAt  int        `json:"createdAt"`
	UpdatedAt  int        `json:"updatedAt"`
}

type Category struct {
	Id        string `json:"id"`
	AccountId string `json:"accountId"`
	Name      string `json:"name"`
	Color     int    `json:"color"`
	UpdatedAt int    `json:"updatedAt"`
	CreatedAt int    `json:"createdAt"`
}

type Transaction struct {
	Id          string  `json:"id"`
	AccountId   string  `json:"accountId"`
	BudgetId    string  `json:"budgetId"`
	Description *string `json:"description"`
	Amount      float64 `json:"amount"`
	Month       int     `json:"month"`
	Day         int     `json:"day"`
	Year        int     `json:"year"`
	CreatedAt   int     `json:"createdAt"`
	UpdatedAt   int     `json:"updatedAt"`
}
