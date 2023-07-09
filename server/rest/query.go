package rest

import "github.com/J-Obog/paidoff/data"

type Query map[string]any

type BudgetQuery struct {
	Month *int `json:"month"`
	Year  *int `json:"year"`
}

type TransactionQuery struct {
	CreatedBefore *int64          `json:"createdBefore"`
	CreatedAfter  *int64          `json:"createdAfter"`
	AmountGte     *float64        `json:"amountGte"`
	AmountLte     *float64        `json:"amountLte"`
	Categories    []data.Category `json:"categories"`
}

func (q Query) TransactionQuery() TransactionQuery {
	return TransactionQuery{}
}

func (q Query) BudgetQuery() BudgetQuery {
	return BudgetQuery{}
}
