package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type BudgetAPI struct {
	budgetManager      *manager.BudgetManager
	transactionManager *manager.TransactionManager
	categoryManager    *manager.CategoryManager
}

func NewBudgetAPI(
	budgetManager *manager.BudgetManager,
	transactionManager *manager.TransactionManager,
	categoryManager *manager.CategoryManager,
) *BudgetAPI {
	return &BudgetAPI{
		budgetManager:      budgetManager,
		transactionManager: transactionManager,
		categoryManager:    categoryManager,
	}
}

func (api *BudgetAPI) Get(req *rest.Request) *rest.Response {
	id := req.Params.GetBudgetId()
	accountId := testAccountId

	budget, err := api.budgetManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if budget == nil {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	budgetm, err := api.getMaterializedBudget(*budget)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(budgetm)
}

// TODO: Check if this shpuld be converted to 'GetForPeriod'
func (api *BudgetAPI) Filter(req *rest.Request) *rest.Response {
	return nil
}

func (api *BudgetAPI) Create(req *rest.Request) *rest.Response {
	accountId := testAccountId
	body, err := rest.ParseBody[rest.BudgetCreateBody](req.Body)

	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateCreate(accountId, body); err != nil {
		return rest.Err(err)
	}

	newBudget, err := api.budgetManager.Create(accountId, body)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(newBudget)
}

func (api *BudgetAPI) Update(req *rest.Request) *rest.Response {
	accountId := testAccountId
	id := req.Params.GetBudgetId()
	body, err := rest.ParseBody[rest.BudgetUpdateBody](req.Body)

	if err != nil {
		return rest.Err(err)
	}

	budget, err := api.budgetManager.Get(id, accountId)

	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(budget, body); err != nil {
		return rest.Err(err)
	}

	ok, err := api.budgetManager.Update(budget, body)

	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Ok(budget)
}

func (api *BudgetAPI) Delete(req *rest.Request) *rest.Response {
	accountId := testAccountId
	id := req.Params.GetBudgetId()
	ok, err := api.budgetManager.Delete(id, accountId)

	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Success()
}

func (api *BudgetAPI) validateCreate(accountId string, body rest.BudgetCreateBody) error {
	d := data.NewDate(body.Month, 1, body.Year)

	if err := d.IsValid(); err != nil {
		return err
	}

	ok, err := api.categoryManager.Exists(body.CategoryId, accountId)
	if err != nil {
		return err
	}

	if !ok {
		return rest.ErrInvalidCategoryId
	}

	ok, err = api.budgetManager.CategoryIsUniqueForPeriod(
		body.CategoryId,
		accountId,
		body.Month,
		body.Year,
	)

	if err != nil {
		return err
	}

	if !ok {
		return rest.ErrCategoryAlreadyInBudgetPeriod
	}

	return nil
}

func (api *BudgetAPI) validateUpdate(existing *data.Budget, body rest.BudgetUpdateBody) error {
	if existing == nil {
		return rest.ErrInvalidBudgetId
	}

	if body.CategoryId != existing.CategoryId {
		ok, err := api.categoryManager.Exists(body.CategoryId, existing.AccountId)
		if err != nil {
			return err
		}

		if !ok {
			return rest.ErrInvalidCategoryId
		}

		ok, err = api.budgetManager.CategoryIsUniqueForPeriod(
			body.CategoryId,
			existing.AccountId,
			existing.Month,
			existing.Year,
		)

		if err != nil {
			return err
		}

		if !ok {
			return rest.ErrCategoryAlreadyInBudgetPeriod
		}
	}

	return nil
}

func (api *BudgetAPI) getMaterializedBudget(budget data.Budget) (data.BudgetMaterialized, error) {
	budgetMaterialized := data.BudgetMaterialized{Budget: budget}
	total, err := api.transactionManager.GetTotalForPeriodCategory(
		budget.AccountId,
		budget.CategoryId,
		budget.Month,
		budget.Year,
	)

	budgetMaterialized.Actual = total
	return budgetMaterialized, err
}
