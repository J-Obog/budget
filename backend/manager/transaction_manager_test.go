package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/types"
	"github.com/stretchr/testify/suite"
)

func TestTransactionManager(t *testing.T) {
	suite.Run(t, new(TransactionManagerTestSuite))
}

type TransactionManagerTestSuite struct {
	ManagerTestSuite
	manager *TransactionManager
}

func (s *TransactionManagerTestSuite) SetupSuite() {
	s.initMocks()
	s.manager = &TransactionManager{
		store:        s.transactionStore,
		clock:        s.clock,
		uuidProvider: s.uuidProvider,
	}
}

func (s *TransactionManagerTestSuite) TestGet() {
	expected := &data.Transaction{
		Id:        "transaction-123",
		AccountId: "account-456",
	}

	s.transactionStore.EXPECT().Get(
		expected.Id,
		expected.AccountId,
	).Return(expected, nil)

	actual, err := s.manager.Get(expected.Id, expected.AccountId)
	s.Equal(*expected, *actual)
	s.NoError(err)
}

func (s *TransactionManagerTestSuite) TestCreate() {
	accountId := "account-123"
	body := rest.TransactionCreateBody{
		CategoryId: types.StringPtr("cat-123"),
		Note:       types.StringPtr("some note"),
		Type:       data.BudgetType_Expense,
		Amount:     123.45,
		Month:      8,
		Day:        5,
		Year:       2022,
	}

	expected := data.Transaction{
		Id:         testUuid,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
		CreatedAt:  testTimestamp,
		UpdatedAt:  testTimestamp,
	}

	s.clock.EXPECT().Now().Return(testTimestamp)
	s.uuidProvider.EXPECT().GetUuid().Return(testUuid)
	s.transactionStore.EXPECT().Insert(expected).Return(nil)

	actual, err := s.manager.Create(accountId, body)
	s.Equal(actual, expected)
	s.NoError(err)
}

func (s *TransactionManagerTestSuite) TestUpdate() {
	expected := true
	transaction := &data.Transaction{Id: "txn-123"}
	body := rest.TransactionUpdateBody{
		CategoryId: types.StringPtr("cat-123"),
		Note:       types.StringPtr("some note"),
		Type:       data.BudgetType_Expense,
		Amount:     123.45,
		Month:      8,
		Day:        5,
		Year:       2022,
	}

	updatedTransaction := data.Transaction{
		Id:         transaction.Id,
		AccountId:  transaction.AccountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  testTimestamp,
	}

	s.clock.EXPECT().Now().Return(testTimestamp)
	s.transactionStore.EXPECT().Update(updatedTransaction).Return(expected, nil)

	actual, err := s.manager.Update(transaction, body)
	s.Equal(expected, actual)
	s.NoError(err)
	s.Equal(*transaction, updatedTransaction)
}

func (s *TransactionManagerTestSuite) TestDelete() {
	expected := true
	id := "some-id"
	accountId := "some-other-id"

	s.transactionStore.EXPECT().Delete(id, accountId).Return(expected, nil)

	actual, err := s.manager.Delete(id, accountId)

	s.Equal(expected, actual)
	s.NoError(err)
}

func (s *TransactionManagerTestSuite) TestTotalForPeriodCategory() {
	categoryId := "some-id"
	accountId := "some-other-id"
	month := 10
	year := 2025

	transactions := []data.Transaction{
		{Type: data.BudgetType_Expense, Amount: 95.99},
		{Type: data.BudgetType_Income, Amount: 35.12},
		{Type: data.BudgetType_Income, Amount: 45.89},
	}

	expected := -14.98

	s.transactionStore.EXPECT().GetByPeriodCategory(
		accountId,
		categoryId,
		month,
		year,
	).Return(transactions, nil)

	actual, err := s.manager.GetTotalForPeriodCategory(
		accountId,
		categoryId,
		month,
		year,
	)

	s.Equal(expected, actual)
	s.NoError(err)
}
