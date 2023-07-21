package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func TestAccountStore(t *testing.T) {
	it := NewStoreIntegrationTest()
	account := testAccount()
	timeNow := int64(12345)

	t.Run("it inserts and gets", func(t *testing.T) {
		err := it.AccountStore.Insert(account)
		assert.NoError(t, err)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.True(t, found.NotEmpty())
		assert.Equal(t, account, found.Get())
	})

	t.Run("it updates", func(t *testing.T) {
		newEmail := "some-new-email@gmail.com"
		update := data.AccountUpdate{Email: newEmail}

		ok, err := it.AccountStore.Update(account.Id, update, timeNow)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.True(t, found.NotEmpty())
		assert.Equal(t, found.Get().Email, newEmail)
	})

	t.Run("it marks as deleted", func(t *testing.T) {
		ok, err := it.AccountStore.SetDeleted(account.Id)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.True(t, found.NotEmpty())
		assert.Equal(t, found.Get().IsDeleted, true)
	})

	t.Run("it deletes", func(t *testing.T) {
		ok, err := it.AccountStore.Delete(account.Id)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.True(t, found.Empty())
	})
}
