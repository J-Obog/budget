package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
)

func checkCategory(req *data.RestRequest, manager *manager.CategoryManager) (data.Category, *data.RestResponse) {
	id := getCategoryId(req)
	category, err := manager.Get(id)

	if err != nil {
		return data.Category{}, buildServerError(err)
	}

	if category == nil {
		return data.Category{}, buildNotFoundError()
	}

	if category.AccountId != getAccountCtx(req).Id {
		return data.Category{}, buildForbiddenError()
	}

	return *category, nil
}

func checkTransaction(req *data.RestRequest, manager *manager.TransactionManager) (data.Transaction, *data.RestResponse) {
	id := getTransactionId(req)
	transaction, err := manager.Get(id)

	if err != nil {
		return data.Transaction{}, buildServerError(err)
	}

	if transaction == nil {
		return data.Transaction{}, buildNotFoundError()
	}

	if transaction.AccountId != getAccountCtx(req).Id {
		return data.Transaction{}, buildForbiddenError()
	}

	return *transaction, nil
}

func checkBudget(req *data.RestRequest, manager *manager.BudgetManager) (data.Budget, *data.RestResponse) {
	id := getBudgetId(req)
	budget, err := manager.Get(id)

	if err != nil {
		return data.Budget{}, buildServerError(err)
	}

	if budget == nil {
		return data.Budget{}, buildNotFoundError()
	}

	if budget.AccountId != getAccountCtx(req).Id {
		return data.Budget{}, buildForbiddenError()
	}

	return *budget, nil
}

func validateCategoryId(categoryId *string, req *data.RestRequest, manager *manager.CategoryManager) *data.RestResponse {
	if categoryId != nil {
		_, errRes := checkCategory(req, manager)
		return errRes
	}

	return nil
}
