package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/assert"
)

func TestBudgetManagerGetsByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := bugdetManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		budget := data.Budget{Id: "budget-id-1", CategoryId: "some-cat-id"}

		relatedTransactions := []data.Transaction{
			{Type: data.BudgetType_Income, Amount: 50.00},
			{Type: data.BudgetType_Expense, Amount: 45.12},
			{Type: data.BudgetType_Income, Amount: 32.17},
		}

		expectedTotal := 37.05
		expectedResponse := getExpectedMaterializedBudget(budget, expectedTotal)

		manager.MockBudgetStore.On("Get", req.ResourceId, req.Account.Id).Return(&budget, nil)
		manager.MockTransactionStore.On("GetByPeriodCategory", req.Account.Id, budget.CategoryId, budget.Month, budget.Year).Return(relatedTransactions, nil)

		res := manager.GetByRequest(req)
		assert.Equal(t, res.Data, expectedResponse)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if budget doesn't exist", func(t *testing.T) {
		manager := bugdetManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		manager.MockBudgetStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.GetByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})
}

func TestBudgetManagerGetsAllByRequest(t *testing.T) {
	t.Run("it succeeds with user query", func(t *testing.T) {

	})

	t.Run("it succeeds with default values", func(t *testing.T) {

	})
}

func TestBudgetManagerCreatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetCreateBody{
			CategoryId: "cat-id",
			Projected:  100.50,
			Month:      10,
			Year:       2023,
		}

		req := testRequest()
		req.Body = body

		d := data.NewDate(body.Month, 1, body.Year)
		expected := getExpectedCreatedBudget(body, req.Account.Id)
		existing := data.Category{Id: "cat-id"}

		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockCategoryStore.On("Get", body.CategoryId, req.Account.Id).Return(&existing, nil)
		manager.MockBudgetStore.On("GetByPeriodCategory", req.Account.Id, body.CategoryId, body.Month, body.Year).Return(nil, nil)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockUid.On("GetId").Return(testUuid)
		manager.MockBudgetStore.On("Insert", expected).Return(nil)

		res := manager.CreateByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if period is invalid", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetCreateBody{
			Month: 10,
			Year:  2023,
		}

		req := testRequest()
		req.Body = body

		d := data.NewDate(body.Month, 1, body.Year)

		manager.MockClock.On("IsDateValid", d).Return(false)

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidDate)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetCreateBody{
			CategoryId: "cat-id",
		}

		req := testRequest()
		req.Body = body

		d := data.NewDate(body.Month, 1, body.Year)

		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockCategoryStore.On("Get", body.CategoryId, req.Account.Id).Return(nil, nil)

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

	t.Run("it fails if category id is already in use", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetCreateBody{
			CategoryId: "cat-id",
		}

		req := testRequest()
		req.Body = body

		d := data.NewDate(body.Month, 1, body.Year)
		existing := data.Category{Id: "cat-id"}
		budgetThatHasCategory := data.Budget{Id: "cat-id"}

		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockCategoryStore.On("Get", body.CategoryId, req.Account.Id).Return(&existing, nil)
		manager.MockBudgetStore.On("GetByPeriodCategory", req.Account.Id, body.CategoryId, body.Month, body.Year).Return(&budgetThatHasCategory, nil)

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrCategoryAlreadyInBudgetPeriod)
	})
}

func TestBudgetManagerUpdatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetUpdateBody{
			CategoryId: "new-category-id",
		}

		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		expected := getExpectedBudgetUpdate(body)
		category := data.Category{Id: "new-category-id"}
		existingBudget := data.Budget{CategoryId: "old-category-id", Month: 10, Year: 2021}

		manager.MockCategoryStore.On("Get", body.CategoryId, req.Account.Id).Return(&category, nil)
		manager.MockBudgetStore.On("Get", req.ResourceId, req.Account.Id).Return(&existingBudget, nil)
		manager.MockBudgetStore.On("GetByPeriodCategory", req.Account.Id, body.CategoryId, existingBudget.Month, existingBudget.Year).Return(nil, nil)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockUid.On("GetId").Return(testUuid)
		manager.MockBudgetStore.On("Update", req.ResourceId, req.Account.Id, expected, testTimestamp).Return(true, nil)

		res := manager.UpdateByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if budget wasn't updated", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetUpdateBody{
			CategoryId: "new-category-id",
		}

		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		expected := getExpectedBudgetUpdate(body)
		category := data.Category{Id: "new-category-id"}
		existingBudget := data.Budget{CategoryId: "old-category-id", Month: 10, Year: 2021}

		manager.MockCategoryStore.On("Get", body.CategoryId, req.Account.Id).Return(&category, nil)
		manager.MockBudgetStore.On("Get", req.ResourceId, req.Account.Id).Return(&existingBudget, nil)
		manager.MockBudgetStore.On("GetByPeriodCategory", req.Account.Id, body.CategoryId, existingBudget.Month, existingBudget.Year).Return(nil, nil)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockUid.On("GetId").Return(testUuid)
		manager.MockBudgetStore.On("Update", req.ResourceId, req.Account.Id, expected, testTimestamp).Return(false, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidBudgetId)
	})

	t.Run("it fails if budget doesn't exist", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetUpdateBody{
			CategoryId: "new-category-id",
		}

		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		manager.MockBudgetStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidBudgetId)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetUpdateBody{
			CategoryId: "new-category-id",
		}

		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		existingBudget := data.Budget{CategoryId: "old-category-id", Month: 10, Year: 2021}

		manager.MockCategoryStore.On("Get", body.CategoryId, req.Account.Id).Return(nil, nil)
		manager.MockBudgetStore.On("Get", req.ResourceId, req.Account.Id).Return(&existingBudget, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

	t.Run("it fails if category id is already in use", func(t *testing.T) {
		manager := bugdetManagerMock()
		body := rest.BudgetUpdateBody{
			CategoryId: "new-category-id",
		}

		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		category := data.Category{Id: "new-category-id"}
		existingBudget := data.Budget{CategoryId: "old-category-id", Month: 10, Year: 2021}
		budgetThatHasCategoryId := data.Budget{CategoryId: "new-category-id"}

		manager.MockCategoryStore.On("Get", body.CategoryId, req.Account.Id).Return(&category, nil)
		manager.MockBudgetStore.On("Get", req.ResourceId, req.Account.Id).Return(&existingBudget, nil)
		manager.MockBudgetStore.On("GetByPeriodCategory", req.Account.Id, body.CategoryId, existingBudget.Month, existingBudget.Year).Return(&budgetThatHasCategoryId, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrCategoryAlreadyInBudgetPeriod)
	})
}

func TestBudgetManagerUDeletesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := bugdetManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		manager.MockBudgetStore.On("Delete", req.ResourceId, req.Account.Id).Return(true, nil)

		res := manager.DeleteByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if category wasn't deleted", func(t *testing.T) {
		manager := bugdetManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		manager.MockBudgetStore.On("Delete", req.ResourceId, req.Account.Id).Return(false, nil)

		res := manager.DeleteByRequest(req)
		assert.Error(t, res.Error, rest.ErrInvalidBudgetId)
	})
}

func getExpectedMaterializedBudget(budget data.Budget, actual float64) data.BudgetMaterialized {
	return data.BudgetMaterialized{
		Budget: budget,
		Actual: actual,
	}
}

func getExpectedCreatedBudget(body rest.BudgetCreateBody, accountId string) data.Budget {
	return data.Budget{
		Id:         testUuid,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Month:      body.Month,
		Year:       body.Year,
		Projected:  body.Projected,
		CreatedAt:  testTimestamp,
		UpdatedAt:  testTimestamp,
	}
}

func getExpectedBudgetUpdate(body rest.BudgetUpdateBody) data.BudgetUpdate {
	return data.BudgetUpdate{
		CategoryId: body.CategoryId,
		Projected:  body.Projected,
	}
}
