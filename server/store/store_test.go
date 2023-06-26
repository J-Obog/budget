package store

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

const (
	testAccountId = "test-account-1"
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

func TestAccounts(t *testing.T) {
	store := getStore(t)

	t.Run("inserts and gets an account", func(t *testing.T) {
		setup(t, store)

		account := makeAccount()
		account.Id = testAccountId

		err := store.InsertAccount(account)
		assert.NoError(t, err)

		fetchedAccount, err := store.GetAccount(testAccountId)
		assert.NoError(t, err)
		assert.Equal(t, account, *fetchedAccount)
	})

	t.Run("updates an account", func(t *testing.T) {
		setup(t, store)

		oldEmail := "jdoe@gmail.com"
		newEmail := "jdoe@yahoo.com"

		account := makeAccount()
		account.Id = testAccountId
		account.Email = oldEmail

		err := store.InsertAccount(account)
		assert.NoError(t, err)

		account.Email = newEmail

		err = store.UpdateAccount(account)
		assert.NoError(t, err)

		fetchedAccount, err := store.GetAccount(testAccountId)
		assert.NoError(t, err)
		assert.Equal(t, account, *fetchedAccount)
	})

	t.Run("deletes an account", func(t *testing.T) {
		setup(t, store)

		account := makeAccount()
		account.Id = testAccountId

		err := store.InsertAccount(account)
		assert.NoError(t, err)

		err = store.DeleteAccount(testAccountId)
		assert.NoError(t, err)

		fetchedAccount, err := store.GetAccount(testAccountId)
		assert.NoError(t, err)

		assert.Nil(t, fetchedAccount)
	})
}
