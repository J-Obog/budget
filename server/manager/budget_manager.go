package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type BudgetManager struct {
	store           store.BudgetStore
	categoryManager *CategoryManager
	clock           clock.Clock
	uid             uid.UIDProvider
}

type budgetValidateCommon struct {
	month      int
	year       int
	accountId  string
	categoryId string
}

func (manager *BudgetManager) Get(id string) (*data.Budget, error) {
	budget, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (manager *BudgetManager) GetByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	budgetId := req.Params.BudgetId()

	budget := manager.getBudgetByAccount(res, budgetId, accountId)

	if res.IsErr() {
		return
	}

	res.Ok(budget)
}

func (manager *BudgetManager) GetAllByRequest(req *rest.Request, res *rest.Response) {
	query := req.Query.BudgetQuery()
	accountId := req.Account.Id

	budgets, err := manager.store.GetByAccount(accountId)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	filtered := filter[data.Budget](budgets, func(b *data.Budget) bool {
		if query.Month != nil && b.Month != *query.Month {
			return false
		}
		if query.Year != nil && b.Year != *query.Year {
			return false
		}

		return true
	})

	res.Ok(filtered)
}

func (manager *BudgetManager) CreateByRequest(req *rest.Request, res *rest.Response) {
	body, err := req.Body.BudgetCreateBody()
	if err != nil {
		res.ErrBadRequest()
		return
	}

	validateCommon := budgetValidateCommon{month: body.Month, year: body.Year, categoryId: body.CategoryId, accountId: req.Account.Id}

	manager.validate(res, validateCommon)
	if res.IsErr() {
		return
	}

	now := manager.clock.Now()
	id := manager.uid.GetId()

	newBudget := newBudget(body, id, req.Account.Id, now)

	if err := manager.store.Insert(newBudget); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	budgetId := req.Params.BudgetId()

	budget := manager.getBudgetByAccount(res, budgetId, accountId)
	if res.IsErr() {
		return
	}

	body, err := req.Body.BudgetUpdateBody()
	if err != nil {
		res.ErrBadRequest()
		return
	}

	validateCommon := budgetValidateCommon{month: body.Month, year: body.Year, categoryId: body.CategoryId, accountId: req.Account.Id}

	manager.validate(res, validateCommon)
	if res.IsErr() {
		return
	}

	now := manager.clock.Now()
	updateBudget(body, budget, now)

	if err = manager.store.Update(*budget); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	budgetId := req.Params.BudgetId()

	manager.getBudgetByAccount(res, budgetId, accountId)
	if res.IsErr() {
		return
	}

	if err := manager.store.Delete(budgetId); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *BudgetManager) categoryInPeriod(id string, accountId string, month int, year int) (bool, error) {
	budgets, err := manager.store.GetByAccount(accountId)
	if err != nil {
		return false, err
	}

	filtered := filter[data.Budget](budgets, func(b *data.Budget) bool {
		return b.Month == month && b.Year == year
	})

	for _, budget := range filtered {
		if budget.CategoryId == id {
			return true, nil
		}
	}

	return false, nil
}

func (manager *BudgetManager) validate(res *rest.Response, validateCom budgetValidateCommon) {
	//check if date is valid

	categoryId := validateCom.categoryId
	accountId := validateCom.accountId
	month := validateCom.month
	year := validateCom.year

	ok, err := manager.categoryManager.Exists(categoryId, accountId)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if !ok {
		res.ErrCategoryNotFound()
		return
	}

	ok, err = manager.categoryInPeriod(categoryId, accountId, month, year)

	if err != nil {
		res.ErrInternal(err)
		return
	}

	if ok {
		res.ErrCategoryInBudgetPeriod()
		return
	}
}

func (manager *BudgetManager) getBudgetByAccount(res *rest.Response, id string, accountId string) *data.Budget {
	budget, err := manager.Get(id)

	if err != nil {
		res.ErrInternal(err)
		return nil
	}

	if budget == nil || budget.AccountId != accountId {
		res.ErrBudgetNotFound()
		return nil
	}

	return budget
}
