package store

import (
	"fmt"
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/suite"
)

func TestCategoryStore(t *testing.T) {
	suite.Run(t, new(CategoryStoreTestSuite))
}

type CategoryStoreTestSuite struct {
	StoreTestSuite
}

func (s *CategoryStoreTestSuite) SetupTest() {
	err := s.categoryStore.DeleteAll()
	s.NoError(err)
}

func (s *CategoryStoreTestSuite) TestInsertAndGet() {
	category := data.Category{
		Id:        "category-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.categoryStore.Insert(category)
	s.NoError(err)

	actual, err := s.categoryStore.Get(category.Id, category.AccountId)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(category, *actual)
}

func (s *CategoryStoreTestSuite) TestGetByName() {
	category := data.Category{
		Id:        "category-id",
		AccountId: "account-id",
		Name:      "FooBar",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.categoryStore.Insert(category)
	s.NoError(err)

	actual, err := s.categoryStore.GetByName(category.AccountId, category.Name)
	s.NoError(err)
	s.NotNil(actual)
	s.Equal(category, *actual)
}

func (s *CategoryStoreTestSuite) TestGetAll() {
	accountId := "some-account-id"

	categories := []data.Category{}

	for i := 0; i < 5; i++ {
		category := data.Category{
			Id:        fmt.Sprintf("id-%d", i),
			AccountId: accountId,
			CreatedAt: testTimestamp,
			UpdatedAt: testTimestamp,
		}

		categories = append(categories, category)

		err := s.categoryStore.Insert(category)
		s.NoError(err)
	}

	actual, err := s.categoryStore.GetAll(accountId)
	s.NoError(err)
	s.ElementsMatch(actual, categories)
}

func (s *CategoryStoreTestSuite) TestUpdate() {
	category := data.Category{
		Id:        "category-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.categoryStore.Insert(category)
	s.NoError(err)

	category.Color = 12345

	ok, err := s.categoryStore.Update(category)
	s.NoError(err)
	s.True(ok)

	actual, err := s.categoryStore.Get(category.Id, category.AccountId)
	s.NoError(err)
	s.Equal(*actual, category)
}

func (s *CategoryStoreTestSuite) TestDelete() {
	category := data.Category{Id: "category-id"}

	err := s.categoryStore.Insert(category)
	s.NoError(err)

	ok, err := s.categoryStore.Delete(category.Id, category.AccountId)
	s.NoError(err)
	s.True(ok)

	actual, err := s.categoryStore.Get(category.Id, category.AccountId)
	s.NoError(err)
	s.Nil(actual)
}
