package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountStore(t *testing.T) {
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
