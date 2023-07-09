package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type BudgetAPI struct {
	budgetManager   *manager.BudgetManager
	categoryManager *manager.CategoryManager
}

func (api *BudgetAPI) GetBudget(r *rest.Request) *rest.Response {
	budget, errResp := api.bugetCtx(r)
	if errResp != nil {
		return errResp
	}

	return buildOKResponse(budget)
}

func (api *BudgetAPI) GetBudgets(r *rest.Request) *rest.Response {
	if errResp := api.validateGetsRequest(r); errResp != nil {
		return errResp
	}

	budgets, err := api.budgetManager.GetByAccount(r.Account.Id, r.Query.BudgetQuery())
	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(budgets)
}

func (api *BudgetAPI) CreateBudget(r *rest.Request) *rest.Response {
	if errResp := api.validateCreateRequest(r); errResp != nil {
		return errResp
	}

	if err := api.budgetManager.Create(r.Account.Id, r.Body.BudgetCreateBody()); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) UpdateBudget(r *rest.Request) *rest.Response {
	budget, errResp := api.bugetCtx(r)
	if errResp != nil {
		return errResp
	}

	if errResp = api.validateUpdateRequest(r); errResp != nil {
		return errResp
	}

	if err := api.budgetManager.Update(&budget, r.Body.BudgetUpdateBody()); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) DeleteBudget(r *rest.Request) *rest.Response {
	budget, errResp := api.bugetCtx(r)
	if errResp != nil {
		return errResp
	}

	if err := api.budgetManager.Delete(budget.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) validateUpdateRequest(req *rest.Request) *rest.Response {
	return nil
}

func (api *BudgetAPI) validateCreateRequest(req *rest.Request) *rest.Response {
	return nil
}

func (api *BudgetAPI) validateGetsRequest(req *rest.Request) *rest.Response {
	return nil
}

func (api *BudgetAPI) bugetCtx(r *rest.Request) (data.Budget, *rest.Response) {
	budget, err := api.budgetManager.Get(r.Params.BudgetId())

	if err != nil {
		return data.Budget{}, buildServerError(err)
	}
	if budget == nil || budget.AccountId != r.Account.Id {
		return data.Budget{}, buildBadRequestError()
	}

	return *budget, nil
}
