package store

/*
func TestTransactionStore(t *testing.T) {
	it := dbIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(it)
		transaction := testTransaction()

		err := it.TransactionStore.Insert(transaction)
		assert.NoError(t, err)

		fetched, err := it.TransactionStore.Get(transaction.Id)
		assert.NoError(t, err)
		assert.Equal(t, transaction, *fetched)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(it)
		transaction := testTransaction()
		transaction.Year = 1999

		it.TransactionStore.Insert(transaction)

		transaction.Year = 2001

		err := it.TransactionStore.Update(transaction)
		assert.NoError(t, err)

		fetched, _ := it.TransactionStore.Get(transaction.Id)
		assert.Equal(t, transaction, *fetched)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(it)
		transaction := testTransaction()

		it.TransactionStore.Insert(transaction)

		err := it.TransactionStore.Delete(transaction.Id)
		assert.NoError(t, err)

		fetched, _ := it.TransactionStore.Get(transaction.Id)
		assert.Nil(t, fetched)
	})

	t.Run("it gets transactions by account", func(t *testing.T) {
		setup(it)

		accountId := "test-12345"

		t1 := testTransaction()
		t1.Id = "t-1"
		t1.AccountId = accountId

		t2 := testTransaction()
		t2.Id = "t-2"
		t2.AccountId = accountId

		t3 := testTransaction()
		t3.Id = "t-3"
		t3.AccountId = "notid"

		it.TransactionStore.Insert(t1)
		it.TransactionStore.Insert(t2)
		it.TransactionStore.Insert(t3)

		expected := []data.Transaction{t1, t2}

		actual, err := it.TransactionStore.GetByAccount(accountId)
		assert.NoError(t, err)
		assert.ElementsMatch(t, actual, expected)
	})

}
*/
