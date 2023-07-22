package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
	"github.com/stretchr/testify/assert"
)

func TestTransactionStore(t *testing.T) {
	it := NewStoreIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		it.Setup()

		transaction := testTransaction()

		err := it.TransactionStore.Insert(transaction)
		assert.NoError(t, err)

		found, err := it.TransactionStore.Get(transaction.Id, transaction.AccountId)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, transaction, *found)
	})

	t.Run("it gets by period category", func(t *testing.T) {
		it.Setup()

		month := 10
		year := 2023
		categoryId := types.StringPtr("some-category-id")

		t1 := testTransaction()
		t1.Id = "id-1"
		t1.Month = month
		t1.Year = year
		t1.CategoryId = categoryId

		t2 := testTransaction()
		t2.Id = "id-2"
		t2.Month = month
		t2.Year = year
		t2.CategoryId = categoryId

		expected := []data.Transaction{t1, t2}

		for _, transaction := range expected {
			err := it.TransactionStore.Insert(transaction)
			assert.NoError(t, err)
		}

		found, err := it.TransactionStore.GetByPeriodCategory(t1.AccountId, *categoryId, month, year)
		assert.NoError(t, err)
		assert.ElementsMatch(t, found, expected)
	})

	t.Run("it gets by filter", func(t *testing.T) {
		it.Setup()

		accountId := "someId"

		lowestAmount := 10.90
		highestAmount := 99.99

		lowestDay := 10
		highestDay := 15

		month := 10
		year := 2021

		t1 := testTransaction()
		t1.Id = "id-1"
		t1.AccountId = accountId
		t1.Month = month
		t1.Year = year
		t1.Day = lowestDay
		t1.Amount = lowestAmount

		err := it.TransactionStore.Insert(t1)
		assert.NoError(t, err)

		t2 := testTransaction()
		t2.Id = "id-2"
		t2.AccountId = accountId
		t2.Month = month
		t2.Year = year
		t2.Day = highestDay
		t2.Amount = highestAmount + 10

		err = it.TransactionStore.Insert(t2)
		assert.NoError(t, err)

		expected := []data.Transaction{t1}

		filter := data.TransactionFilter{
			Before:      data.NewDate(month, highestDay, year),
			After:       data.NewDate(month, lowestDay, year),
			GreaterThan: lowestAmount,
			LessThan:    highestAmount,
		}

		found, err := it.TransactionStore.GetBy(accountId, filter)
		assert.NoError(t, err)
		assert.ElementsMatch(t, expected, found)
	})

	t.Run("it updates", func(t *testing.T) {
		it.Setup()

		oldAmount := 1.23
		newAmount := 2.23

		transaction := testTransaction()
		transaction.Amount = oldAmount

		err := it.TransactionStore.Insert(transaction)
		assert.NoError(t, err)

		update := data.TransactionUpdate{
			Amount: newAmount,
		}

		ok, err := it.TransactionStore.Update(transaction.Id, transaction.AccountId, update, 1234)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.TransactionStore.Get(transaction.Id, transaction.AccountId)
		assert.NoError(t, err)
		assert.Equal(t, found.Amount, newAmount)
	})

	t.Run("it deletes", func(t *testing.T) {
		it.Setup()

		transaction := testTransaction()

		err := it.TransactionStore.Insert(transaction)
		assert.NoError(t, err)

		ok, err := it.TransactionStore.Delete(transaction.Id, transaction.AccountId)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.TransactionStore.Get(transaction.Id, transaction.AccountId)
		assert.NoError(t, err)
		assert.Nil(t, found)
	})
}
