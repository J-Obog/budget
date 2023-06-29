package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func TestAccounts(t *testing.T) {
	store := getStore(t)

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(t, store)

		testId := "testing-123"

		account := makeAccount()
		account.Id = testId

		err := store.InsertAccount(account)
		assert.NoError(t, err)

		fetchedAccount, err := store.GetAccount(testId)
		assert.NoError(t, err)
		assert.Equal(t, account, *fetchedAccount)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(t, store)

		testId := "testing-123"

		oldEmail := "jdoe@gmail.com"
		newEmail := "jdoe@yahoo.com"

		account := makeAccount()
		account.Id = testId
		account.Email = oldEmail

		store.InsertAccount(account)

		account.Email = newEmail

		err := store.UpdateAccount(account)
		assert.NoError(t, err)

		fetchedAccount, _ := store.GetAccount(testId)
		assert.Equal(t, account, *fetchedAccount)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(t, store)

		testId := "testing-123"

		account := makeAccount()
		account.Id = testId

		store.InsertAccount(account)

		err := store.DeleteAccount(testId)
		assert.NoError(t, err)

		fetchedAccount, _ := store.GetAccount(testId)
		assert.Nil(t, fetchedAccount)
	})
}

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
}

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

	t.Run("it gets categories", func(t *testing.T) {
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
