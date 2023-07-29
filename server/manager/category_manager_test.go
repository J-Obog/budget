package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

type CategoryManagerTestSuite struct {
	suite.Suite
	store   *mocks.CategoryStore
	clock   *mocks.Clock
	uid     *mocks.UIDProvider
	queue   *mocks.Queue
	manager *CategoryManager
}

func (s *CategoryManagerTestSuite) SetupSuite() {
	s.store = new(mocks.CategoryStore)
	s.clock = new(mocks.Clock)
	s.uid = new(mocks.UIDProvider)
	s.queue = new(mocks.Queue)

	s.manager = &CategoryManager{
		store: s.store,
		clock: s.clock,
		uid:   s.uid,
		queue: s.queue,
	}
}

func (s *CategoryManagerTestSuite) TestGetsCategory() {
	expected := &data.Category{
		Id:        "some-category-id",
		AccountId: "some-account",
	}

	s.store.On("Get", expected.Id, expected.AccountId).Return(expected, nil)

	actual, err := s.manager.Get(expected.Id, expected.AccountId)

	s.NoError(err)
	s.Equal(*expected, *actual)
}

func (s *CategoryManagerTestSuite) TestGetsCategoryByName() {
	expected := &data.Category{
		AccountId: "some-account",
		Name:      "Some Category Name",
	}

	s.store.On("GetByName", expected.AccountId, expected.Name).Return(expected, nil)

	actual, err := s.manager.GetByName(expected.AccountId, expected.Name)

	s.NoError(err)
	s.Equal(*expected, *actual)
}

func (s *CategoryManagerTestSuite) TestGetsAllCategories() {
	accountId := "account-id"

	expected := []data.Category{
		{AccountId: accountId},
	}

	s.store.On("GetAll", accountId).Return(expected, nil)

	actual, err := s.manager.GetAll(accountId)
	s.NoError(err)
	s.ElementsMatch(expected, actual)
}

func (s *CategoryManagerTestSuite) TestCreatesCategory() {
	accountId := "account-id"

	body := rest.CategoryCreateBody{
		Name:  "Foobar",
		Color: 111111,
	}

	expected := data.Category{
		Id:        testUuid,
		AccountId: accountId,
		Name:      body.Name,
		Color:     body.Color,
		UpdatedAt: testTimestamp,
		CreatedAt: testTimestamp,
	}

	s.clock.On("Now").Return(testTimestamp)
	s.uid.On("GetId").Return(testUuid)
	s.store.On("Insert", expected).Return(nil)

	actual, err := s.manager.Create(accountId, body)
	s.NoError(err)
	s.Equal(expected, actual)
}

func (s *CategoryManagerTestSuite) TestUpdatesCategory() {
	existing := &data.Category{
		Id:        "category-id",
		AccountId: "account-id",
	}

	body := rest.CategoryUpdateBody{
		Name:  "Foobar",
		Color: 111111,
	}

	update := data.CategoryUpdate{
		Name:  body.Name,
		Color: body.Color,
	}

	s.clock.On("Now").Return(testTimestamp, nil)
	s.store.On("Update", existing.Id, existing.AccountId, update, testTimestamp).Return(true, nil)

	ok, err := s.manager.Update(existing, body)

	s.NoError(err)
	s.True(ok)
	s.Equal(existing, &data.Category{
		Id:        existing.Id,
		AccountId: existing.AccountId,
		Name:      body.Name,
		Color:     body.Color,
		UpdatedAt: testTimestamp,
	})
}

func (s *CategoryManagerTestSuite) TestDeletesCategory() {
	id := "some-category-to-delete"
	account := "some-account-id"

	expectedMessage := queue.Message{
		Id: testUuid,
		Data: queue.CategoryDeletedMessage{
			CategoryId: id,
		},
	}

	s.queue.On("Push", expectedMessage, queue.QueueName_CategoryDeleted).Return(nil)
	s.store.On("Delete", id, account).Return(true, nil)

	ok, err := s.manager.Delete(id, account)

	s.NoError(err)
	s.True(ok)
}
