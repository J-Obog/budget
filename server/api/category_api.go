package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type CategoryAPI struct {
	categoryManager *manager.CategoryManager
}

func (api *CategoryAPI) GetCategory(req *rest.Request) *rest.Response {
	return api.categoryManager.GetByRequest(req)
}

func (api *CategoryAPI) GetCategories(req *rest.Request) *rest.Response {
	return api.categoryManager.GetAllByRequest(req)
}

func (api *CategoryAPI) CreateCategory(req *rest.Request) *rest.Response {
	return api.categoryManager.CreateByRequest(req)
}

func (api *CategoryAPI) UpdateCategory(req *rest.Request) *rest.Response {
	return api.categoryManager.UpdateByRequest(req)
}

func (api *CategoryAPI) DeleteCategory(req *rest.Request) *rest.Response {
	return api.categoryManager.DeleteByRequest(req)
}
