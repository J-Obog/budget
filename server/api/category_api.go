package api

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type CategoryAPI struct {
	categoryStore store.CategoryStore
	budgetStore   store.BudgetStore
	clock         clock.Clock
	uuidProvider  uuid.UuidProvider
}

func NewCategoryAPI(
	categoryStore store.CategoryStore,
	budgetStore store.BudgetStore,
	clock clock.Clock,
	uuidProvider uuid.UuidProvider,
) *CategoryAPI {
	return &CategoryAPI{
		categoryStore: categoryStore,
		budgetStore:   budgetStore,
		clock:         clock,
		uuidProvider:  uuidProvider,
	}
}

func getCategoryId(req *rest.Request) string {
	return ""
}

func (api *CategoryAPI) Get(req *rest.Request) *rest.Response {
	id := getCategoryId(req)
	accountId := testAccountId

	category, err := api.categoryStore.Get(id, accountId)
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

	categories, err := api.categoryStore.GetAll(accountId)

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

	if err := api.validateCreate(body, accountId); err != nil {
		return rest.Err(err)
	}

	timestamp := api.clock.Now()
	id := api.uuidProvider.GetUuid()

	newCategory := data.Category{
		Id:        id,
		AccountId: accountId,
		Name:      body.Name,
		Color:     body.Color,
		UpdatedAt: timestamp,
		CreatedAt: timestamp,
	}

	err = api.categoryStore.Insert(newCategory)
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

	existing, err := api.categoryStore.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(existing, body); err != nil {
		return rest.Err(err)
	}

	//timestamp :=
	ok, err := api.categoryStore.Update(existing, body, timestamp)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

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
