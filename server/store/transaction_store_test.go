package store

import (
	"fmt"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
	"github.com/stretchr/testify/suite"
)

type TransactionStoreTestSuite struct {
	suite.Suite
	store TransactionStore
}

func (s *TransactionStoreTestSuite) SetupSuite() {
	cfg := config.Get()
	svc := GetConfiguredStoreService(cfg)
	s.store = svc.TransactionStore
}

func (s *TransactionStoreTestSuite) SetupTest() {
	err := s.store.DeleteAll()
	s.NoError(err)
}

func (s *TransactionStoreTestSuite) TestInsertsAndGets() {
	transaction := data.Transaction{Id: "transaction-id"}

	err := s.store.Insert(transaction)
	s.NoError(err)

	found, err := s.store.Get(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(transaction, *found)
}

// TODO: implement
func (s *TransactionStoreTestSuite) TestGetsFilter() {
}

func (s *TransactionStoreTestSuite) TestGetByPeriodCategory() {
	accountId := "some-account-id"
	categoryId := "some-category-id"
	month := 10
	year := 2024

	expected := []data.Transaction{}

	for i := 0; i < 5; i++ {
		transaction := data.Transaction{
			Id:         fmt.Sprintf("id-%d", i),
			AccountId:  accountId,
			CategoryId: &categoryId,
			Month:      month,
			Year:       year,
		}

		err := s.store.Insert(transaction)
		s.NoError(err)
	}

	found, err := s.store.GetByPeriodCategory(accountId, categoryId, month, year)
	s.NoError(err)
	s.ElementsMatch(found, expected)
}

func (s *TransactionStoreTestSuite) TestUpdates() {
	transaction := data.Transaction{Id: "transaction-id"}

	err := s.store.Insert(transaction)
	s.NoError(err)

	update := data.TransactionUpdate{
		CategoryId: types.StringPtr("category-id"),
		Note:       types.StringPtr("Some note"),
		Type:       data.BudgetType_Income,
		Amount:     123.45,
		Month:      11,
		Day:        7,
		Year:       2023,
	}

	ok, err := s.store.Update(transaction.Id, transaction.AccountId, update, testTimestamp)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.Equal(*found.CategoryId, *update.CategoryId)
	s.Equal(*found.Note, *update.Note)
	s.Equal(found.Type, update.Type)
	s.Equal(found.Amount, update.Amount)
	s.Equal(found.Month, update.Month)
	s.Equal(found.Day, update.Day)
	s.Equal(found.Year, update.Year)
	s.Equal(*found.CategoryId, *update.CategoryId)

	s.Equal(found.UpdatedAt, testTimestamp)
}

func (s *TransactionStoreTestSuite) TestDeletes() {
	transaction := data.Transaction{Id: "transaction-id"}

	err := s.store.Insert(transaction)
	s.NoError(err)

	ok, err := s.store.Delete(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.Nil(found)
}
