package store

import (
	"fmt"

	"github.com/J-Obog/paidoff/data"
)

type BudgetStoreTestSuite struct {
	StoreTestSuite
}

func (s *BudgetStoreTestSuite) SetupTest() {
	err := s.budgetStore.DeleteAll()
	s.NoError(err)
}

func (s *BudgetStoreTestSuite) TestInsertAndGet() {
	budget := data.Budget{
		Id:        "budget-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.budgetStore.Insert(budget)
	s.NoError(err)

	actual, err := s.budgetStore.Get(budget.Id, budget.AccountId)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(budget, *actual)
}

func (s *BudgetStoreTestSuite) TestGetByPeriodCategory() {
	budget := data.Budget{
		Id:         "some-account-id",
		CategoryId: "some-category-id",
		Month:      10,
		Year:       2024,
		CreatedAt:  testTimestamp,
		UpdatedAt:  testTimestamp,
	}

	err := s.budgetStore.Insert(budget)
	s.NoError(err)

	actual, err := s.budgetStore.GetByPeriodCategory(
		budget.AccountId,
		budget.CategoryId,
		budget.Month,
		budget.Year,
	)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(budget, *actual)
}

func (s *BudgetStoreTestSuite) TestGetByCategory() {
	categoryId := "some-category-id"
	accountId := "some-account-id"

	budgets := []data.Budget{}

	for i := 0; i < 5; i++ {
		budget := data.Budget{
			Id:         fmt.Sprintf("id-%d", i),
			AccountId:  accountId,
			CategoryId: categoryId,
			CreatedAt:  testTimestamp,
			UpdatedAt:  testTimestamp,
		}

		budgets = append(budgets, budget)

		err := s.budgetStore.Insert(budget)
		s.NoError(err)
	}

	found, err := s.budgetStore.GetByCategory(accountId, categoryId)
	s.NoError(err)
	s.ElementsMatch(found, budgets)
}

// TODO: implement
func (s *BudgetStoreTestSuite) TestGetsBudgetsByFilter() {
}

func (s *BudgetStoreTestSuite) TestUpdate() {
	budget := data.Budget{Id: "budget-id"}

	err := s.budgetStore.Insert(budget)
	s.NoError(err)

	budget.Month = 10

	ok, err := s.budgetStore.Update(budget)
	s.NoError(err)
	s.True(ok)

	actual, err := s.budgetStore.Get(budget.Id, budget.AccountId)
	s.NoError(err)
	s.Equal(actual, budget)
}

func (s *BudgetStoreTestSuite) TestDelete() {
	budget := data.Budget{Id: "budget-id"}

	err := s.budgetStore.Insert(budget)
	s.NoError(err)

	ok, err := s.budgetStore.Delete(budget.Id, budget.AccountId)
	s.NoError(err)
	s.True(ok)

	actual, err := s.budgetStore.Get(budget.Id, budget.AccountId)
	s.NoError(err)
	s.Nil(actual)
}
