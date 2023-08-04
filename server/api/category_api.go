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

func NewCategoryAPI(categoryManager *manager.CategoryManager, budgetManager *manager.BudgetManager) *CategoryAPI {
	return &CategoryAPI{
		budgetManager:   budgetManager,
		categoryManager: categoryManager,
	}
}

func getCategoryId(req *rest.Request) string {
	return ""
}

func (api *CategoryAPI) Get(req *rest.Request) *rest.Response {
	id := getCategoryId(req)

	category, err := api.categoryManager.Get(id, req.Account.Id)
	if err != nil {
		return rest.Err(err)
	}

	if category == nil {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

	return rest.Ok(category)
}

func (api *CategoryAPI) GetAll(req *rest.Request) *rest.Response {
	categories, err := api.categoryManager.GetAll(req.Account.Id)

	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(categories)
}

func (api *CategoryAPI) Create(req *rest.Request) *rest.Response {
	accountId := req.Account.Id

	body, err := rest.ParseBody[rest.CategoryCreateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateCreate(body, accountId); err != nil {
		return rest.Err(err)
	}

	category, err := api.categoryManager.Create(accountId, body)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(category)
}

func (api *CategoryAPI) Update(req *rest.Request) *rest.Response {
	id := getCategoryId(req)
	accountId := req.Account.Id

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

	_, err := api.categoryManager.Update(existing, body)
	if err != nil {
		return rest.Err(err)
	}

	/*if !ok {
		return rest.Err(rest.ErrInvalidCategoryId)
	}*/

	return rest.Ok(existing)
}

func (api *CategoryAPI) Delete(req *rest.Request) *rest.Response {
	categoryId := req.ResourceId
	accountId := req.Account.Id

	if err := api.checkCategoryBeingUsed(categoryId, accountId); err != nil {
		return rest.Err(err)
	}

	_, err := api.categoryManager.Delete(categoryId, accountId)
	if err != nil {
		return rest.Err(err)
	}

	/*if !ok {
		return rest.Err(rest.ErrInvalidCategoryId)
	}*/

	return rest.Success()
}

func (api *CategoryAPI) validateUpdate(existing *data.Category, body rest.CategoryUpdateBody) error {
	if existing == nil {
		return rest.ErrInvalidCategoryId
	}

	if body.Name != existing.Name {
		if err := api.checkName(body.Name); err != nil {
			return err
		}
		if err := api.checkNameAlreadyExists(existing.AccountId, body.Name); err != nil {
			return err
		}
	}

	return nil
}

func (api *CategoryAPI) validateCreate(body rest.CategoryCreateBody, accountId string) error {
	if err := api.checkName(body.Name); err != nil {
		return err
	}

	if err := api.checkNameAlreadyExists(accountId, body.Name); err != nil {
		return err
	}

	return nil
}

// CheckIfNameIsTaken
func (api *CategoryAPI) checkCategoryBeingUsed(categoryId string, accountId string) error {
	budgets, err := api.budgetManager.GetByCategory(accountId, categoryId)
	if err != nil {
		return err
	}

	if len(budgets) != 0 {
		return rest.ErrCategoryCurrentlyInUse
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

func (api *CategoryAPI) checkNameAlreadyExists(accountId string, name string) error {
	category, err := api.categoryManager.GetByName(accountId, name)
	if err != nil {
		return err
	}

	if category != nil {
		return rest.ErrCategoryNameAlreadyExists
	}
	return nil
}
