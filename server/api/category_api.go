package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type CategoryAPI struct {
	categoryManager *manager.CategoryManager
}

func (api *CategoryAPI) GetCategory(req *rest.Request, res *rest.Response) {
	api.categoryManager.GetByRequest(req, res)
}

func (api *CategoryAPI) GetCategories(req *rest.Request, res *rest.Response) {
	api.categoryManager.GetAllByRequest(req, res)
}

func (api *CategoryAPI) CreateCategory(req *rest.Request, res *rest.Response) {
	api.categoryManager.CreateByRequest(req, res)
}

func (api *CategoryAPI) UpdateCategory(req *rest.Request, res *rest.Response) {
	api.categoryManager.UpdateByRequest(req, res)
}

func (api *CategoryAPI) DeleteCategory(req *rest.Request, res *rest.Response) {
	api.categoryManager.DeleteByRequest(req, res)
}
