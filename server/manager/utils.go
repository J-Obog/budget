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

func newCategory(body rest.CategoryCreateBody, id string, accountId string, timestamp int64) data.Category {
	return data.Category{
		Id:        id,
		AccountId: accountId,
		Name:      body.Name,
		Color:     body.Color,
		UpdatedAt: timestamp,
		CreatedAt: timestamp,
	}
}

func updateCategory(body rest.CategoryUpdateBody, category *data.Category, timestamp int64) {
	category.Color = body.Color
	category.Name = body.Name
	category.UpdatedAt = timestamp
}

func updateBudget(body rest.BudgetUpdateBody, budget *data.Budget, timestamp int64) {
	budget.CategoryId = body.CategoryId
	budget.Month = body.Month
	budget.Year = body.Year
	budget.Projected = body.Projected
	budget.UpdatedAt = timestamp
}
