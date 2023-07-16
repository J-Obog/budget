package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
)

func newBudget(body rest.BudgetCreateBody, id string, accountId string, timestamp int64) data.Budget {
	return data.Budget{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Month:      body.Month,
		Year:       body.Year,
		Projected:  body.Projected,
		Actual:     0,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}
}

func updateBudget(body rest.BudgetUpdateBody, budget *data.Budget, timestamp int64) {
	budget.CategoryId = body.CategoryId
	budget.Month = body.Month
	budget.Year = body.Year
	budget.Projected = body.Projected
	budget.UpdatedAt = timestamp
}
