package store

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, store Store) {
	err := store.Flush()
	assert.NoError(t, err)
}

func getStore(t *testing.T) Store {
	cfg, err := config.MakeConfig("local")
	assert.NoError(t, err)

	s, err := MakeStore(cfg)
	assert.NoError(t, err)

	return s
}

func makeAccount() data.Account {
	return data.Account{
		Id:          "test-12345",
		Name:        "John Doe",
		Email:       "jdoe@gmail.com",
		Password:    "foobar",
		IsActivated: true,
		CreatedAt:   1234,
		UpdatedAt:   4567,
	}
}

func makeBudget() data.Budget {
	return data.Budget{
		Id:         "test-123",
		AccountId:  "test-456",
		CategoryId: nil,
		Name:       "eating",
		Type:       data.BudgetType_EXPENSE,
		Month:      8,
		Year:       2023,
		Projected:  115.60,
		Actual:     67.89,
		CreatedAt:  123,
		UpdatedAt:  456,
	}
}

func makeTransaction() data.Transaction {
	return data.Transaction{
		Id:          "t-12354",
		AccountId:   "testing-45678",
		BudgetId:    "testy-123",
		Description: nil,
		Amount:      590.20,
		Month:       9,
		Day:         5,
		Year:        2028,
		CreatedAt:   123,
		UpdatedAt:   456,
	}
}

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

		err := store.InsertAccount(account)
		assert.NoError(t, err)

		account.Email = newEmail

		err = store.UpdateAccount(account)
		assert.NoError(t, err)

		fetchedAccount, err := store.GetAccount(testId)
		assert.NoError(t, err)
		assert.Equal(t, account, *fetchedAccount)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(t, store)

		testId := "testing-123"

		account := makeAccount()
		account.Id = testId

		err := store.InsertAccount(account)
		assert.NoError(t, err)

		err = store.DeleteAccount(testId)
		assert.NoError(t, err)

		fetchedAccount, err := store.GetAccount(testId)
		assert.NoError(t, err)

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

		err := store.InsertBudget(budget)
		assert.NoError(t, err)

		budget.Projected = float64(newProjectedBudget)

		err = store.UpdateBudget(budget)
		assert.NoError(t, err)

		fetchedBudget, err := store.GetBudget(testId)
		assert.NoError(t, err)
		assert.Equal(t, budget, *fetchedBudget)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(t, store)

		testId := "t-1"

		budget := makeBudget()
		budget.Id = testId

		err := store.InsertBudget(budget)
		assert.NoError(t, err)

		err = store.DeleteBudget(testId)
		assert.NoError(t, err)

		fetchedBudget, err := store.GetBudget(testId)
		assert.NoError(t, err)
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

		err := store.InsertTransaction(transaction)
		assert.NoError(t, err)

		transaction.Amount = newAmount

		err = store.UpdateTransaction(transaction)
		assert.NoError(t, err)

		fetchedTransaction, err := store.GetTransaction(testId)
		assert.NoError(t, err)
		assert.Equal(t, transaction, *fetchedTransaction)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(t, store)

		testId := "t-1"

		transaction := makeTransaction()
		transaction.Id = testId

		err := store.InsertTransaction(transaction)
		assert.NoError(t, err)

		err = store.DeleteTransaction(testId)
		assert.NoError(t, err)

		fetchedTransaction, err := store.GetTransaction(testId)
		assert.NoError(t, err)
		assert.Nil(t, fetchedTransaction)
	})
}
