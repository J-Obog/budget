package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func TestCategories(t *testing.T) {
	store := getStore(t)

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(t, store)

		testId := "test-1"

		category := makeCategory()
		category.Id = testId

		err := store.InsertCategory(category)
		assert.NoError(t, err)

		fetchedCategory, err := store.GetCategory(testId)
		assert.NoError(t, err)
		assert.Equal(t, category, *fetchedCategory)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(t, store)

		oldColorCode := 1233445
		newColorCode := 4556652

		testId := "t-123456"

		category := makeCategory()
		category.Id = testId
		category.Color = oldColorCode

		store.InsertCategory(category)

		category.Color = newColorCode

		err := store.UpdateCategory(category)
		assert.NoError(t, err)

		fetchedCategory, _ := store.GetCategory(testId)
		assert.Equal(t, category, *fetchedCategory)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(t, store)

		testId := "t-1"

		category := makeCategory()
		category.Id = testId

		store.InsertCategory(category)

		err := store.DeleteCategory(testId)
		assert.NoError(t, err)

		fetchedCategory, _ := store.GetCategory(testId)
		assert.Nil(t, fetchedCategory)
	})

	t.Run("it gets categories by account", func(t *testing.T) {
		setup(t, store)

		accountId := "test-12345"

		cat1 := makeCategory()
		cat1.Id = "t-1"
		cat1.AccountId = accountId

		cat2 := makeCategory()
		cat2.Id = "t-2"
		cat2.AccountId = accountId

		cat3 := makeCategory()
		cat3.Id = "t-3"
		cat3.AccountId = accountId

		store.InsertCategory(cat1)
		store.InsertCategory(cat2)
		store.InsertCategory(cat3)

		expected := []data.Category{cat1, cat2, cat3}

		actual, err := store.GetCategories(accountId)
		assert.NoError(t, err)
		assert.ElementsMatch(t, actual, expected)
	})
}
