package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func TestBudgetStore(t *testing.T) {
	it := NewStoreIntegrationTest()
	t.Run("it inserts and gets", func(t *testing.T) {
		it.Setup()

		budget := testBudget()

		err := it.BudgetStore.Insert(budget)
		assert.NoError(t, err)

		found, err := it.BudgetStore.Get(budget.Id, budget.AccountId)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, budget, *found)
	})

	t.Run("it gets by period category", func(t *testing.T) {
		it.Setup()

		categoryId := "some-category-id"
		month := 10
		year := 2024

		budget := testBudget()
		budget.Id = "test-id-1"
		budget.CategoryId = categoryId
		budget.Month = month
		budget.Year = year

		err := it.BudgetStore.Insert(budget)
		assert.NoError(t, err)

		found, err := it.BudgetStore.GetByPeriodCategory(budget.AccountId, categoryId, month, year)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, budget, *found)
	})

	t.Run("it gets by category", func(t *testing.T) {
		it.Setup()

		categoryId := "some-category-id"
		accountId := "some-account-id"

		b1 := testBudget()
		b1.Id = "test-id-1"
		b1.AccountId = accountId
		b1.CategoryId = categoryId

		b2 := testBudget()
		b2.Id = "test-id-2"
		b2.AccountId = accountId
		b2.CategoryId = categoryId

		expected := []data.Budget{b1, b2}

		for _, budget := range expected {
			err := it.BudgetStore.Insert(budget)
			assert.NoError(t, err)
		}

		found, err := it.BudgetStore.GetByCategory(accountId, categoryId)
		assert.NoError(t, err)
		assert.ElementsMatch(t, found, expected)
	})

	t.Run("it gets by filter", func(t *testing.T) {
		it.Setup()

		month := 10
		year := 2023
		accountId := "some-account-id"

		b1 := testBudget()
		b1.Id = "test-id-1"
		b1.AccountId = accountId
		b1.Month = month
		b1.Year = year

		b2 := testBudget()
		b2.Id = "test-id-2"
		b2.AccountId = accountId
		b2.Month = month
		b2.Year = year

		expected := []data.Budget{b1, b2}

		for _, budget := range expected {
			err := it.BudgetStore.Insert(budget)
			assert.NoError(t, err)
		}

		filter := data.BudgetFilter{
			Month: month,
			Year:  year,
		}

		found, err := it.BudgetStore.GetBy(accountId, filter)
		assert.NoError(t, err)
		assert.ElementsMatch(t, found, expected)
	})

	t.Run("it updates", func(t *testing.T) {
		it.Setup()

		oldProjected := float64(1234)
		newProjected := float64(4321)

		budget := testBudget()
		budget.Projected = oldProjected

		err := it.BudgetStore.Insert(budget)
		assert.NoError(t, err)

		update := data.BudgetUpdate{
			Projected: newProjected,
		}

		ok, err := it.BudgetStore.Update(budget.Id, budget.AccountId, update, 12345)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.BudgetStore.Get(budget.Id, budget.AccountId)
		assert.NoError(t, err)
		assert.Equal(t, found.Projected, newProjected)
	})

	t.Run("it deletes", func(t *testing.T) {
		it.Setup()

		budget := testBudget()

		err := it.BudgetStore.Insert(budget)
		assert.NoError(t, err)

		ok, err := it.BudgetStore.Delete(budget.Id, budget.AccountId)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.BudgetStore.Get(budget.Id, budget.AccountId)
		assert.NoError(t, err)
		assert.Nil(t, found)
	})
}
