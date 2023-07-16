package api

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type BudgetAPI struct {
	budgetManager *manager.BudgetManager
}

func (api *BudgetAPI) GetBudget(req *rest.Request, res *rest.Response) {
	api.budgetManager.GetByRequest(req, res)
}

func (api *BudgetAPI) GetBudgets(req *rest.Request, res *rest.Response) {
	api.budgetManager.GetAllByRequest(req, res)
}

func (api *BudgetAPI) CreateBudget(req *rest.Request, res *rest.Response) {
	api.budgetManager.CreateByRequest(req, res)
}

func (api *BudgetAPI) UpdateBudget(req *rest.Request, res *rest.Response) {
	api.budgetManager.UpdateByRequest(req, res)
}

func (api *BudgetAPI) DeleteBudget(req *rest.Request, res *rest.Response) {
	api.budgetManager.DeleteByRequest(req, res)
}
