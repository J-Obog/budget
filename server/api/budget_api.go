package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type BudgetAPI struct {
	budgetManager      *manager.BudgetManager
	transactionManager *manager.TransactionManager
}

func NewBudgetAPI(budgetManager *manager.BudgetManager, transactionManager *manager.TransactionManager) *BudgetAPI {
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

	if budget == nil {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	total, err := api.getBudgetTotal(*budget)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(data.BudgetMaterialized{
		Budget: *budget,
		Actual: total,
	})
}

func (api *BudgetAPI) Filter(req *rest.Request) *rest.Response {
	accountId := testAccountId

	query, err := rest.ParseQuery[rest.BudgetQuery](req.Query)
	if err != nil {
		return rest.Err(err)
	}

	budgets, err := api.budgetManager.GetByQuery(accountId, query)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(budgets)
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

	budget, err := api.budgetManager.Create(accountId, body)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(budget)
}

func (api *BudgetAPI) Update(req *rest.Request) *rest.Response {
	accountId := testAccountId
	budgetId := req.ResourceId

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

	ok, err := api.budgetManager.Update(existing, body)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Ok(existing)
}

func (api *BudgetAPI) Delete(req *rest.Request) *rest.Response {
	accountId := testAccountId

	ok, err := api.budgetManager.Delete(req.ResourceId, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Success()
}

func (api *BudgetAPI) validateCreate(body rest.BudgetCreateBody, accountId string) error {
	if err := isDateValid(body.Month, 1, body.Year); err != nil {
		return err
	}

	if err := api.checkCategoryExists(body.CategoryId, accountId); err != nil {
		return err
	}

	if err := api.checkCategoryIsUnique(body.CategoryId, accountId, body.Month, body.Year); err != nil {
		return err
	}

	return nil
}

func (api *BudgetAPI) validateUpdate(existing *data.Budget, body rest.BudgetUpdateBody) error {
	if existing == nil {
		return rest.ErrInvalidBudgetId
	}

	if body.CategoryId != existing.CategoryId {
		if err := api.checkCategoryExists(body.CategoryId, existing.AccountId); err != nil {
			return err
		}

		if err := api.checkCategoryIsUnique(body.CategoryId, existing.AccountId, existing.Month, existing.Year); err != nil {
			return err
		}
	}

	return nil
}

func (api *BudgetAPI) checkCategoryIsUnique(categoryId string, accountId string, month int, year int) error {
	budget, err := api.budgetManager.GetByPeriodCategory(accountId, categoryId, month, year)
	if err != nil {
		return err
	}

	if budget != nil {
		return rest.ErrCategoryAlreadyInBudgetPeriod
	}

	return nil
}

func (api *BudgetAPI) checkCategoryExists(categoryId string, accountId string) error {
	category, err := api.budgetManager.Get(categoryId, accountId)
	if err != nil {
		return err
	}

	if category == nil {
		return rest.ErrInvalidCategoryId
	}

	return nil
}

func (api *BudgetAPI) getBudgetTotal(budget data.Budget) (float64, error) {
	accountId := budget.AccountId
	categoryId := budget.CategoryId
	month := budget.Month
	year := budget.Year

	total := 0.00
	transactions, err := api.transactionManager.GetByPeriodCategory(accountId, categoryId, month, year)

	if err != nil {
		return total, err

	}

	for _, transaction := range transactions {
		netMove := transaction.Amount
		if transaction.Type == data.BudgetType_Expense {
			netMove *= -1
		}

		total += netMove
	}

	return total, nil
}
