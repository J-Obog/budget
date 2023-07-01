package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
)

type BudgetAPI struct {
	budgetManager   *manager.BudgetManager
	categoryManager *manager.CategoryManager
}

func (api *BudgetAPI) GetBudget(req *data.RestRequest) *data.RestResponse {
	budget, errRes := checkBudget(req, api.budgetManager)
	if errRes != nil {
		return errRes
	}

	return buildOKResponse(budget)
}

func (api *BudgetAPI) GetBudgets(req *data.RestRequest) *data.RestResponse {
	q, err := getBugetGetQuery(req)

	if err != nil {
		return buildServerError(err)
	}

	accountId := getAccountCtx(req).Id

	budgets, err := api.budgetManager.GetByAccount(accountId)
	if err != nil {
		return buildServerError(err)
	}

	filter := NewFilter[data.Budget]()
	filter.AddCheck(filterBudget(q))
	filtered := filter.Filter(budgets)

	return buildOKResponse(filtered)
}

func (api *BudgetAPI) CreateBudget(req *data.RestRequest) *data.RestResponse {
	createReq, err := getBugetCreateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	errRes := validateCategoryId(createReq.CategoryId, req, api.categoryManager)
	if errRes != nil {
		return errRes
	}

	if err := api.budgetManager.Create(getAccountCtx(req).Id, createReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) UpdateBudget(req *data.RestRequest) *data.RestResponse {
	budget, errRes := checkBudget(req, api.budgetManager)
	if errRes != nil {
		return errRes
	}

	updateReq, err := getBugetUpdateBody(req)
	if err != nil {
		return buildServerError(err)
	}

	errRes = validateCategoryId(updateReq.CategoryId, req, api.categoryManager)
	if errRes != nil {
		return errRes
	}

	if err := api.budgetManager.Update(&budget, updateReq); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) DeleteBudget(req *data.RestRequest) *data.RestResponse {
	budget, errRes := checkBudget(req, api.budgetManager)
	if errRes != nil {
		return errRes
	}

	if err := api.budgetManager.Delete(budget.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}
