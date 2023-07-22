package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionStore(t *testing.T) {
	it := NewStoreIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		it.Setup()

		transaction := testTransaction()

		//fmt.Println(transaction.CategoryId.Empty())

		err := it.TransactionStore.Insert(transaction)
		fmt.Println(err)
		assert.NoError(t, err)

		//found, err := it.TransactionStore.Get(transaction.Id, transaction.AccountId)
		//assert.NoError(t, err)
		//assert.Equal(t, transaction, found.Get())
	})

	/*t.Run("it gets by period category", func(t *testing.T) {
		it.Setup()

		month := 10
		year := 2023
		categoryId := types.OptionalString("some-category-id")

		t1 := testTransaction()
		t1.Id = "id-1"
		t1.Month = month
		t1.Year = year
		t1.CategoryId = categoryId

		err := it.TransactionStore.Insert(t1)
		assert.NoError(t, err)

		t2 := testTransaction()
		t2.Id = "id-2"
		t2.Month = month
		t2.Year = year
		t2.CategoryId = categoryId

		err = it.TransactionStore.Insert(t2)
		assert.NoError(t, err)

		transactions := []data.Transaction{t1, t2}

		found, err := it.TransactionStore.GetByPeriodCategory(t1.AccountId, categoryId.Get(), month, year)
		assert.NoError(t, err)
		assert.ElementsMatch(t, found, transactions)
	})

	t.Run("it gets by filter", func(t *testing.T) {
		it.Setup()

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
		assert.Equal(t, found.Get().Amount, newAmount)
	})*/
}
