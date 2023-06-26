package store

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

const (
	testAccountId = "account-1"
)

func setupDbIntegrationTest() Store {
	cfg, _ := config.MakeConfig("local")
	s, _ := MakeStore(cfg)
	s.Flush()
	return s
}

func TestAccounts(t *testing.T) {
	store := setupDbIntegrationTest()
	t.Run("inserts and gets an account", func(t *testing.T) {
		account := data.Account{
			Id:          testAccountId,
			Name:        "John Doe",
			Email:       "jdoe@gmail.com",
			Password:    "foobar",
			IsActivated: true,
			CreatedAt:   1234,
			UpdatedAt:   4567,
		}

		err := store.InsertAccount(account)
		assert.NoError(t, err)

		savedAccount, err := store.GetAccount(testAccountId)

		assert.NoError(t, err)

		assert.Equal(t, account, *savedAccount)
	})

	t.Run("updates an account", func(t *testing.T) {

	})

	t.Run("deletes an account", func(t *testing.T) {

	})

}
