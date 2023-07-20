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
	accountId := req.Account.Get().Id
	budgetId := req.ResourceId

	budget, err := manager.store.Get(budgetId, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if budget.Empty() {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	materializedBudget, err := manager.toMaterializedBudget(budget.Get())
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(materializedBudget)
}

func (manager *BudgetManager) GetAllByRequest(req *rest.Request) *rest.Response {
	accountId := req.Account.Get().Id

	query := req.Query.(rest.BudgetQuery)

	filter := manager.getFilterForBudgetQuery(query)

	budgets, err := manager.store.GetBy(accountId, filter)
	if err != nil {
		return rest.Err(err)
	}

	materializedBudgets := make([]data.BudgetMaterialized, len(budgets))
	for _, budget := range budgets {
		materializedBudget, err := manager.toMaterializedBudget(budget)
		if err != nil {
			return rest.Err(err)
		}

		materializedBudgets = append(materializedBudgets, materializedBudget)
	}

	return rest.Ok(materializedBudgets)
}

func (manager *BudgetManager) CreateByRequest(req *rest.Request) *rest.Response {
	body := req.Body.(rest.BudgetCreateBody)
	accountId := req.Account.Get().Id

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
	accountId := req.Account.Get().Id
	budgetId := req.ResourceId

	timestamp := manager.clock.Now()
	update := getUpdateForBudgetUpdate(body)

	if err := manager.validateUpdate(body, budgetId, accountId); err != nil {
		return rest.Err(err)
	}

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
	ok, err := manager.store.Delete(req.ResourceId, req.Account.Get().Id)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidBudgetId)
	}

	return rest.Success()
}

func (manager *BudgetManager) validateSet(body rest.BudgetSetBody, accountId string, month int, year int) error {
	if err := manager.checkCategoryExists(body.CategoryId, accountId); err != nil {
		return err
	}

	if err := manager.checkCategoryIsUnique(body.CategoryId, accountId, month, year); err != nil {
		return err
	}

	return nil
}

func (manager *BudgetManager) validateCreate(body rest.BudgetCreateBody, accountId string) error {
	period := data.NewDate(body.Month, 1, body.Year)

	if ok := manager.clock.IsDateValid(period); !ok {
		return rest.ErrInvalidDate
	}

	if err := manager.validateSet(body.BudgetSetBody, accountId, period.Month, period.Year); err != nil {
		return err
	}

	return nil
}

func (manager *BudgetManager) validateUpdate(body rest.BudgetUpdateBody, budgetId string, accountId string) error {
	budget, err := manager.store.Get(budgetId, accountId)
	if err != nil {
		return err
	}

	if budget.Empty() {
		return rest.ErrInvalidBudgetId
	}

	if body.CategoryId != budget.Get().CategoryId {
		return manager.validateSet(body.BudgetSetBody, accountId, budget.Get().Month, budget.Get().Year)
	}

	return nil
}

func (manager *BudgetManager) checkCategoryIsUnique(categoryId string, accountId string, month int, year int) error {
	budget, err := manager.store.GetByPeriodCategory(accountId, categoryId, month, year)
	if err != nil {
		return err
	}

	if budget.Empty() {
		return rest.ErrCategoryAlreadyInBudgetPeriod
	}

	return nil
}

func (manager *BudgetManager) checkCategoryExists(categoryId string, accountId string) error {
	category, err := manager.categoryStore.Get(categoryId, accountId)
	if err != nil {
		return err
	}

	if category.Empty() {
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
	return data.BudgetFilter{
		Month: q.Month.GetOr(manager.clock.CurrentMonth()),
		Year:  q.Year.GetOr(manager.clock.CurrentYear()),
	}
}

func getUpdateForBudgetUpdate(body rest.BudgetUpdateBody) data.BudgetUpdate {
	return data.BudgetUpdate{
		CategoryId: body.CategoryId,
		Projected:  body.Projected,
	}
}

func (manager *BudgetManager) toMaterializedBudget(budget data.Budget) (data.BudgetMaterialized, error) {
	accountId := budget.AccountId
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
