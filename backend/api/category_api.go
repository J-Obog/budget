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

func (api *CategoryAPI) Get(req *rest.Request) *rest.Response {
	id := req.Params.GetCategoryId()
	accountId := testAccountId

	category, err := api.categoryManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if category == nil {
		return rest.Err(rest.ErrInvalidCategoryId)
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
	var body rest.CategoryCreateBody
	accountId := testAccountId

	if err := req.Body.To(&body); err != nil {
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
	var body rest.CategoryUpdateBody
	id := req.Params.GetCategoryId()
	accountId := testAccountId

	if err := req.Body.To(&body); err != nil {
		return rest.Err(err)
	}

	existing, err := api.categoryManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(existing, body); err != nil {
		return rest.Err(err)
	}

	ok, err := api.categoryManager.Update(existing, body)

	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

	return rest.Ok(existing)
}

func (api *CategoryAPI) Delete(req *rest.Request) *rest.Response {
	id := req.Params.GetCategoryId()
	accountId := testAccountId

	if err := api.validateDelete(id, accountId); err != nil {
		return rest.Err(err)
	}

	ok, err := api.categoryManager.Delete(id, accountId)

	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

	return rest.Success()
}

func (api *CategoryAPI) validateDelete(id string, accountId string) error {
	ok, err := api.budgetManager.CategoryIsNotUsed(id, accountId)

	if err != nil {
		return err
	}

	if !ok {
		return rest.ErrCategoryCurrentlyInUse
	}

	return nil
}

func (api *CategoryAPI) validateUpdate(existing *data.Category, body rest.CategoryUpdateBody) error {
	if existing == nil {
		return rest.ErrInvalidCategoryId
	}

	if body.Name != existing.Name {
		if err := api.checkName(body.Name); err != nil {
			return err
		}

		ok, err := api.categoryManager.NameIsUnique(existing.AccountId, body.Name)

		if err != nil {
			return err
		}

		if !ok {
			return rest.ErrCategoryNameAlreadyExists
		}
	}

	return nil
}

func (api *CategoryAPI) validateCreate(accountId string, body rest.CategoryCreateBody) error {
	if err := api.checkName(body.Name); err != nil {
		return err
	}

	ok, err := api.categoryManager.NameIsUnique(accountId, body.Name)

	if err != nil {
		return err
	}

	if !ok {
		return rest.ErrCategoryNameAlreadyExists
	}

	return nil
}

func (api *CategoryAPI) checkName(name string) error {
	nameLen := len(name)
	if nameLen < config.LimitMinCategoryNameChars || nameLen > config.LimitMaxCategoryNameChars {
		return rest.ErrInvalidCategoryName
	}

	return nil
}
