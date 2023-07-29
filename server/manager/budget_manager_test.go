package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/types"
	"github.com/stretchr/testify/suite"
)

type BudgetManagerTestSuite struct {
	suite.Suite
	store   *mocks.BudgetStore
	clock   *mocks.Clock
	uid     *mocks.UIDProvider
	manager *BudgetManager
}

func (s *BudgetManagerTestSuite) SetupSuite() {
	s.store = new(mocks.BudgetStore)
	s.clock = new(mocks.Clock)
	s.uid = new(mocks.UIDProvider)

	s.manager = &BudgetManager{
		store: s.store,
		clock: s.clock,
		uid:   s.uid,
	}
}

func (s *BudgetManagerTestSuite) TestGetsBudget() {
	expected := &data.Budget{
		Id:        "some-budget-id",
		AccountId: "some-account",
	}

	s.store.On("Get", expected.Id, expected.AccountId).Return(expected, nil)

	actual, err := s.manager.Get(expected.Id, expected.AccountId)

	s.NoError(err)
	s.Equal(*expected, *actual)
}

func (s *BudgetManagerTestSuite) TestGetsByQuery() {
	accountId := "some-account"

	expected := []data.Budget{
		{Id: "some-budget"},
	}

	defaultMonth := 10
	defaultYear := 2024

	s.clock.On("CurrentMonth").Return(defaultMonth)
	s.clock.On("CurrentYear").Return(defaultYear)

	s.Run("sets defaults", func() {
		q := rest.BudgetQuery{}
		expectedFilter := data.BudgetFilter{
			Month: defaultMonth,
			Year:  defaultYear,
		}

		s.store.On("GetBy", accountId, expectedFilter).Return(expected, nil)

		actual, err := s.manager.GetByQuery(accountId, q)
		s.NoError(err)
		s.ElementsMatch(expected, actual)
	})

	s.Run("uses query values", func() {
		q := rest.BudgetQuery{
			Month: types.IntPtr(11),
			Year:  types.IntPtr(4),
		}
		expectedFilter := data.BudgetFilter{
			Month: *q.Month,
			Year:  *q.Year,
		}

		s.store.On("GetBy", accountId, expectedFilter).Return(expected, nil)

		actual, err := s.manager.GetByQuery(accountId, q)
		s.NoError(err)
		s.ElementsMatch(expected, actual)
	})
}

func (s *BudgetManagerTestSuite) TestGetsBudgetsByPeriodCategory() {
	expected := &data.Budget{
		AccountId:  "some-account",
		CategoryId: "some-category-id",
		Month:      10,
		Year:       2023,
	}

	s.store.On("GetByPeriodCategory",
		expected.AccountId,
		expected.CategoryId,
		expected.Month,
		expected.Year,
	).Return(expected, nil)

	actual, err := s.manager.GetByPeriodCategory(
		expected.AccountId,
		expected.CategoryId,
		expected.Month,
		expected.Year,
	)

	s.NoError(err)
	s.Equal(*expected, *actual)
}

func (s *BudgetManagerTestSuite) TestGetsBudgetsByCategory() {
	accountId := "account-id"
	categoryId := "category-id"

	expected := []data.Budget{
		{AccountId: accountId, CategoryId: categoryId},
	}

	s.store.On("GetByCategory", accountId, categoryId).Return(expected, nil)

	actual, err := s.manager.GetByCategory(accountId, categoryId)
	s.NoError(err)
	s.ElementsMatch(expected, actual)
}

func (s *BudgetManagerTestSuite) TestCreatesBudget() {
	accountId := "account-id"

	body := rest.BudgetCreateBody{
		CategoryId: "category-id",
		Month:      15,
		Year:       2023,
		Projected:  10.57,
	}

	expected := data.Budget{
		Id:         testUuid,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Month:      body.Month,
		Year:       body.Year,
		Projected:  body.Projected,
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

func (s *BudgetManagerTestSuite) TestUpdatesBudget() {
	existing := &data.Budget{
		Id:        "budget-id",
		AccountId: "account-id",
	}

	body := rest.BudgetUpdateBody{
		CategoryId: "some-category-id",
		Projected:  500.87,
	}

	update := data.BudgetUpdate{
		CategoryId: body.CategoryId,
		Projected:  body.Projected,
	}

	s.clock.On("Now").Return(testTimestamp, nil)
	s.store.On("Update", existing.Id, existing.AccountId, update, testTimestamp).Return(true, nil)

	ok, err := s.manager.Update(existing, body)

	s.NoError(err)
	s.True(ok)
	s.Equal(existing, &data.Budget{
		Id:         existing.Id,
		AccountId:  existing.AccountId,
		CategoryId: body.CategoryId,
		Projected:  body.Projected,
		UpdatedAt:  testTimestamp,
	})
}

func (s *BudgetManagerTestSuite) TestDeletesBudget() {
	id := "some-budget-to-delete"
	account := "some-account-id"

	s.store.On("Delete", id, account).Return(true, nil)

	ok, err := s.manager.Delete(id, account)

	s.NoError(err)
	s.True(ok)
}
