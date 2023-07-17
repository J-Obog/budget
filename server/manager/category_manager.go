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

func (manager *CategoryManager) GetByRequest(req *rest.Request, res *rest.Response) {
	category := manager.getCategory(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	res.Ok(category)
}

func (manager *CategoryManager) GetAllByRequest(req *rest.Request, res *rest.Response) {
	filter := data.CategoryFilter{AccountId: &req.Account.Id}
	categories, err := manager.store.GetBy(filter)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(categories)
}

func (manager *CategoryManager) CreateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.CategoryCreateBody)
	timestamp := manager.clock.Now()
	id := manager.uid.GetId()
	category := newCategory(body, id, req.Account.Id, timestamp)

	if manager.validate(res, body.Name, req.Account.Id); res.IsErr() {
		return
	}

	if err := manager.store.Insert(category); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *CategoryManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.CategoryUpdateBody)
	timestamp := manager.clock.Now()

	category := manager.getCategory(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	if manager.validate(res, body.Name, req.Account.Id); res.IsErr() {
		return
	}

	updateCategory(body, category, timestamp)

	if err := manager.store.Update(*category); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *CategoryManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	filter := data.BudgetFilter{CategoryId: &req.ResourceId, AccountId: &req.Account.Id}
	budgets, err := manager.budgetStore.GetBy(filter)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if len(budgets) != 0 {
		res.ErrCategoryIsCurrentlyUsed()
		return
	}

	if err := manager.sendMsg(req.ResourceId); err != nil {
		res.ErrInternal(err)
		return
	}

	ok, err := manager.store.Delete(req.ResourceId, req.Account.Id)

	if err != nil {
		res.ErrInternal(err)
		return
	}

	if !ok {
		res.ErrCategoryNotFound()
		return
	}

	res.Ok(nil)
}

func (manager *CategoryManager) getCategory(res *rest.Response, id string, accountId string) *data.Category {
	category, err := manager.store.Get(id, accountId)

	if err != nil {
		res.ErrInternal(err)
		return nil
	}

	if category == nil {
		res.ErrBudgetNotFound()
		return nil
	}

	return category
}

func (manager *CategoryManager) validate(res *rest.Response, name string, accountId string) {
	filter := data.CategoryFilter{AccountId: &accountId, Name: &name}
	category, err := manager.store.GetBy(filter)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if category != nil {
		res.ErrCategoryNameAlreadyUsed()
		return
	}
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
