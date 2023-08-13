package manager

import (
	"fmt"
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

func TestCategoryManagerTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryManagerTestSuite))
}

type CategoryManagerTestSuite struct {
	ManagerTestSuite
	manager *CategoryManager
}

func (s *CategoryManagerTestSuite) SetupSuite() {
	s.initMocks()
	s.manager = &CategoryManager{
		store:        s.categoryStore,
		clock:        s.clock,
		uuidProvider: s.uuidProvider,
		queue:        s.queue,
	}
}

func (s *CategoryManagerTestSuite) TestGet() {
	expected := &data.Category{
		Id:        "category-123",
		AccountId: "account-456",
	}

	s.categoryStore.EXPECT().Get(
		expected.Id,
		expected.AccountId,
	).Return(expected, nil)

	actual, err := s.manager.Get(expected.Id, expected.AccountId)
	s.Equal(*expected, *actual)
	s.NoError(err)
}

/*func (s *CategoryManagerTestSuite) TestCreate() {
	accountId := "account-123"
	body := rest.CategoryCreateBody{
		Name:  "some-name",
		Color: 12345,
	}

	expected := data.Category{
		Id:        testUuid,
		AccountId: accountId,
		Name:      body.Name,
		Color:     body.Color,
		CreatedAt: testTimestamp,
		UpdatedAt: testTimestamp,
	}

	s.clock.EXPECT().Now().Return(testTimestamp)
	s.uuidProvider.EXPECT().GetUuid().Return(testUuid)
	s.categoryStore.EXPECT().Insert(expected).Return(nil)

	actual, err := s.manager.Create(accountId, body)
	s.Equal(actual, expected)
	s.NoError(err)
}*/

func (s *CategoryManagerTestSuite) TestUpdate() {
	expected := true
	category := &data.Category{Id: "category-123"}
	body := rest.CategoryUpdateBody{
		Name:  "some-name",
		Color: 12345,
	}

	updatedCategory := data.Category{
		Id:        category.Id,
		AccountId: category.AccountId,
		Name:      body.Name,
		Color:     body.Color,
		CreatedAt: category.CreatedAt,
		UpdatedAt: testTimestamp,
	}

	s.clock.EXPECT().Now().Return(testTimestamp)
	s.categoryStore.EXPECT().Update(updatedCategory).Return(expected, nil)

	actual, err := s.manager.Update(category, body)
	s.Equal(expected, actual)
	s.NoError(err)
	s.Equal(*category, updatedCategory)
}

func (s *CategoryManagerTestSuite) TestDelete() {
	expected := true
	id := "some-id"
	accountId := "some-other-id"

	expectedMsg := queue.ToMessage(
		testUuid,
		queue.CategoryDeletedMessage{CategoryId: id, AccountId: accountId},
	)

	s.uuidProvider.EXPECT().GetUuid().Return(testUuid)
	s.queue.EXPECT().Push(expectedMsg, queue.QueueName_CategoryDeleted).Return(nil)
	s.categoryStore.EXPECT().Delete(id, accountId).Return(expected, nil)

	actual, err := s.manager.Delete(id, accountId)

	fmt.Println(actual, expected)
	s.Equal(expected, actual)
	s.NoError(err)
}

func (s *CategoryManagerTestSuite) TestExists() {
	expected := false
	id := "some-id"
	accountId := "some-other-id"

	s.categoryStore.EXPECT().Get(
		id,
		accountId,
	).Return(nil, nil)

	actual, err := s.manager.Exists(
		id,
		accountId,
	)

	s.Equal(expected, actual)
	s.NoError(err)
}

func (s *CategoryManagerTestSuite) TestNameIsUnique() {
	expected := false
	name := "some-name"
	accountId := "some-other-id"

	s.categoryStore.EXPECT().GetByName(
		accountId,
		name,
	).Return(&data.Category{Id: "some-category"}, nil)

	actual, err := s.manager.NameIsUnique(
		accountId,
		name,
	)

	s.Equal(expected, actual)
	s.NoError(err)
}
