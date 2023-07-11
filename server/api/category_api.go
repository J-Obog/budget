package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type CategoryAPI struct {
	categoryManager *manager.CategoryManager
}

func (api *CategoryAPI) GetCategory(r *rest.Request) *rest.Response {
	category, err := api.categoryManager.Get(r.Params.CategoryId(), r.Account.Id)
	if err != nil {
		return buildServerError(err)
	}
	if category == nil {
		return buildBadRequestError()
	}

	return buildOKResponse(category)
}

func (api *CategoryAPI) GetCategories(r *rest.Request) *rest.Response {
	categories, err := api.categoryManager.GetByAccount(r.Account.Id)

	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(categories)
}

func (api *CategoryAPI) CreateCategory(r *rest.Request) *rest.Response {
	reqBody, err := r.Body.CategoryCreateBody()
	if err != nil {
		return buildBadRequestError()
	}

	if err := api.validateCreate(reqBody, r.Account.Id); err != nil {
		return buildBadRequestError()
	}

	if err := api.categoryManager.Create(r.Account.Id, reqBody); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) UpdateCategory(r *rest.Request) *rest.Response {
	category, err := api.categoryManager.Get(r.Params.CategoryId(), r.Account.Id)
	if err != nil {
		return buildServerError(err)
	}
	if category == nil {
		return buildBadRequestError()
	}

	reqBody, err := r.Body.CategoryUpdateBody()
	if err != nil {
		return buildBadRequestError()
	}

	if err := api.validateUpdate(reqBody, r.Account.Id); err != nil {
		return buildBadRequestError()
	}

	if err := api.categoryManager.Update(category, reqBody); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) DeleteCategory(r *rest.Request) *rest.Response {
	category, err := api.categoryManager.Get(r.Params.CategoryId(), r.Account.Id)
	if err != nil {
		return buildServerError(err)
	}
	if category == nil {
		return buildBadRequestError()
	}

	if err := api.categoryManager.Delete(category.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) validateCreate(reqBody rest.CategoryCreateBody, accountId string) error {
	ok, err := api.categoryManager.NameExists(accountId, reqBody.Name)
	if err != nil {
		return err
	}
	if ok {
		return nil //RETURN SOME BAD REQ ERROR
	}

	return nil
}

func (api *CategoryAPI) validateUpdate(reqBody rest.CategoryUpdateBody, accountId string) error {
	return api.validateCreate(reqBody.CategoryCreateBody, accountId)
}
