package rest

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/gorilla/schema"
)

type BudgetQuery struct {
	Month *int `schema:"month"`
	Year  *int `schema:"year"`
}

type TransactionQuery struct {
	StartDate *data.Date `schema:"startDate"`
	EndDate   *data.Date `schema:"endDate"`
	MinAmount *float64   `schema:"minAmount"`
	MaxAmount *float64   `schema:"maxAmount"`
}

func ParseQuery[T any](queryMap map[string][]string) (T, error) {
	var t T

	decoder := schema.NewDecoder()

	if err := decoder.Decode(&t, queryMap); err != nil {
		return t, &RestError{Msg: "invalid value for parsing query"}
	}

	return t, nil
}
