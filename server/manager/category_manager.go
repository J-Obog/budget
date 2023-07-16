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

type categoryValidateCommon struct {
	name      string
	accountId string
}

func (manager *CategoryManager) GetByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	id := req.Params.CategoryId()

	category := manager.getCategoryByAccount(res, id, accountId)
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

	validateCommon := categoryValidateCommon{name: body.Name}
	manager.validate(res, validateCommon)
	if res.IsErr() {
		return
	}

	now := manager.clock.Now()
	id := manager.uid.GetId()
	category := newCategory(body, id, req.Account.Id, now)

	if err := manager.store.Insert(category); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *CategoryManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	id := req.Params.CategoryId()

	category := manager.getCategoryByAccount(res, id, accountId)
	if res.IsErr() {
		return
	}

	body := req.Body.(rest.CategoryUpdateBody)

	validateCommon := categoryValidateCommon{name: body.Name}
	manager.validate(res, validateCommon)
	if res.IsErr() {
		return
	}

	now := manager.clock.Now()
	updateCategory(body, category, now)

	if err := manager.store.Update(*category); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

// TODO: Check that category is not being used
// TODO: Submit category.deleted message
func (manager *CategoryManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	id := req.Params.CategoryId()

	manager.getCategoryByAccount(res, id, accountId)
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

func (manager *CategoryManager) validate(res *rest.Response, validateCom categoryValidateCommon) {
	ok, err := manager.nameUsed(validateCom.accountId, validateCom.name)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if ok {
		res.ErrCategoryNameAlreadyUsed()
		return
	}
}
