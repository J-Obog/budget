package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/types"
	"github.com/stretchr/testify/suite"
)

type TransactionManagerTestSuite struct {
	suite.Suite
	store   *mocks.TransactionStore
	clock   *mocks.Clock
	uid     *mocks.UIDProvider
	manager *TransactionManager
}

func (s *TransactionManagerTestSuite) SetupSuite() {
	s.store = new(mocks.TransactionStore)
	s.clock = new(mocks.Clock)
	s.uid = new(mocks.UIDProvider)

	s.manager = &TransactionManager{
		store: s.store,
		clock: s.clock,
		uid:   s.uid,
	}
}

func (s *TransactionManagerTestSuite) TestGetsTransaction() {
	expected := &data.Transaction{
		Id:        "some-transaction-id",
		AccountId: "some-account",
	}

	s.store.On("Get", expected.Id, expected.AccountId).Return(expected, nil)

	actual, err := s.manager.Get(expected.Id, expected.AccountId)

	s.NoError(err)
	s.Equal(*expected, *actual)
}

func (s *TransactionManagerTestSuite) TestGetsTransactionByPeriodCategory() {
	accountId := "some-account"
	categoryId := "category-id"
	month := 10
	year := 2023

	expected := []data.Transaction{
		{Id: "some-transaction-id"},
	}

	s.store.On("GetByPeriodCategory",
		accountId,
		categoryId,
		month,
		year,
	).Return(expected, nil)

	actual, err := s.manager.GetByPeriodCategory(
		accountId,
		categoryId,
		month,
		year,
	)

	s.NoError(err)
	s.ElementsMatch(expected, actual)
}

func (s *TransactionManagerTestSuite) TestCreatesTransaction() {
	accountId := "account-id"

	body := rest.TransactionCreateBody{
		CategoryId: types.StringPtr("category-id"),
		Note:       types.StringPtr("some note content"),
		Type:       data.BudgetType_Expense,
		Amount:     10.89,
		Month:      4,
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

	s.clock.On("Now").Return(testTimestamp)
	s.uid.On("GetId").Return(testUuid)
	s.store.On("Insert", expected).Return(nil)

	actual, err := s.manager.Create(accountId, body)
	s.NoError(err)
	s.Equal(expected, actual)
}

func (s *TransactionManagerTestSuite) TestUpdatesTransaction() {
	existing := &data.Transaction{
		Id:        "transaction-id",
		AccountId: "account-id",
	}

	body := rest.TransactionUpdateBody{
		CategoryId: types.StringPtr("category-id"),
		Note:       types.StringPtr("some note content"),
		Type:       data.BudgetType_Expense,
		Amount:     10.89,
		Month:      4,
		Day:        5,
		Year:       2022,
	}

	update := data.TransactionUpdate{
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
	}

	s.clock.On("Now").Return(testTimestamp, nil)
	s.store.On("Update", existing.Id, existing.AccountId, update, testTimestamp).Return(true, nil)

	ok, err := s.manager.Update(existing, body)

	s.NoError(err)
	s.True(ok)
	s.Equal(existing, &data.Transaction{
		Id:         existing.Id,
		AccountId:  existing.AccountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
		UpdatedAt:  testTimestamp,
	})
}

func (s *TransactionManagerTestSuite) TestDeletesTransaction() {
	id := "some-transaction-to-delete"
	account := "some-account-id"

	s.store.On("Delete", id, account).Return(true, nil)

	ok, err := s.manager.Delete(id, account)

	s.NoError(err)
	s.True(ok)
}
