package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
)

func boolPtr(v bool) *bool {
	b := new(bool)
	*b = v
	return b
}

func newBudget(body rest.BudgetCreateBody, id string, accountId string, timestamp int64, period int64) data.Budget {
	return data.Budget{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Period:     period,
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

func newTransaction(body rest.TransactionCreateBody, id string, accountId string, timestamp int64, period int64) data.Transaction {
	return data.Transaction{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Amount:     body.Amount,
		Period:     period,
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
	budget.Projected = body.Projected
	budget.UpdatedAt = timestamp
}

func updateTransaction(body rest.TransactionUpdateBody, transaction *data.Transaction, timestamp int64, period int64) {
	transaction.CategoryId = body.CategoryId
	transaction.Note = body.Note
	transaction.Amount = body.Amount
	transaction.Period = period
	transaction.UpdatedAt = timestamp
}

func updateAccount(body rest.AccountUpdateBody, account *data.Account, timestamp int64) {
	account.Name = body.Name
	account.UpdatedAt = timestamp
}
