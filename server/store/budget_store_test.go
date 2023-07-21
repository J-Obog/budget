package store

/*
func TestBudgetStore(t *testing.T) {
	it := dbIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(it)
		budget := testBudget()

		err := it.BudgetStore.Insert(budget)
		assert.NoError(t, err)

		fetched, err := it.BudgetStore.Get(budget.Id)
		assert.NoError(t, err)
		assert.Equal(t, budget, *fetched)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(it)
		budget := testBudget()
		budget.Month = 10

		it.BudgetStore.Insert(budget)

		budget.Month = 9

		err := it.BudgetStore.Update(budget)
		assert.NoError(t, err)

		fetched, _ := it.BudgetStore.Get(budget.Id)
		assert.Equal(t, budget, *fetched)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(it)
		budget := testBudget()

		it.BudgetStore.Insert(budget)

		err := it.BudgetStore.Delete(budget.Id)
		assert.NoError(t, err)

		fetched, _ := it.BudgetStore.Get(budget.Id)
		assert.Nil(t, fetched)
	})
}
*/
