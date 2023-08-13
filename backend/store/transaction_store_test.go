package store

import (
	"fmt"
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/suite"
)

func TestTransactionStore(t *testing.T) {
	suite.Run(t, new(TransactionStoreTestSuite))
}

type TransactionStoreTestSuite struct {
	StoreTestSuite
}

func (s *TransactionStoreTestSuite) SetupTest() {
	err := s.transactionStore.DeleteAll()
	s.NoError(err)
}

func (s *TransactionStoreTestSuite) TestInsertsAndGetsTransaction() {
	transaction := data.Transaction{
		Id:        "transaction-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.transactionStore.Insert(transaction)
	s.NoError(err)

	actual, err := s.transactionStore.Get(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(transaction, *actual)
}

// TODO: implement
func (s *TransactionStoreTestSuite) TestGetsTransactionsByFilter() {
}

func (s *TransactionStoreTestSuite) TestGetByPeriodCategory() {
	accountId := "some-account-id"
	categoryId := "some-category-id"
	month := 10
	year := 2024

	transactions := []data.Transaction{}

	for i := 0; i < 5; i++ {
		transaction := data.Transaction{
			Id:         fmt.Sprintf("id-%d", i),
			AccountId:  accountId,
			CategoryId: &categoryId,
			Month:      month,
			Year:       year,
			CreatedAt:  testTimestamp,
			UpdatedAt:  testTimestamp,
		}

		transactions = append(transactions, transaction)

		err := s.transactionStore.Insert(transaction)
		s.NoError(err)
	}

	actual, err := s.transactionStore.GetByPeriodCategory(accountId, categoryId, month, year)
	s.NoError(err)
	s.ElementsMatch(actual, transactions)
}

func (s *TransactionStoreTestSuite) TestUpdate() {
	transaction := data.Transaction{
		Id:        "transaction-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.transactionStore.Insert(transaction)
	s.NoError(err)

	transaction.Amount = 123.45

	ok, err := s.transactionStore.Update(transaction)
	s.NoError(err)
	s.True(ok)

	actual, err := s.transactionStore.Get(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.Equal(*actual, transaction)
}

func (s *TransactionStoreTestSuite) TestDelete() {
	transaction := data.Transaction{Id: "transaction-id"}

	err := s.transactionStore.Insert(transaction)
	s.NoError(err)

	ok, err := s.transactionStore.Delete(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.True(ok)

	actual, err := s.transactionStore.Get(transaction.Id, transaction.AccountId)
	s.NoError(err)
	s.Nil(actual)
}
