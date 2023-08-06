package api

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type CategoryAPI struct {
	categoryManager *manager.CategoryManager
	budgetManager   *manager.BudgetManager
}

func NewCategoryAPI(
	categoryManager *manager.CategoryManager,
	budgetManager *manager.BudgetManager,
) *CategoryAPI {
	return &CategoryAPI{
		categoryManager: categoryManager,
		budgetManager:   budgetManager,
	}
}

func getCategoryId(req *rest.Request) string {
	return ""
}

func (api *CategoryAPI) Get(req *rest.Request) *rest.Response {
	id := getCategoryId(req)
	accountId := testAccountId

	category, err := api.categoryManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(category)
}

func (api *CategoryAPI) GetAll(req *rest.Request) *rest.Response {
	accountId := testAccountId
	categories, err := api.categoryManager.GetAll(accountId)

	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(categories)
}

func (api *CategoryAPI) Create(req *rest.Request) *rest.Response {
	accountId := testAccountId

	body, err := rest.ParseBody[rest.CategoryCreateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateCreate(accountId, body); err != nil {
		return rest.Err(err)
	}

	newCategory, err := api.categoryManager.Create(accountId, body)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(newCategory)
}

func (api *CategoryAPI) Update(req *rest.Request) *rest.Response {
	id := getCategoryId(req)
	accountId := testAccountId

	body, err := rest.ParseBody[rest.CategoryUpdateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	existing, err := api.categoryManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(existing, body); err != nil {
		return rest.Err(err)
	}

	if err := api.categoryManager.Update(existing, body); err != nil {
		return rest.Err(err)
	}

	return rest.Ok(existing)
}

func (api *CategoryAPI) Delete(req *rest.Request) *rest.Response {
	id := getCategoryId(req)
	accountId := testAccountId

	if err := api.validateDelete(id, accountId); err != nil {
		return rest.Err(err)
	}

	if err := api.categoryManager.Delete(id, accountId); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (api *CategoryAPI) validateDelete(id string, accountId string) error {
	return api.budgetManager.CheckCategoryNotInUse(id, accountId)
}

func (api *CategoryAPI) validateUpdate(existing *data.Category, body rest.CategoryUpdateBody) error {
	if body.Name != existing.Name {
		if err := api.checkName(body.Name); err != nil {
			return err
		}
		if err := api.budgetManager.CheckCategoryNotInUse(existing.AccountId, body.Name); err != nil {
			return err
		}
	}

	return nil
}

func (api *CategoryAPI) validateCreate(accountId string, body rest.CategoryCreateBody) error {
	if err := api.checkName(body.Name); err != nil {
		return err
	}

	return api.categoryManager.CheckNameNotTaken(accountId, body.Name)
}

func (api *CategoryAPI) checkName(name string) error {
	nameLen := len(name)
	if nameLen < config.LimitMinCategoryNameChars || nameLen > config.LimitMaxCategoryNameChars {
		return rest.ErrInvalidCategoryName
	}

	return nil
}
