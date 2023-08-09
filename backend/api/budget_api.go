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

	budgetm, err := api.getMaterializedBudget(*budget)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(budgetm)
}

// TODO: Check if
func (api *BudgetAPI) Filter(req *rest.Request) *rest.Response {
	return nil
}

func (api *BudgetAPI) Create(req *rest.Request) *rest.Response {
	accountId := testAccountId

	body, err := rest.ParseBody[rest.BudgetCreateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateCreate(body, accountId); err != nil {
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

	existing, err := api.budgetManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(existing, body); err != nil {
		return rest.Err(err)
	}

	if err := api.budgetManager.Update(existing, body); err != nil {
		return rest.Err(err)
	}

	return rest.Ok(existing)
}

func (api *BudgetAPI) Delete(req *rest.Request) *rest.Response {
	accountId := testAccountId
	id := req.Params.GetBudgetId()

	if err := api.budgetManager.Delete(id, accountId); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (api *BudgetAPI) validateCreate(body rest.BudgetCreateBody, accountId string) error {
	d := data.NewDate(body.Month, 1, body.Year)

	if err := d.IsValid(); err != nil {
		return err
	}

	if _, err := api.categoryManager.Get(body.CategoryId, accountId); err != nil {
		return err
	}

	if err := api.budgetManager.CheckCategoryNotInPeriod(body.CategoryId, accountId, body.Month, body.Year); err != nil {
		return err
	}

	return nil
}

func (api *BudgetAPI) validateUpdate(existing *data.Budget, body rest.BudgetUpdateBody) error {
	if body.CategoryId != existing.CategoryId {
		if _, err := api.categoryManager.Get(body.CategoryId, existing.AccountId); err != nil {
			return err
		}

		if err := api.budgetManager.CheckCategoryNotInPeriod(body.CategoryId, existing.AccountId, existing.Month, existing.Year); err != nil {
			return err
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
