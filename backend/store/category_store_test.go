package store

import (
	"fmt"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/suite"
)

type CategoryStoreTestSuite struct {
	suite.Suite
	store CategoryStore
}

func (s *CategoryStoreTestSuite) SetupSuite() {
	cfg := config.Get()
	svc := GetConfiguredStoreService(cfg)
	s.store = svc.CategoryStore
}

func (s *CategoryStoreTestSuite) SetupTest() {
	err := s.store.DeleteAll()
	s.NoError(err)
}

func (s *CategoryStoreTestSuite) TestInsertsAndGetsCategory() {
	category := data.Category{
		Id:        "category-id",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.store.Insert(category)
	s.NoError(err)

	found, err := s.store.Get(category.Id, category.AccountId)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(category, *found)
}

func (s *CategoryStoreTestSuite) TestGetsCategoryByName() {
	category := data.Category{
		Id:        "category-id",
		AccountId: "account-id",
		Name:      "FooBar",
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	err := s.store.Insert(category)
	s.NoError(err)

	found, err := s.store.GetByName(category.AccountId, category.Name)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(category, *found)
}

func (s *CategoryStoreTestSuite) TestGetsAllCategories() {
	accountId := "some-account-id"

	expected := []data.Category{}

	for i := 0; i < 5; i++ {
		category := data.Category{
			Id:        fmt.Sprintf("id-%d", i),
			AccountId: accountId,
			CreatedAt: testTimestamp,
			UpdatedAt: testTimestamp,
		}

		expected = append(expected, category)

		err := s.store.Insert(category)
		s.NoError(err)
	}

	found, err := s.store.GetAll(accountId)
	s.NoError(err)
	s.ElementsMatch(found, expected)
}

func (s *CategoryStoreTestSuite) TestUpdatesCategory() {
	category := data.Category{Id: "category-id"}

	err := s.store.Insert(category)
	s.NoError(err)

	update := data.CategoryUpdate{
		Name:  "Baz",
		Color: 11111,
	}

	ok, err := s.store.Update(category.Id, category.AccountId, update, testTimestamp)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(category.Id, category.AccountId)
	s.NoError(err)
	s.Equal(found.Name, update.Name)
	s.Equal(found.Color, update.Color)
	s.Equal(found.UpdatedAt, testTimestamp)
}

func (s *CategoryStoreTestSuite) TestDeletesCategory() {
	category := data.Category{Id: "category-id"}

	err := s.store.Insert(category)
	s.NoError(err)

	ok, err := s.store.Delete(category.Id, category.AccountId)
	s.NoError(err)
	s.True(ok)

	found, err := s.store.Get(category.Id, category.AccountId)
	s.NoError(err)
	s.Nil(found)
}
