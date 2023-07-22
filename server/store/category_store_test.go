package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func TestCategoryStore(t *testing.T) {
	it := NewStoreIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		it.Setup()

		category := testCategory()

		err := it.CategoryStore.Insert(category)
		assert.NoError(t, err)

		found, err := it.CategoryStore.Get(category.Id, category.AccountId)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, category, *found)
	})

	t.Run("it updates", func(t *testing.T) {
		it.Setup()

		category := testCategory()
		category.Name = "some_old_name"

		newName := "some_new_name"
		update := data.CategoryUpdate{Name: newName}

		err := it.CategoryStore.Insert(category)
		assert.NoError(t, err)

		ok, err := it.CategoryStore.Update(category.Id, category.AccountId, update, 1234)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.CategoryStore.Get(category.Id, category.AccountId)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, found.Name, newName)
	})

	t.Run("it gets by name", func(t *testing.T) {
		it.Setup()

		name := "some+cool+name"

		category := testCategory()
		category.Name = name

		err := it.CategoryStore.Insert(category)
		assert.NoError(t, err)

		found, err := it.CategoryStore.GetByName(category.AccountId, name)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, found, category)
	})

	t.Run("it gets all", func(t *testing.T) {
		it.Setup()

		accountId := "some-account-id"

		c1 := testCategory()
		c1.Id = "t1"
		c1.AccountId = accountId

		c2 := testCategory()
		c2.Id = "t2"
		c2.AccountId = accountId

		c3 := testCategory()
		c3.Id = "t3"
		c3.AccountId = accountId

		categories := []data.Category{c1, c2, c3}

		for _, category := range categories {
			err := it.CategoryStore.Insert(category)
			assert.NoError(t, err)
		}

		found, err := it.CategoryStore.GetAll(accountId)
		assert.NoError(t, err)
		assert.ElementsMatch(t, found, categories)
	})

	t.Run("it deletes", func(t *testing.T) {
		it.Setup()

		category := testCategory()

		err := it.CategoryStore.Insert(category)
		assert.NoError(t, err)

		ok, err := it.CategoryStore.Delete(category.Id, category.AccountId)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.CategoryStore.Get(category.Id, category.AccountId)
		assert.NoError(t, err)
		assert.Nil(t, found)
	})
}
