package store

import (
	"fmt"
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/suite"
)

type BudgetStoreTestSuite struct {
	suite.Suite
	store BudgetStore
}

func TestBudgetStore(t *testing.T) {
	suite.Run(t, new(AccountStoreTestSuite))
}

func (s *BudgetStoreTestSuite) SetupSuite() {
	cfg := config.Get()
	svc := GetConfiguredStoreService(cfg)
	s.store = svc.BudgetStore
}

func (s *BudgetStoreTestSuite) SetupTest() {
	err := s.store.DeleteAll()
	s.NoError(err)
}

func (s *BudgetStoreTestSuite) TestInsertsAndGets() {
	budget := data.Budget{Id: "budget-id"}

	err := s.store.Insert(budget)
	s.NoError(err)

	found, err := s.store.Get(budget.Id, budget.AccountId)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(budget, *found)
}

func (s *BudgetStoreTestSuite) TestGetsByPeriodCategory() {
	budget := data.Budget{
		Id:         "some-account-id",
		CategoryId: "some-category-id",
		Month:      10,
		Year:       2024,
	}

	err := s.store.Insert(budget)
	s.NoError(err)

	found, err := s.store.GetByPeriodCategory(budget.AccountId, budget.CategoryId, budget.Month, budget.Year)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(budget, *found)
}

func (s *BudgetStoreTestSuite) TestGetsByCategory() {
	categoryId := "some-category-id"
	accountId := "some-account-id"

	expected := []data.Budget{}

	for i := 0; i < 5; i++ {
		budget := data.Budget{Id: fmt.Sprintf("id-%d", i), AccountId: accountId, CategoryId: categoryId}
		err := s.store.Insert(budget)
		s.NoError(err)
	}

	found, err := s.store.GetByCategory(accountId, categoryId)
	s.NoError(err)
	s.ElementsMatch(found, expected)
}

// TODO: implement
func (s *BudgetStoreTestSuite) TestGetsByFilter() {
}

func (s *BudgetStoreTestSuite) TestUpdates() {
	budget := data.Budget{Id: "budget-id"}

	err := s.store.Insert(budget)
	s.NoError(err)

	update := data.BudgetUpdate{
		CategoryId: "some-category-id",
		Projected:  10.56,
	}

	ok, err := s.store.Update(budget.Id, budget.AccountId, update, 12345)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(budget.Id, budget.AccountId)
	s.NoError(err)
	s.Equal(found.CategoryId, update.CategoryId)
	s.Equal(found.Projected, update.Projected)
	s.Equal(found.UpdatedAt, testTimestamp)
}

func (s *BudgetStoreTestSuite) TestDeletes() {
	budget := data.Budget{Id: "budget-id"}

	err := s.store.Insert(budget)
	s.NoError(err)

	ok, err := s.store.Delete(budget.Id, budget.AccountId)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(budget.Id, budget.AccountId)
	s.NoError(err)
	s.Nil(found)
}
