package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func TestTransactions(t *testing.T) {
	store := getStore(t)

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(t, store)

		testId := "test-1"

		transaction := makeTransaction()
		transaction.Id = testId

		err := store.InsertTransaction(transaction)
		assert.NoError(t, err)

		fetchedTransaction, err := store.GetTransaction(testId)
		assert.NoError(t, err)
		assert.Equal(t, transaction, *fetchedTransaction)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(t, store)

		oldAmount := float64(123)
		newAmount := float64(9000)

		testId := "t-123456"

		transaction := makeTransaction()
		transaction.Id = testId
		transaction.Amount = oldAmount

		store.InsertTransaction(transaction)

		transaction.Amount = newAmount

		err := store.UpdateTransaction(transaction)
		assert.NoError(t, err)

		fetchedTransaction, _ := store.GetTransaction(testId)
		assert.Equal(t, transaction, *fetchedTransaction)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(t, store)

		testId := "t-1"

		transaction := makeTransaction()
		transaction.Id = testId

		store.InsertTransaction(transaction)

		err := store.DeleteTransaction(testId)
		assert.NoError(t, err)

		fetchedTransaction, _ := store.GetTransaction(testId)
		assert.Nil(t, fetchedTransaction)
	})

	t.Run("it gets categories by account", func(t *testing.T) {
		setup(t, store)

		accountId := "test-12345"

		t1 := makeTransaction()
		t1.Id = "t-1"
		t1.AccountId = accountId

		t2 := makeTransaction()
		t2.Id = "t-2"
		t2.AccountId = accountId

		t3 := makeTransaction()
		t3.Id = "t-3"
		t3.AccountId = accountId

		store.InsertTransaction(t1)
		store.InsertTransaction(t2)
		store.InsertTransaction(t3)

		expected := []data.Transaction{t1, t2, t3}

		actual, err := store.GetTransactions(accountId)
		assert.NoError(t, err)
		assert.ElementsMatch(t, actual, expected)
	})

}
