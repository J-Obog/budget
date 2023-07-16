package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
)

// TODO: implement
func isDateValid(month int, day int, year int) bool {
	return false
}

func newBudget(body rest.BudgetCreateBody, id string, accountId string, timestamp int64) data.Budget {
	return data.Budget{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Month:      body.Month,
		Year:       body.Year,
		Projected:  body.Projected,
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

func newTransaction(body rest.TransactionCreateBody, id string, accountId string, timestamp int64) data.Transaction {
	return data.Transaction{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
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

func updateTransaction(body rest.TransactionUpdateBody, transaction *data.Transaction, timestamp int64) {
	transaction.CategoryId = body.CategoryId
	transaction.Note = body.Note
	transaction.Amount = body.Amount
	transaction.Month = body.Month
	transaction.Day = body.Day
	transaction.Year = body.Year
	transaction.UpdatedAt = timestamp
}

func updateAccount(body rest.AccountUpdateBody, account *data.Account, timestamp int64) {
	account.Name = body.Name
	account.UpdatedAt = timestamp
}

func filter[T any](items []T, filterFn func(t *T) bool) []T {
	end := 0
	for _, item := range items {
		ok := filterFn(&item)
		if ok {
			items[end] = item
			end += 1
		}
	}

	return items[:end-1]
}

func find[T any](items []T, findFn func(t *T) bool) bool {
	for _, item := range items {
		ok := findFn(&item)
		if ok {
			return true
		}
	}

	return false
}
