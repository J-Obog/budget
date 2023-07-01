package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountStore(t *testing.T) {
	it := dbIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		setup(it)

		account := testAccount()

		err := it.AccountStore.Insert(account)
		assert.NoError(t, err)

		fetched, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.Equal(t, account, *fetched)
	})

	t.Run("it updates", func(t *testing.T) {
		setup(it)

		account := testAccount()
		account.Email = "jdoe@gmail.com"

		it.AccountStore.Insert(account)

		account.Email = "jdoe@yahoo.com"

		err := it.AccountStore.Update(account)
		assert.NoError(t, err)

		fetched, _ := it.AccountStore.Get(account.Id)
		assert.Equal(t, account, *fetched)
	})

	t.Run("it deletes", func(t *testing.T) {
		setup(it)

		account := testAccount()

		it.AccountStore.Insert(account)

		err := it.AccountStore.Delete(account.Id)
		assert.NoError(t, err)

		fetched, _ := it.AccountStore.Get(account.Id)
		assert.Nil(t, fetched)
	})
}
