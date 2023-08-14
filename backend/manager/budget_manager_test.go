package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

func TestBudgetManagerTestSuite(t *testing.T) {
	suite.Run(t, new(BudgetManagerTestSuite))
}

type BudgetManagerTestSuite struct {
	ManagerTestSuite
	manager *BudgetManager
}

func (s *BudgetManagerTestSuite) SetupSuite() {
	s.initMocks()
	s.manager = &BudgetManager{
		store:        s.budgetStore,
		clock:        s.clock,
		uuidProvider: s.uuidProvider,
	}
}

func (s *BudgetManagerTestSuite) TestGets() {
	expected := &data.Budget{
		Id:        "budget-123",
		AccountId: "account-456",
	}

	s.budgetStore.EXPECT().Get(
		expected.Id,
		expected.AccountId,
	).Return(expected, nil)

	actual, err := s.manager.Get(expected.Id, expected.AccountId)
	s.Equal(*expected, *actual)
	s.NoError(err)
}

func (s *BudgetManagerTestSuite) TestGetsByPeriod() {
	accountId := "acc-123"
	month := 10
	year := 2023

	expected := []data.Budget{
		{Id: "some-uuid"},
	}

	s.budgetStore.EXPECT().GetByPeriod(
		accountId,
		month,
		year,
	).Return(expected, nil)

	actual, err := s.manager.GetByPeriod(accountId, month, year)
	s.ElementsMatch(expected, actual)
	s.NoError(err)
}

func (s *BudgetManagerTestSuite) TestCreates() {
	accountId := "account-123"
	body := rest.BudgetCreateBody{
		CategoryId: "category-123",
		Projected:  10.45,
		Month:      10,
		Year:       2023,
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

	s.clock.EXPECT().Now().Return(testTimestamp)
	s.uuidProvider.EXPECT().GetUuid().Return(testUuid)
	s.budgetStore.EXPECT().Insert(expected).Return(nil)

	actual, err := s.manager.Create(accountId, body)
	s.Equal(actual, expected)
	s.NoError(err)
}

func (s *BudgetManagerTestSuite) TestUpdates() {
	expected := true
	budget := &data.Budget{Id: "budget-123"}
	body := rest.BudgetUpdateBody{
		CategoryId: "category-4567",
		Projected:  23.78,
	}

	updatedBudget := data.Budget{
		Id:         budget.Id,
		AccountId:  budget.AccountId,
		CategoryId: body.CategoryId,
		Month:      budget.Month,
		Year:       budget.Year,
		Projected:  body.Projected,
		CreatedAt:  budget.CreatedAt,
		UpdatedAt:  testTimestamp,
	}

	s.clock.EXPECT().Now().Return(testTimestamp)
	s.budgetStore.EXPECT().Update(updatedBudget).Return(expected, nil)

	actual, err := s.manager.Update(budget, body)
	s.Equal(expected, actual)
	s.NoError(err)
	s.Equal(*budget, updatedBudget)
}

func (s *BudgetManagerTestSuite) TestDeletes() {
	expected := true
	id := "some-id"
	accountId := "some-other-id"

	s.budgetStore.EXPECT().Delete(id, accountId).Return(expected, nil)

	actual, err := s.manager.Delete(id, accountId)

	s.Equal(expected, actual)
	s.NoError(err)
}

func (s *BudgetManagerTestSuite) TestCategoryIsNotUsed() {
	expected := false
	categoryId := "some-id"
	accountId := "some-other-id"

	s.budgetStore.EXPECT().GetByCategory(
		accountId,
		categoryId,
	).Return([]data.Budget{{Id: "some-budget"}}, nil)

	actual, err := s.manager.CategoryIsNotUsed(
		categoryId,
		accountId,
	)

	s.Equal(expected, actual)
	s.NoError(err)
}

func (s *BudgetManagerTestSuite) TestCategoryIsUniqueForPeriod() {
	expected := false
	categoryId := "some-id"
	accountId := "some-other-id"
	month := 8
	year := 2023

	s.budgetStore.EXPECT().GetByPeriodCategory(
		accountId,
		categoryId,
		month,
		year,
	).Return(&data.Budget{Id: "some-budget"}, nil)

	actual, err := s.manager.CategoryIsUniqueForPeriod(
		categoryId,
		accountId,
		month,
		year,
	)

	s.Equal(expected, actual)
	s.NoError(err)
}
