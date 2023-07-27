package store

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/suite"
)

type AccountStoreTestSuite struct {
	suite.Suite
	store AccountStore
}

func (s *AccountStoreTestSuite) SetupSuite() {
	cfg := config.Get()
	svc := GetConfiguredStoreService(cfg)
	s.store = svc.AccountStore
}

func (s *AccountStoreTestSuite) SetupTest() {
	err := s.store.DeleteAll()
	s.NoError(err)
}

func (s *AccountStoreTestSuite) TestInsertsAndGets() {
	account := data.Account{
		Id:        "account-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.store.Insert(account)
	s.NoError(err)

	found, err := s.store.Get(account.Id)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(account, *found)
}

func (s *AccountStoreTestSuite) TestUpdates() {
	account := data.Account{Id: "account-id"}
	update := data.AccountUpdate{Name: "New Account Name"}

	err := s.store.Insert(account)
	s.NoError(err)

	ok, err := s.store.Update(account.Id, update, testTimestamp)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(account.Id)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(found.Name, update.Name)
	s.Equal(found.UpdatedAt, testTimestamp)
}

func (s *AccountStoreTestSuite) TestMarksAccountAsDeleted() {
	account := data.Account{Id: "account-id", IsDeleted: false}

	err := s.store.Insert(account)
	s.NoError(err)

	ok, err := s.store.SetDeleted(account.Id)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(account.Id)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(found.IsDeleted, true)
}

func (s *AccountStoreTestSuite) TestDeletesAccount() {
	account := data.Account{Id: "account-id"}

	err := s.store.Insert(account)
	s.NoError(err)

	ok, err := s.store.Delete(account.Id)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(account.Id)
	s.NoError(err)
	s.Nil(found)
}
