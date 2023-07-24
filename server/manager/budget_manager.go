package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type BudgetManager struct {
	store            store.BudgetStore
	categoryStore    store.CategoryStore
	transactionStore store.TransactionStore
	clock            clock.Clock
	uid              uid.UIDProvider
}

func (manager *BudgetManager) GetByRequest(req *rest.Request) *rest.Response {
	accountId := req.Account.Id
	budgetId := req.ResourceId

	budget, err := manager.store.Get(budgetId, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if budget == nil {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	materializedBudget, err := manager.toMaterializedBudget(*budget, accountId)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(materializedBudget)
}

func (manager *BudgetManager) GetAllByRequest(req *rest.Request) *rest.Response {
	accountId := req.Account.Id

	query := req.Query.(rest.BudgetQuery)

	filter := manager.getFilterForBudgetQuery(query)

	budgets, err := manager.store.GetBy(accountId, filter)
	if err != nil {
		return rest.Err(err)
	}

	materializedBudgets := make([]data.BudgetMaterialized, len(budgets))
	for _, budget := range budgets {
		materializedBudget, err := manager.toMaterializedBudget(budget, accountId)
		if err != nil {
			return rest.Err(err)
		}

		materializedBudgets = append(materializedBudgets, materializedBudget)
	}

	return rest.Ok(materializedBudgets)
}

func (manager *BudgetManager) CreateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.BudgetCreateBody)
	accountId := req.Account.Id

	if err := manager.validateCreate(body, accountId); err != nil {
		return rest.Err(err)
	}

	budget := manager.getBudgetForCreate(accountId, body)

	if err := manager.store.Insert(budget); err != nil {
		return rest.Err(err)
	}

	return rest.Success()
}

func (manager *BudgetManager) UpdateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.BudgetUpdateBody)
	accountId := req.Account.Id
	budgetId := req.ResourceId

	if err := manager.validateUpdate(body, budgetId, accountId); err != nil {
		return rest.Err(err)
	}

	timestamp := manager.clock.Now()
	update := getUpdateForBudgetUpdate(body)

	ok, err := manager.store.Update(budgetId, accountId, update, timestamp)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Success()
}

func (manager *BudgetManager) DeleteByRequest(req *rest.Request) *rest.Response {
	ok, err := manager.store.Delete(req.ResourceId, req.Account.Id)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Success()
}

func (manager *BudgetManager) validateCreate(body rest.BudgetCreateBody, accountId string) error {
	if err := manager.checkPeriod(body.Month, body.Year); err != nil {
		return err
	}

	if err := manager.checkCategoryExists(body.CategoryId, accountId); err != nil {
		return err
	}

	if err := manager.checkCategoryIsUnique(body.CategoryId, accountId, body.Month, body.Year); err != nil {
		return err
	}

	return nil
}

func (manager *BudgetManager) validateUpdate(body rest.BudgetUpdateBody, budgetId string, accountId string) error {
	budget, err := manager.store.Get(budgetId, accountId)
	if err != nil {
		return err
	}

	if budget == nil {
		return rest.ErrInvalidBudgetId
	}

	if body.CategoryId != budget.CategoryId {
		if err := manager.checkCategoryExists(body.CategoryId, accountId); err != nil {
			return err
		}

		if err := manager.checkCategoryIsUnique(body.CategoryId, accountId, budget.Month, budget.Year); err != nil {
			return err
		}
	}

	return nil
}

func (manager *BudgetManager) checkPeriod(month int, year int) error {
	d := data.NewDate(month, 1, year)

	if ok := manager.clock.IsDateValid(d); !ok {
		return rest.ErrInvalidDate
	}

	return nil
}

func (manager *BudgetManager) checkCategoryIsUnique(categoryId string, accountId string, month int, year int) error {
	budget, err := manager.store.GetByPeriodCategory(accountId, categoryId, month, year)
	if err != nil {
		return err
	}

	if budget != nil {
		return rest.ErrCategoryAlreadyInBudgetPeriod
	}

	return nil
}

func (manager *BudgetManager) checkCategoryExists(categoryId string, accountId string) error {
	category, err := manager.categoryStore.Get(categoryId, accountId)
	if err != nil {
		return err
	}

	if category == nil {
		return rest.ErrInvalidCategoryId
	}

	return nil
}

func (manager *BudgetManager) getBudgetForCreate(accountId string, body rest.BudgetCreateBody) data.Budget {
	id := manager.uid.GetId()
	timestamp := manager.clock.Now()

	return data.Budget{
		Id:         id,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Month:      body.Month,
		Year:       body.Year,
		Projected:  body.Projected,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}
}

func (manager *BudgetManager) getFilterForBudgetQuery(q rest.BudgetQuery) data.BudgetFilter {
	filter := data.BudgetFilter{
		Month: manager.clock.CurrentMonth(),
		Year:  manager.clock.CurrentYear(),
	}

	if q.Month != nil {
		filter.Month = *q.Month
	}

	if q.Year != nil {
		filter.Year = *q.Year
	}

	return filter
}

func getUpdateForBudgetUpdate(body rest.BudgetUpdateBody) data.BudgetUpdate {
	return data.BudgetUpdate{
		CategoryId: body.CategoryId,
		Projected:  body.Projected,
	}
}

func (manager *BudgetManager) toMaterializedBudget(budget data.Budget, accountId string) (data.BudgetMaterialized, error) {
	categoryId := budget.CategoryId
	month := budget.Month
	year := budget.Year

	total := 0.00
	transactions, err := manager.transactionStore.GetByPeriodCategory(accountId, categoryId, month, year)

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

	materialized := data.BudgetMaterialized{
		Budget: budget,
		Actual: total,
	}

	return materialized, nil
}
