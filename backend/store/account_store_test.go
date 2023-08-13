package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/suite"
)

func TestAccountStore(t *testing.T) {
	suite.Run(t, new(AccountStoreTestSuite))
}

type AccountStoreTestSuite struct {
	StoreTestSuite
}

func (s *AccountStoreTestSuite) SetupTest() {
	err := s.accountStore.DeleteAll()
	s.NoError(err)
}

func (s *AccountStoreTestSuite) TestInsertAndGet() {
	account := data.Account{
		Id:        "account-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.accountStore.Insert(account)
	s.NoError(err)

	actual, err := s.accountStore.Get(account.Id)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(account, *actual)
}

func (s *AccountStoreTestSuite) TestUpdate() {
	account := data.Account{
		Id:        "account-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.accountStore.Insert(account)
	s.NoError(err)

	account.Email = "some-email"

	ok, err := s.accountStore.Update(account)
	s.NoError(err)
	s.True(ok)

	actual, err := s.accountStore.Get(account.Id)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(account, *actual)
}

func (s *AccountStoreTestSuite) TestSoftDelete() {
	account := data.Account{
		Id:        "account-id",
		IsDeleted: false,
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.accountStore.Insert(account)
	s.NoError(err)

	ok, err := s.accountStore.SoftDelete(account.Id)
	s.NoError(err)
	s.True(ok)

	actual, err := s.accountStore.Get(account.Id)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(actual.IsDeleted, true)
}

func (s *AccountStoreTestSuite) TestDelete() {
	account := data.Account{Id: "account-id"}

	err := s.accountStore.Insert(account)
	s.NoError(err)

	ok, err := s.accountStore.Delete(account.Id)
	s.NoError(err)
	s.True(ok)

	actual, err := s.accountStore.Get(account.Id)
	s.NoError(err)
	s.Nil(actual)
}
