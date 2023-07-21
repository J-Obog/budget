package store

/*
func TestCategoryStore(t *testing.T) {
	it := dbIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(it)
		category := testCategory()

		err := it.CategoryStore.Insert(category)
		assert.NoError(t, err)

		fetched, err := it.CategoryStore.Get(category.Id)
		assert.NoError(t, err)
		assert.Equal(t, category, *fetched)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(it)
		category := testCategory()
		category.Color = 1234

		it.CategoryStore.Insert(category)

		category.Color = 56789

		err := it.CategoryStore.Update(category)
		assert.NoError(t, err)

		fetched, _ := it.CategoryStore.Get(category.Id)
		assert.Equal(t, category, *fetched)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(it)
		category := testCategory()

		it.CategoryStore.Insert(category)

		err := it.CategoryStore.Delete(category.Id)
		assert.NoError(t, err)

		fetched, _ := it.CategoryStore.Get(category.Id)
		assert.Nil(t, fetched)
	})

	t.Run("it gets categories by account", func(t *testing.T) {
		setup(it)

		accountId := "acc12345"

		cat1 := testCategory()
		cat1.Id = "t-1"
		cat1.AccountId = accountId

		cat2 := testCategory()
		cat2.Id = "t-2"
		cat2.AccountId = accountId

		cat3 := testCategory()
		cat3.Id = "t-3"
		cat3.AccountId = "notid"

		it.CategoryStore.Insert(cat1)
		it.CategoryStore.Insert(cat2)
		it.CategoryStore.Insert(cat3)

		expected := []data.Category{cat1, cat2}

		actual, err := it.CategoryStore.GetByAccount(accountId)
		assert.NoError(t, err)
		assert.ElementsMatch(t, actual, expected)
	})
}
*/
