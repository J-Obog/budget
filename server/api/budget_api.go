package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/validation"
)

type BudgetAPI struct {
	budgetManager   *manager.BudgetManager
	categoryManager *manager.CategoryManager
}

func (api *BudgetAPI) getBudgetCtx(req *data.RestRequest) (data.Budget, *data.RestResponse) {
	budget, err := api.budgetManager.Get(req.UrlParams["budgetId"].(string))

	if err != nil {
		return data.Budget{}, buildServerError(err)
	}
	if budget == nil || budget.AccountId != getAccountCtx(req).Id {
		return data.Budget{}, buildBadRequestError()
	}

	return *budget, nil
}

func (api *BudgetAPI) GetBudget(req *data.RestRequest) *data.RestResponse {
	budget, errResp := api.getBudgetCtx(req)
	if errResp != nil {
		return errResp
	}

	return buildOKResponse(budget)
}

func (api *BudgetAPI) GetBudgets(req *data.RestRequest) *data.RestResponse {
	q, err := getBugetGetQuery(req)

	if err != nil {
		return buildServerError(err)
	}

	accountId := getAccountCtx(req).Id

	budgets, err := api.budgetManager.GetByAccount(accountId, q)
	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(budgets)
}

func (api *BudgetAPI) CreateBudget(req *data.RestRequest) *data.RestResponse {
	if err := validation.ValidateBudgetCreateReq(req.Body); err != nil {
		return buildBadRequestError()
	}

	createReq, err := getBugetCreateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	accountId := getAccountCtx(req).Id

	if errResp := checkCategoryExists(&createReq.CategoryId, accountId, api.categoryManager); errResp != nil {
		return errResp
	}

	if err := api.budgetManager.Create(accountId, createReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) UpdateBudget(req *data.RestRequest) *data.RestResponse {
	if err := validation.ValidateBudgetUpdateReq(req.Body); err != nil {
		return buildBadRequestError()
	}

	budget, errResp := api.getBudgetCtx(req)
	if errResp != nil {
		return errResp
	}

	updateReq, err := getBugetUpdateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	if errResp = checkCategoryExists(&updateReq.CategoryId, budget.AccountId, api.categoryManager); errResp != nil {
		return errResp
	}

	if err := api.budgetManager.Update(&budget, updateReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) DeleteBudget(req *data.RestRequest) *data.RestResponse {
	budget, errResp := api.getBudgetCtx(req)
	if errResp != nil {
		return errResp
	}

	if err := api.budgetManager.Delete(budget.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}
