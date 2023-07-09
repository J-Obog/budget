package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

func validateAccountName(name string) error {
	return nil
}

func checkCategoryExists(categoryId string, accountId string, catManager *manager.CategoryManager) *rest.Response {
	cat, err := catManager.Get(categoryId)
	if err != nil {
		return buildServerError(err)
	}
	if cat == nil || cat.AccountId != accountId {
		return buildBadRequestError()
	}

	return nil
}

func isValidDate(month int, day int, year int) error {
	return nil
}
