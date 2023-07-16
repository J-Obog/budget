package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type BudgetManager struct {
	store         store.BudgetStore
	categoryStore store.CategoryStore
	clock         clock.Clock
	uid           uid.UIDProvider
}

func (manager *BudgetManager) GetByRequest(req *rest.Request, res *rest.Response) {
	budget := manager.getBudget(res, req.ResourceId, req.Account.Id)

	if res.IsErr() {
		return
	}

	res.Ok(budget)
}

func (manager *BudgetManager) GetAllByRequest(req *rest.Request, res *rest.Response) {
	query := req.Query.(rest.BudgetQuery)

	filter := data.BudgetFilter{
		AccountId: &req.Account.Id,
		Month:     query.Month,
		Year:      query.Year,
	}

	budgets, err := manager.store.GetBy(filter)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(budgets)
}

func (manager *BudgetManager) CreateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.BudgetCreateBody)
	timestamp := manager.clock.Now()
	id := manager.uid.GetId()
	newBudget := newBudget(body, id, req.Account.Id, timestamp)

	if manager.validate(res, body.Month, body.Year, body.CategoryId, req.Account.Id); res.IsErr() {
		return
	}

	if err := manager.store.Insert(newBudget); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	body := req.Body.(rest.BudgetUpdateBody)
	timestamp := manager.clock.Now()

	budget := manager.getBudget(res, req.ResourceId, req.Account.Id)
	if res.IsErr() {
		return
	}

	updateBudget(body, budget, timestamp)

	if manager.validate(res, body.Month, body.Year, body.CategoryId, req.Account.Id); res.IsErr() {
		return
	}

	if err := manager.store.Update(*budget); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	if manager.getBudget(res, req.ResourceId, req.Account.Id); res.IsErr() {
		return
	}

	if err := manager.store.Delete(req.ResourceId); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) validate(res *rest.Response, month int, year int, accountId string, categoryId string) {
	if ok := isDateValid(month, 1, year); !ok {
		res.ErrInvalidDate()
		return
	}

	category, err := manager.categoryStore.Get(categoryId, accountId)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if category == nil {
		res.ErrCategoryNotFound()
		return
	}

	filter := data.BudgetFilter{
		CategoryId: &categoryId,
		AccountId:  &accountId,
		Month:      &month,
		Year:       &year,
	}

	budgets, err := manager.store.GetBy(filter)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if budgets.First() != nil {
		res.ErrCategoryInBudgetPeriod()
		return
	}
}

func (manager *BudgetManager) getBudget(res *rest.Response, id string, accountId string) *data.Budget {
	budget, err := manager.store.Get(id, accountId)

	if err != nil {
		res.ErrInternal(err)
		return nil
	}

	if budget == nil {
		res.ErrBudgetNotFound()
		return nil
	}

	return budget
}
