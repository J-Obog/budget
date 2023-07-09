package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

func checkCategoryExists(categoryId *string, accountId string, catManager *manager.CategoryManager) *rest.Response {
	if categoryId == nil {
		return nil
	}

	cat, err := catManager.Get(*categoryId)
	if err != nil {
		return buildServerError(err)
	}
	if cat == nil || cat.AccountId != accountId {
		return buildBadRequestError()
	}

	return nil
}
