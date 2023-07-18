package manager

import (
	"github.com/J-Obog/paidoff/clock"
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
	category, err := manager.store.Get(req.ResourceId, req.Account.Get().Id)
	if err != nil {
		return rest.Err(err)
	}

	if category.Empty() {
		return rest.Err(rest.ErrInvalidCategoryId)
	}

	return rest.Ok(category)
}

func (manager *CategoryManager) GetAllByRequest(req *rest.Request) *rest.Response {
	categories, err := manager.store.GetAll(req.Account.Get().Id)

	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(categories)
}

func (manager *CategoryManager) CreateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.CategorySetBody)
	accountId := req.Account.Get().Id

	category := manager.getCategoryForCreate(accountId, body)

	if err := manager.validateSet(body, true); err != nil {
		return rest.Err(err)
	}

	if err := manager.store.Insert(category); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *CategoryManager) UpdateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.CategorySetBody)
	categoryId := req.ResourceId
	accountId := req.Account.Get().Id

	if err := manager.validateSet(body, true); err != nil {
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
	accountId := req.Account.Get().Id

	budgets, err := manager.budgetStore.GetByCategory(accountId, categoryId)
	if err != nil {
		return rest.Err(err)
	}

	if len(budgets) != 0 {
		return rest.Err(rest.ErrCategoryCurrentlyInUse)
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

func (manager *CategoryManager) validateSet(body rest.CategorySetBody, isUpdate bool) error {
	return nil
}

func (manager *CategoryManager) isNameUnique(accountId string, name string) (bool, error) {
	category, err := manager.store.GetByName(accountId, name)
	if err != nil {
		return false, err
	}

	if category.NotEmpty() {
		return false, nil
	}

	return true, nil
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

func (manager *CategoryManager) getCategoryForCreate(accountId string, body rest.CategorySetBody) data.Category {
	return data.Category{}
}

func getUpdateForCategoryUpdate(body rest.CategorySetBody) data.CategoryUpdate {
	return data.CategoryUpdate{}
}
