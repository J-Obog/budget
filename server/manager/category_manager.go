package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type CategoryManager struct {
	store       store.CategoryStore
	budgetStore store.BudgetStore
	clock       clock.Clock
	uid         uid.UIDProvider
	queue       queue.Queue
}

func (manager *CategoryManager) GetByRequest(req *rest.Request) *rest.Response {
	category, err := manager.store.Get(req.ResourceId, req.Account.Id)
	if err != nil {
		return rest.Err(err)
	}

	if category == nil {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

	return rest.Ok(category)
}

func (manager *CategoryManager) GetAllByRequest(req *rest.Request) *rest.Response {
	categories, err := manager.store.GetAll(req.Account.Id)

	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(categories)
}

func (manager *CategoryManager) CreateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.CategoryCreateBody)
	accountId := req.Account.Id

	if err := manager.validateCreate(body, accountId); err != nil {
		return rest.Err(err)
	}

	category := manager.getCategoryForCreate(accountId, body)

	if err := manager.store.Insert(category); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *CategoryManager) UpdateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.CategoryUpdateBody)
	categoryId := req.ResourceId
	accountId := req.Account.Id

	if err := manager.validateUpdate(body, accountId, categoryId); err != nil {
		return rest.Err(err)
	}

	timestamp := manager.clock.Now()
	update := getUpdateForCategoryUpdate(body)

	ok, err := manager.store.Update(categoryId, accountId, update, timestamp)

	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

	return rest.Success()
}

func (manager *CategoryManager) DeleteByRequest(req *rest.Request) *rest.Response {
	categoryId := req.ResourceId
	accountId := req.Account.Id

	if err := manager.checkCategoryBeingUsed(categoryId, accountId); err != nil {
		return rest.Err(err)
	}

	if err := manager.sendMsg(req.ResourceId); err != nil {
		return rest.Err(err)
	}

	ok, err := manager.store.Delete(categoryId, accountId)

	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

	return rest.Success()
}

func (manager *CategoryManager) validateUpdate(body rest.CategoryUpdateBody, accountId string, id string) error {
	existing, err := manager.store.Get(id, accountId)
	if err != nil {
		return err
	}

	if existing == nil {
		return rest.ErrInvalidCategoryId
	}

	if body.Name != existing.Name {
		if err := manager.checkName(body.Name); err != nil {
			return err
		}
		if err := manager.checkNameAlreadyExists(accountId, body.Name); err != nil {
			return err
		}
	}

	return nil
}

func (manager *CategoryManager) validateCreate(body rest.CategoryCreateBody, accountId string) error {
	if err := manager.checkName(body.Name); err != nil {
		return err
	}

	if err := manager.checkNameAlreadyExists(accountId, body.Name); err != nil {
		return err
	}

	return nil
}

func (manager *CategoryManager) checkCategoryBeingUsed(categoryId string, accountId string) error {
	budgets, err := manager.budgetStore.GetByCategory(accountId, categoryId)
	if err != nil {
		return err
	}

	if len(budgets) != 0 {
		return rest.ErrCategoryCurrentlyInUse
	}

	return nil
}

func (manager *CategoryManager) checkName(name string) error {
	nameLen := len(name)
	if nameLen < config.LimitMinCategoryNameChars || nameLen > config.LimitMaxCategoryNameChars {
		return rest.ErrInvalidCategoryName
	}

	return nil
}

func (manager *CategoryManager) checkNameAlreadyExists(accountId string, name string) error {
	category, err := manager.store.GetByName(accountId, name)
	if err != nil {
		return err
	}

	if category != nil {
		return rest.ErrCategoryNameAlreadyExists
	}
	return nil
}

// TODO: better msg id?
func (manager *CategoryManager) sendMsg(id string) error {
	msgId := manager.uid.GetId()

	msg := queue.Message{
		Id: msgId,
		Data: queue.CategoryDeletedMessage{
			CategoryId: id,
		},
	}

	if err := manager.queue.Push(msg, queue.QueueName_CategoryDeleted); err != nil {
		return err
	}

	return nil
}

func (manager *CategoryManager) getCategoryForCreate(accountId string, body rest.CategoryCreateBody) data.Category {
	id := manager.uid.GetId()
	timestamp := manager.clock.Now()

	return data.Category{
		Id:        id,
		AccountId: accountId,
		Name:      body.Name,
		Color:     body.Color,
		UpdatedAt: timestamp,
		CreatedAt: timestamp,
	}
}

func getUpdateForCategoryUpdate(body rest.CategoryUpdateBody) data.CategoryUpdate {
	return data.CategoryUpdate{
		Name:  body.Name,
		Color: body.Color,
	}
}
