package models

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
