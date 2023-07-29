package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

type AccountManagerTestSuite struct {
	suite.Suite
	store   *mocks.AccountStore
	clock   *mocks.Clock
	manager *AccountManager
}

func (s *AccountManagerTestSuite) SetupSuite() {
	s.store = new(mocks.AccountStore)
	s.clock = new(mocks.Clock)

	s.manager = &AccountManager{
		store: s.store,
		clock: s.clock,
	}
}

func (s *AccountManagerTestSuite) TestUpdatesAccount() {
	existing := &data.Account{}

	body := rest.AccountUpdateBody{
		Name: "Some Name",
	}

	update := data.AccountUpdate{
		Name: body.Name,
	}

	s.clock.On("Now").Return(testTimestamp, nil)
	s.store.On("Update", existing.Id, update, testTimestamp).Return(true, nil)

	ok, err := s.manager.Update(existing, body)

	s.NoError(err)
	s.True(ok)
	s.Equal(existing, &data.Account{
		Name:      update.Name,
		UpdatedAt: testTimestamp,
	})
}

func (s *AccountManagerTestSuite) TestDeletesAccount() {
	id := "account-id"

	s.store.On("SetDeleted", id).Return(true, nil)

	ok, err := s.manager.Delete(id)
	s.NoError(err)
	s.True(ok)
}
