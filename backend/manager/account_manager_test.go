package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

func TestAccountManagerTestSuite(t *testing.T) {
	suite.Run(t, new(AccountManagerTestSuite))
}

type AccountManagerTestSuite struct {
	ManagerTestSuite
	manager *AccountManager
}

func (s *AccountManagerTestSuite) SetupSuite() {
	s.initMocks()
	s.manager = &AccountManager{
		store: s.accountStore,
		clock: s.clock,
	}
}

func (s *AccountManagerTestSuite) TestGets() {
	expected := &data.Account{Id: "account-123"}

	s.accountStore.EXPECT().Get(expected.Id).Return(expected, nil)

	actual, err := s.manager.Get(expected.Id)
	s.Equal(*expected, *actual)
	s.NoError(err)
}

func (s *AccountManagerTestSuite) TestUpdates() {
	expected := true
	account := &data.Account{Id: "account-123"}
	body := rest.AccountUpdateBody{
		Name: "John Doe",
	}

	updatedAccount := data.Account{
		Id:          account.Id,
		Name:        body.Name,
		Email:       account.Email,
		Password:    account.Password,
		IsActivated: account.IsActivated,
		IsDeleted:   account.IsDeleted,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   testTimestamp,
	}

	s.clock.EXPECT().Now().Return(testTimestamp)
	s.accountStore.EXPECT().Update(updatedAccount).Return(expected, nil)

	actual, err := s.manager.Update(account, body)
	s.Equal(expected, actual)
	s.NoError(err)
	s.Equal(*account, updatedAccount)
}

func (s *AccountManagerTestSuite) TestSoftDeletes() {
	expected := true
	accountId := "some-id"

	s.accountStore.EXPECT().SoftDelete(accountId).Return(expected, nil)

	actual, err := s.manager.SoftDelete(accountId)

	s.Equal(expected, actual)
	s.NoError(err)
}
