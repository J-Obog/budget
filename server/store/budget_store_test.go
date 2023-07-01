package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBudgets(t *testing.T) {
	store := getStore(t)

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(t, store)

		testId := "test-1"

		budget := makeBudget()
		budget.Id = testId

		err := store.InsertBudget(budget)
		assert.NoError(t, err)

		fetchedBudget, err := store.GetBudget(testId)
		assert.NoError(t, err)
		assert.Equal(t, budget, *fetchedBudget)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(t, store)

		oldProjectedBudget := 12345
		newProjectedBudget := 50000

		testId := "t-123456"

		budget := makeBudget()
		budget.Id = testId
		budget.Projected = float64(oldProjectedBudget)

		store.InsertBudget(budget)

		budget.Projected = float64(newProjectedBudget)

		err := store.UpdateBudget(budget)
		assert.NoError(t, err)

		fetchedBudget, _ := store.GetBudget(testId)
		assert.Equal(t, budget, *fetchedBudget)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(t, store)

		testId := "t-1"

		budget := makeBudget()
		budget.Id = testId

		store.InsertBudget(budget)

		err := store.DeleteBudget(testId)
		assert.NoError(t, err)

		fetchedBudget, _ := store.GetBudget(testId)
		assert.Nil(t, fetchedBudget)
	})
}
