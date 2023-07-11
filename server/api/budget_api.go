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
	budgets, err := api.budgetManager.Filter(r.Account.Id, r.Query.BudgetQuery())
	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(budgets)
}

func (api *BudgetAPI) CreateBudget(r *rest.Request) *rest.Response {
	reqBody, err := r.Body.BudgetCreateBody()
	if err != nil {
		return buildBadRequestError()
	}

	if err := api.validateCreate(&reqBody, r.Account.Id); err != nil {
		return buildBadRequestError()
	}

	if err := api.budgetManager.Create(r.Account.Id, reqBody); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *BudgetAPI) UpdateBudget(r *rest.Request) *rest.Response {
	budget, errResp := api.bugetCtx(r)
	if errResp != nil {
		return errResp
	}

	reqBody, err := r.Body.BudgetUpdateBody()
	if err != nil {
		return buildBadRequestError()
	}

	if err := api.validateUpdate(&reqBody, r.Account.Id); err != nil {
		return buildBadRequestError()
	}

	if err := api.budgetManager.Update(&budget, reqBody); err != nil {
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

func (api *BudgetAPI) validateCreate(reqBody *rest.BudgetCreateBody, accountId string) error {
	month := reqBody.Month
	year := reqBody.Year
	categoryId := reqBody.CategoryId

	if err := checkDate(month, 1, year); err != nil {
		return err
	}

	ok, err := api.budgetManager.CategoryInPeriod(categoryId, accountId, month, year)
	if err != nil {
		return err
	}
	if ok {
		return nil //UPDATE TO SOME BAD REQ ERROR
	}

	ok, err = api.categoryManager.Exists(categoryId, accountId)
	if err != nil {
		return err
	}
	if !ok {
		return nil //UPDATE TO SOME BAD REQ ERROR
	}

	return nil
}

func (api *BudgetAPI) validateUpdate(reqBody *rest.BudgetUpdateBody, accountId string) error {
	return api.validateCreate(&reqBody.BudgetCreateBody, accountId)
}
