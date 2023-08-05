package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type BudgetAPI struct {
	budgetManager      manager.BudgetManager
	transactionManager manager.TransactionManager
	categoryManager    manager.CategoryManager
}

func NewBudgetAPI(
	budgetManager manager.BudgetManager,
	transactionManager manager.TransactionManager,
) *BudgetAPI {
	return &BudgetAPI{
		budgetManager:      budgetManager,
		transactionManager: transactionManager,
	}
}

func getBudgetId(req *rest.Request) string {
	return ""
}

func (api *BudgetAPI) Get(req *rest.Request) *rest.Response {
	id := getBudgetId(req)
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
	budgetId := getBudgetId(req)

	body, err := rest.ParseBody[rest.BudgetUpdateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	existing, err := api.budgetManager.Get(budgetId, accountId)
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
	id := getBudgetId(req)

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
	accountId := budget.AccountId
	categoryId := budget.CategoryId
	month := budget.Month
	year := budget.Year

	total := 0.00
	transactions, err := api.transactionManager.GetByPeriodCategory(accountId, categoryId, month, year)

	if err != nil {
		return data.BudgetMaterialized{}, err

	}

	for _, transaction := range transactions {
		netMove := transaction.Amount
		if transaction.Type == data.BudgetType_Expense {
			netMove *= -1
		}

		total += netMove
	}

	m := data.BudgetMaterialized{
		Budget: budget,
		Actual: total,
	}

	return m, nil
}
