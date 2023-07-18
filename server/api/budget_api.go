package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type BudgetAPI struct {
	budgetManager *manager.BudgetManager
}

func (api *BudgetAPI) GetBudget(req *rest.Request) *rest.Response {
	return api.budgetManager.GetByRequest(req)
}

func (api *BudgetAPI) GetBudgets(req *rest.Request) *rest.Response {
	return api.budgetManager.GetAllByRequest(req)
}

func (api *BudgetAPI) CreateBudget(req *rest.Request) *rest.Response {
	return api.budgetManager.CreateByRequest(req)
}

func (api *BudgetAPI) UpdateBudget(req *rest.Request) *rest.Response {
	return api.budgetManager.UpdateByRequest(req)
}

func (api *BudgetAPI) DeleteBudget(req *rest.Request) *rest.Response {
	return api.budgetManager.DeleteByRequest(req)
}
