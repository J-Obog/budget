package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
)

type CategoryAPI struct {
	categoryManager *manager.CategoryManager
}

func (api *CategoryAPI) getCategoryCtx(req *data.RestRequest) (data.Category, *data.RestResponse) {
	category, err := api.categoryManager.Get(req.UrlParams["categoryId"].(string))

	if err != nil {
		return data.Category{}, buildServerError(err)
	}
	if category == nil || category.AccountId != getAccountCtx(req).Id {
		return data.Category{}, buildBadRequestError()
	}

	return *category, nil
}

func (api *CategoryAPI) GetCategory(req *data.RestRequest) *data.RestResponse {
	category, errRes := api.getCategoryCtx(req)
	if errRes != nil {
		return errRes
	}

	return buildOKResponse(category)
}

func (api *CategoryAPI) GetCategories(req *data.RestRequest) *data.RestResponse {
	accountId := getAccountCtx(req).Id
	categories, err := api.categoryManager.GetByAccount(accountId)

	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(categories)
}

func (api *CategoryAPI) CreateCategory(req *data.RestRequest) *data.RestResponse {
	createReq, err := getCategoryCreateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	if err := api.categoryManager.Create(getAccountCtx(req).Id, createReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) UpdateCategory(req *data.RestRequest) *data.RestResponse {
	category, errRes := api.getCategoryCtx(req)
	if errRes != nil {
		return errRes
	}

	updateReq, err := getCategoryUpdateBody(req)

	if err != nil {
		return buildServerError(err)
	}

	if err := api.categoryManager.Update(&category, updateReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) DeleteCategory(req *data.RestRequest) *data.RestResponse {
	category, errRes := api.getCategoryCtx(req)
	if errRes != nil {
		return errRes
	}

	if err := api.categoryManager.Delete(category.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}
