package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type CategoryManager struct {
	store store.CategoryStore
	clock clock.Clock
	uid   uid.UIDProvider
}

func (manager *CategoryManager) GetByRequest(req *rest.Request, res *rest.Response) {
	category := manager.getCategoryByAccount(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	res.Ok(category)
}

func (manager *CategoryManager) GetAllByRequest(req *rest.Request, res *rest.Response) {
	categories, err := manager.store.GetByAccount(req.Account.Id)
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

	category := manager.getCategoryByAccount(res, req.ResourceId, req.Account.Id)
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

// TODO: Check that category is not being used
// TODO: Submit category.deleted message
func (manager *CategoryManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	manager.getCategoryByAccount(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}
}

func (manager *CategoryManager) Exists(id string, accountId string) (bool, error) {
	category, err := manager.store.Get(id)
	if err != nil {
		return false, err
	}

	if category == nil || category.AccountId != accountId {
		return false, nil
	}

	return true, nil
}

func (manager *CategoryManager) Get(id string) (*data.Category, error) {
	category, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (manager *CategoryManager) GetByAccount(accountId string) ([]data.Category, error) {
	return manager.store.GetByAccount(accountId)
}

func (manager *CategoryManager) getCategoryByAccount(res *rest.Response, id string, accountId string) *data.Category {
	category, err := manager.Get(id)

	if err != nil {
		res.ErrInternal(err)
		return nil
	}

	if category == nil || category.AccountId != accountId {
		res.ErrBudgetNotFound()
		return nil
	}

	return category
}

func (manager *CategoryManager) nameUsed(accountId string, name string) (bool, error) {
	categories, err := manager.store.GetByAccount(accountId)
	if err != nil {
		return false, err
	}

	for _, category := range categories {
		if category.Name == name {
			return true, nil
		}
	}

	return false, nil
}

func (manager *CategoryManager) validate(res *rest.Response, name string, accountId string) {
	ok, err := manager.nameUsed(accountId, name)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if ok {
		res.ErrCategoryNameAlreadyUsed()
		return
	}
}
