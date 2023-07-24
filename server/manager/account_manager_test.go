package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/assert"
)

func TestAccountManagerGetsByRequest(t *testing.T) {
	manager := accountManagerMock()
	req := testRequest()

	res := manager.GetByRequest(req)
	assert.Equal(t, res.Data, req.Account)
	assert.NoError(t, res.Error)
}

func TestAccountManagerUpdatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := accountManagerMock()
		body := rest.AccountUpdateBody{Name: "Some Name"}
		req := testRequest()
		req.Body = body

		update := getExpectedAccountUpdate(body)

		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockAccountStore.On("Update", req.Account.Id, update, testTimestamp).Return(true, nil)

		res := manager.UpdateByRequest(req)

		assert.NoError(t, res.Error)
	})

	t.Run("it fails when account name is too short", func(t *testing.T) {
		manager := accountManagerMock()
		body := rest.AccountUpdateBody{Name: genString(config.LimitMinAccountNameChars - 5)}
		req := testRequest()
		req.Body = body
		req.Account.Name = "Some old name"

		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidAccountName)
	})

	t.Run("it fails when account name is too long", func(t *testing.T) {
		manager := accountManagerMock()
		body := rest.AccountUpdateBody{Name: genString(config.LimitMaxAccountNameChars + 5)}
		req := testRequest()
		req.Body = body
		req.Account.Name = "Some old name"

		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidAccountName)
	})

}

func TestAccountManagerDeletesByRequest(t *testing.T) {
	manager := accountManagerMock()
	req := testRequest()

	manager.MockAccountStore.On("SetDeleted", req.Account.Id).Return(true, nil)

	res := manager.DeleteByRequest(req)

	assert.NoError(t, res.Error)
}

func getExpectedAccountUpdate(body rest.AccountUpdateBody) data.AccountUpdate {
	return data.AccountUpdate{
		Name: body.Name,
	}
}
