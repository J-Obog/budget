package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/assert"
)

func TestAccountManagerGetsByRequest(t *testing.T) {
	manager := accountManagerMock()
	account := testAccount()

	req := &rest.Request{
		Account: account,
	}

	res := manager.GetByRequest(req)

	assert.Equal(t, res.Data, account)
	assert.NoError(t, res.Error)
}

func TestAccountManagerUpdatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := accountManagerMock()
		timeNow := int64(1234555)

		req := accountUpdate()

		update := data.AccountUpdate{
			Name: req.Body.(rest.AccountUpdateBody).Name,
		}

		manager.MockClock.On("Now").Return(timeNow)
		manager.MockAccountStore.On("Update", req.Account.Id, update, timeNow).Return(true, nil)

		res := manager.UpdateByRequest(req)

		assert.NoError(t, res.Error)
	})

	t.Run("it fails when account name is too short", func(t *testing.T) {
		manager := accountManagerMock()
		req := accountUpdateNameTooShort()
		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidAccountName)
	})

	t.Run("it fails when account name is too long", func(t *testing.T) {
		manager := accountManagerMock()
		req := accountUpdateNameTooLong()
		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidAccountName)
	})

}

func TestAccountManagerDeletesByRequest(t *testing.T) {
	manager := accountManagerMock()

	req := accountDelete()

	manager.MockAccountStore.On("SetDeleted", req.Account.Id).Return(true, nil)

	res := manager.DeleteByRequest(req)

	assert.NoError(t, res.Error)
}
