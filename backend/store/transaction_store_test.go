package store

import (
	"fmt"
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
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

func (s *TransactionStoreTestSuite) TestGetByFilter() {
	accountId := "acct-123"
	minAmount := 123.45
	maxAmount := 678.90
	minDate := data.NewDate(1, 2, 2022)
	maxDate := data.NewDate(2, 3, 2023)

	filter := data.TransactionFilter{
		MinAmount: types.Float64Ptr(minAmount),
		MaxAmount: types.Float64Ptr(maxAmount),
		StartDate: types.Ptr[data.Date](minDate),
		EndDate:   types.Ptr[data.Date](maxDate),
	}

	transactions := []data.Transaction{
		{
			Id:        "txn-1",
			AccountId: accountId,
			Amount:    minAmount,
			Month:     minDate.Month,
			Day:       minDate.Day,
			Year:      minDate.Year,
			CreatedAt: testTimestamp,
			UpdatedAt: testTimestamp,
		},
		{
			Id:        "txn-2",
			AccountId: accountId,
			Amount:    maxAmount,
			Month:     maxDate.Month,
			Day:       maxDate.Day,
			Year:      maxDate.Year,
			CreatedAt: testTimestamp,
			UpdatedAt: testTimestamp,
		},
	}

	for _, transaction := range transactions {
		err := s.transactionStore.Insert(transaction)
		s.NoError(err)
	}

	actual, err := s.transactionStore.GetByFilter(accountId, filter)
	s.NoError(err)
	s.ElementsMatch(transactions, actual)
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
