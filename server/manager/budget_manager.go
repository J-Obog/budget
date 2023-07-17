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
	budget, err := manager.store.Get(req.ResourceId, req.Account.Id)

	if err != nil {
		res.ErrInternal(err)
		return
	}

	if budget == nil {
		res.ErrBudgetNotFound()
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
	bodyDate := data.NewDate(body.Month, 1, body.Year)
	period := manager.clock.FromDate(bodyDate)
	newBudget := newBudget(body, id, req.Account.Id, timestamp, period)

	if ok := manager.clock.IsDateValid(bodyDate); !ok {
		res.ErrInvalidDate()
		return
	}

	if manager.validate(res, body.CategoryId, req.Account.Id); res.IsErr() {
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
	update := data.BudgetUpdate{
		Projected: &body.Projected,
		UpdatedAt: &timestamp,
	}

	if manager.validate(res, body.Month, body.Year, body.CategoryId, req.Account.Id); res.IsErr() {
		return
	}

	ok, err := manager.store.Update(req.ResourceId, update)

	if err != nil {
		res.ErrInternal(err)
		return
	}

	if !ok {
		res.ErrBudgetNotFound()
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	ok, err := manager.store.Delete(req.ResourceId, req.Account.Id)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if !ok {
		res.ErrBudgetNotFound()
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) validate(res *rest.Response, accountId string, categoryId string) {
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
