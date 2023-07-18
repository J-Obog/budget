package data

import "github.com/J-Obog/paidoff/types"

type Transaction struct {
	Id         string                 `json:"id"`
	AccountId  string                 `json:"accountId"`
	CategoryId types.Optional[string] `json:"categoryId"`
	Note       types.Optional[string] `json:"note"`
	Type       BudgetType             `json:"budgetType"`
	Amount     float64                `json:"amount"`
	Period     int64                  `json:"period"`
	CreatedAt  int64                  `json:"createdAt"`
	UpdatedAt  int64                  `json:"updatedAt"`
}

type TransactionFilter struct {
	Before      Date
	After       Date
	GreaterThan float64
	LessThan    float64
}

type TransactionUpdate struct {
	CategoryId types.Optional[string]
	Note       types.Optional[string]
	Type       BudgetType
	Amount     float64
	Period     int64
}
