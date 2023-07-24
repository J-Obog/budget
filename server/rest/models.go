package rest

import (
	"github.com/J-Obog/paidoff/data"
)

type BudgetQuery struct {
	Month *int `json:"month"`
	Year  *int `json:"year"`
}

type TransactionQuery struct {
	CreatedBefore *int64   `json:"createdBefore"`
	CreatedAfter  *int64   `json:"createdAfter"`
	AmountGte     *float64 `json:"amountGte"`
	AmountLte     *float64 `json:"amountLte"`
}

type AccountUpdateBody struct {
	Name string `json:"name"`
}

type CategoryUpdateBody struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type CategoryCreateBody struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type TransactionUpdateBody struct {
	CategoryId *string         `json:"categoryId"`
	Note       *string         `json:"note"`
	Type       data.BudgetType `json:"budgetType"`
	Amount     float64         `json:"amount"`
	Month      int             `json:"month"`
	Day        int             `json:"day"`
	Year       int             `json:"year"`
}

type TransactionCreateBody struct {
	CategoryId *string         `json:"categoryId"`
	Note       *string         `json:"note"`
	Type       data.BudgetType `json:"budgetType"`
	Amount     float64         `json:"amount"`
	Month      int             `json:"month"`
	Day        int             `json:"day"`
	Year       int             `json:"year"`
}

type BudgetCreateBody struct {
	CategoryId string  `json:"categoryId"`
	Projected  float64 `json:"projected"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
}

type BudgetUpdateBody struct {
	CategoryId string  `json:"categoryId"`
	Projected  float64 `json:"projected"`
}
