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
		assert.Equal(t, budget, found.Get())
	})

	t.Run("it gets by period category", func(t *testing.T) {
		it.Setup()

		categoryId := "some-category-id"

		b1 := testBudget()
		b1.Id = "test-id-1"
		b1.CategoryId = categoryId

		b2 := testBudget()
		b2.Id = "test-id-2"
		b2.CategoryId = "some-other-id"

		err := it.BudgetStore.Insert(b1)
		assert.NoError(t, err)

		err = it.BudgetStore.Insert(b2)
		assert.NoError(t, err)

		found, err := it.BudgetStore.GetByPeriodCategory(b1.AccountId, b1.CategoryId, b1.Month, b1.Year)
		assert.NoError(t, err)
		assert.Equal(t, b1, found.Get())
	})

	t.Run("it gets by category", func(t *testing.T) {
		it.Setup()

		categoryId := "some-category-id"

		b1 := testBudget()
		b1.Id = "test-id-1"
		b1.CategoryId = categoryId

		b2 := testBudget()
		b2.Id = "test-id-2"
		b2.CategoryId = categoryId

		b3 := testBudget()
		b3.Id = "test-id-3"
		b3.CategoryId = categoryId

		budgets := []data.Budget{b1, b2, b3}

		for _, budget := range budgets {
			err := it.BudgetStore.Insert(budget)
			assert.NoError(t, err)
		}

		found, err := it.BudgetStore.GetByCategory(b1.AccountId, categoryId)
		assert.NoError(t, err)
		assert.ElementsMatch(t, found, budgets)
	})

	t.Run("it gets by filter", func(t *testing.T) {
		it.Setup()

		filterMonth := 10
		otherMonth := 12

		year := 2023

		b1 := testBudget()
		b1.Id = "test-id-1"
		b1.Month = filterMonth
		b1.Year = year

		err := it.BudgetStore.Insert(b1)
		assert.NoError(t, err)

		b2 := testBudget()
		b2.Id = "test-id-2"
		b2.Month = otherMonth
		b2.Year = year

		err = it.BudgetStore.Insert(b2)
		assert.NoError(t, err)

		filter := data.BudgetFilter{
			Month: filterMonth,
			Year:  year,
		}

		expected := []data.Budget{b1}

		found, err := it.BudgetStore.GetBy(b1.AccountId, filter)
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
		assert.Equal(t, found.Get().Projected, newProjected)
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
		assert.True(t, found.Empty())
	})
}
