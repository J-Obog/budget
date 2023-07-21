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
		assert.True(t, found.NotEmpty())
		assert.Equal(t, category, found.Get())
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
		assert.True(t, found.NotEmpty())
		assert.Equal(t, found.Get().Name, newName)
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
		assert.True(t, found.NotEmpty())
		assert.Equal(t, found.Get(), category)
	})

	t.Run("it gets all", func(t *testing.T) {
		it.Setup()

		c1 := testCategory()
		c1.Id = "t1"

		c2 := testCategory()
		c2.Id = "t2"

		c3 := testCategory()
		c3.Id = "t3"

		categories := []data.Category{c1, c2, c3}

		for _, category := range categories {
			err := it.CategoryStore.Insert(category)
			assert.NoError(t, err)
		}

		found, err := it.CategoryStore.GetAll(c1.AccountId)
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
		assert.True(t, found.Empty())
	})
}
