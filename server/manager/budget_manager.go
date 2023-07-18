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

func (manager *BudgetManager) GetByRequest(req *rest.Request) *rest.Response {
	accountId := req.Account.Get().Id
	budgetId := req.ResourceId

	budget, err := manager.store.Get(budgetId, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if budget.Empty() {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Ok(budget)
}

func (manager *BudgetManager) GetAllByRequest(req *rest.Request) *rest.Response {
	accountId := req.Account.Get().Id

	query := req.Query.(rest.BudgetQuery)
	filter := getFilterForBudgetQuery(query)

	budgets, err := manager.store.GetBy(accountId, filter)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(budgets)
}

func (manager *BudgetManager) CreateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.BudgetSetBody)
	accountId := req.Account.Get().Id

	/*bodyDate := data.NewDate(body.Month, 1, body.Year)

	periodStart := manager.clock.FromDate(bodyDate)
	periodEnd := manager.clock.MonthEnd(periodStart)
	*/
	budget := manager.getBudgetForCreate(accountId, body)

	if err := manager.store.Insert(budget); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *BudgetManager) UpdateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.BudgetSetBody)
	accountId := req.Account.Get().Id
	budgetId := req.ResourceId

	timestamp := manager.clock.Now()
	update := getUpdateForBudgetUpdate(body)

	ok, err := manager.store.Update(budgetId, accountId, update, timestamp)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Success()
}

func (manager *BudgetManager) DeleteByRequest(req *rest.Request) *rest.Response {
	ok, err := manager.store.Delete(req.ResourceId, req.Account.Get().Id)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Success()
}

func (manager *BudgetManager) validateSet(body rest.BudgetSetBody, accountId string, isUpdate bool) error {
	/*if ok := manager.clock.IsDateValid(bodyDate); !ok {
		return res.ErrInvalidDate()
	}*/

	category, err := manager.categoryStore.Get(body.CategoryId, accountId)
	if err != nil {
		return err
	}

	if category.Empty() {
		return rest.ErrInvalidCategoryId
	}

	budgets, err := manager.store.GetByPeriodCategory(accountId, body.CategoryId, periodStart, periodEnd)
	if err != nil {
		return err
	}

	if budgets.NotEmpty() {
		return rest.ErrCategoryAlreadyInBudgetPeriod
	}

	return nil
}

func (manager *BudgetManager) getBudgetForCreate(accountId string, body rest.BudgetSetBody) data.Budget {
	return data.Budget{}
}

func getFilterForBudgetQuery(q rest.BudgetQuery) data.BudgetFilter {
	return data.BudgetFilter{}
}

func getUpdateForBudgetUpdate(body rest.BudgetSetBody) data.BudgetUpdate {
	return data.BudgetUpdate{}
}
