package rest

import (
	"github.com/J-Obog/paidoff/types"
)

type BudgetQuery struct {
	Month types.Optional[int] `json:"month"`
	Year  types.Optional[int] `json:"year"`
}

type TransactionQuery struct {
	CreatedBefore types.Optional[int64]   `json:"createdBefore"`
	CreatedAfter  types.Optional[int64]   `json:"createdAfter"`
	AmountGte     types.Optional[float64] `json:"amountGte"`
	AmountLte     types.Optional[float64] `json:"amountLte"`
}

type AccountSetBody struct {
	Name string `json:"name"`
}

type CategorySetBody struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type CategoryUpdateBody struct {
	CategorySetBody
}

type CategoryCreateBody struct {
	CategorySetBody
}

type TransactionSetBody struct {
	CategoryId types.Optional[string] `json:"categoryId"`
	Note       types.Optional[string] `json:"note"`
	Amount     float64                `json:"amount"`
	Month      int                    `json:"month"`
	Day        int                    `json:"day"`
	Year       int                    `json:"year"`
}

type BudgetSetBody struct {
	CategoryId string  `json:"categoryId"`
	Projected  float64 `json:"projected"`
}

type BudgetCreateBody struct {
	BudgetSetBody
	Month int `json:"month"`
	Year  int `json:"year"`
}

type BudgetUpdateBody struct {
	BudgetSetBody
}
