package models

type Transaction struct {
	Id          string
	AccountId   string
	BudgetId    string
	Description *string
	Amount      float64
	Month       int
	Day         int
	Year        int
	CreatedAt   int
	UpdatedAt   int
}
