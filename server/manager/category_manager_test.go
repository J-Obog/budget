package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

type CategoryManagerTestSuite struct {
	suite.Suite
	categoryStore *mocks.CategoryStore
	budgetStore   *mocks.BudgetStore
	clock         *mocks.Clock
	uid           *mocks.UIDProvider
	queue         *mocks.Queue
	manager       *CategoryManager
}

func (s *CategoryManagerTestSuite) SetupSuite() {
	s.categoryStore = new(mocks.CategoryStore)
	s.budgetStore = new(mocks.BudgetStore)
	s.clock = new(mocks.Clock)
	s.uid = new(mocks.UIDProvider)
	s.queue = new(mocks.Queue)

	s.manager = &CategoryManager{
		store:       s.categoryStore,
		budgetStore: s.budgetStore,
		clock:       s.clock,
		uid:         s.uid,
		queue:       s.queue,
	}
}

func (s *CategoryManagerTestSuite) TestGetsCategoryByRequest() {
	req := testRequest()
	req.ResourceId = testCategoryId

	s.Run("it succeeds", func() {
		category := &data.Category{Id: testCategoryId}

		s.categoryStore.On("Get", testCategoryId, testAccountId).Return(category, nil).Once()

		res := s.manager.GetByRequest(req)
		s.Equal(*res, *rest.Ok(category))
	})

	s.Run("it fails if category doesn't exist", func() {
		s.categoryStore.On("Get", testCategoryId, testAccountId).Return(nil, nil).Once()

		res := s.manager.GetByRequest(req)
		s.Equal(*res, *rest.Err(rest.ErrInvalidCategoryId))
	})
}

func (s *CategoryManagerTestSuite) TestGetsAllCategoriesByRequest() {
	req := testRequest()

	categories := []data.Category{{Id: testCategoryId}}

	s.categoryStore.On("GetAll", req.Account.Id).Return(categories, nil)

	res := s.manager.GetAllByRequest(req)
	s.Equal(*res, *rest.Ok(categories))
}

func (s *CategoryManagerTestSuite) TestCreatesCategoryByRequest() {
	req := testRequest()
	req.ResourceId = testCategoryId

	s.Run("it succeeds", func() {
		createObj := rest.CategoryCreateBody{Name: "Some new name"}

		expected := data.Category{
			Id:        testUuid,
			AccountId: req.Account.Id,
			Name:      createObj.Name,
			Color:     createObj.Color,
			UpdatedAt: testTimestamp,
			CreatedAt: testTimestamp,
		}

		s.clock.On("Now").Return(testTimestamp)
		s.uid.On("GetId").Return(testUuid)
		s.categoryStore.On("GetByName", req.Account.Id, body.Name).Return(nil, nil)
		s.categoryStore.On("Insert", expected).Return(nil)

		s.categoryStore.On("Insert", expectedCategory).Return(nil)

		res := s.manager.CreateByRequest(req)
		s.Equal(*res, *rest.Success())
	})

	s.Run("it fails if name is too long", func() {
		createObj := rest.CategoryCreateBody{Name: testLongString}

		res := s.manager.CreateByRequest(req)
		s.Equal(*res, *rest.Err(rest.ErrInvalidCategoryName))
	})

	t.Run("it fails if name is too short", func() {
		createObj := rest.CategoryCreateBody{Name: testShortString}

		res := s.manager.CreateByRequest(req)
		s.Equal(*res, *rest.Err(rest.ErrInvalidCategoryName))
	})

	t.Run("it fails if name already exists", func(t *testing.T) {
		createObj := rest.CategoryCreateBody{Name: "NewName"}

		categoryNameTaken(s.categoryStore, req.Account.Id, createObj.Name)

		res := s.manager.CreateByRequest(req)
		s.Equal(*res, *rest.Err(rest.ErrCategoryNameAlreadyExists))
	})
}

/*
func TestCategoryManagerCreatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()
		req := testRequest()
		body := rest.CategoryCreateBody{Name: "Some new name"}

		expected := getExpectedCreatedCategory(body, req.Account.Id)

		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockUid.On("GetId").Return(testUuid)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(nil, nil)
		manager.MockCategoryStore.On("Insert", expected).Return(nil)

		res := manager.CreateByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if name is too long", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryCreateBody{Name: genString(config.LimitMaxCategoryNameChars + 5)}
		req := testRequest()
		req.Body = body

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name is too short", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryCreateBody{Name: genString(config.LimitMinCategoryNameChars - 5)}
		req := testRequest()
		req.Body = body

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name already exists", func(t *testing.T) {
		manager := categoryManagerMock()
		req := testRequest()
		body := rest.CategoryCreateBody{Name: "NewName"}

		categoryThatHasName := data.Category{Name: "NewName"}

		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(categoryThatHasName, nil)

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrCategoryNameAlreadyExists)
	})
}

func TestCategoryManagerUpdatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryUpdateBody{Name: "Some updated name"}
		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		existing := data.Category{Name: "Some old name"}
		update := getExpectedCategoryUpdate(body)

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(nil, nil)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockCategoryStore.On("Update", req.ResourceId, req.Account.Id, update, testTimestamp).Return(true, nil)

		res := manager.UpdateByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if category wasn't updated", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryUpdateBody{Name: "Some updated name"}
		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		existing := data.Category{Name: "Some old name"}
		update := getExpectedCategoryUpdate(body)


		clock.TimeIs(testTimestamp)
		uuid.IdIs(testUuid)

		categoryStore.categoryExists()

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(nil, nil)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockCategoryStore.On("Update", req.ResourceId, req.Account.Id, update, testTimestamp).Return(false, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryUpdateBody{Name: "Some updated name"}
		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

	t.Run("it fails if name is too long", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryUpdateBody{Name: genString(config.LimitMaxCategoryNameChars + 5)}
		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		existing := data.Category{Name: "Some old name"}

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name is too short", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryUpdateBody{Name: genString(config.LimitMinCategoryNameChars - 5)}
		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		existing := data.Category{Name: "Some old name"}

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name already exists", func(t *testing.T) {
		manager := categoryManagerMock()
		body := rest.CategoryUpdateBody{Name: "NewName"}
		req := testRequest()
		req.Body = body
		req.ResourceId = testResourceId

		existing := data.Category{Name: "Some old name"}
		categoryThatHasName := data.Category{Name: "NewName"}

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(categoryThatHasName, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrCategoryNameAlreadyExists)
	})
}

func TestCategoryManagerDeletesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		expectedMsg := getExpectedCategoryDeleteMessage(req.ResourceId)


		mockHelper.timeIsNow()
		mockHelper.uuidIs()
		mockHelper.pushesCategoryDeletedMessage()



		manager.MockUid.On("GetId").Return(testUuid)
		manager.MockBudgetStore.On("GetByCategory", req.Account.Id, req.ResourceId).Return([]data.Budget{}, nil)
		manager.MockQueue.On("Push", expectedMsg, queue.QueueName_CategoryDeleted).Return(nil)
		manager.MockCategoryStore.On("Delete", req.ResourceId, req.Account.Id).Return(true, nil)

		res := manager.DeleteByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if category is being used", func(t *testing.T) {
		manager := categoryManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		budgets := []data.Budget{{Id: "some-budget-id"}}

		manager.MockBudgetStore.On("GetByCategory", req.Account.Id, req.ResourceId).Return(budgets, nil)

		res := manager.DeleteByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrCategoryCurrentlyInUse)
	})

	t.Run("it fails if category wasn't deleted", func(t *testing.T) {
		manager := categoryManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		expectedMsg := getExpectedCategoryDeleteMessage(req.ResourceId)

		manager.MockUid.On("GetId").Return(testUuid)
		manager.MockBudgetStore.On("GetByCategory", req.Account.Id, req.ResourceId).Return([]data.Budget{}, nil)
		manager.MockQueue.On("Push", expectedMsg, queue.QueueName_CategoryDeleted).Return(nil)
		manager.MockCategoryStore.On("Delete", req.ResourceId, req.Account.Id).Return(false, nil)

		res := manager.DeleteByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

}

func getExpectedCreatedCategory(body rest.CategoryCreateBody, accountId string) data.Category {
	return data.Category{
		Id:        testUuid,
		AccountId: accountId,
		Name:      body.Name,
		Color:     body.Color,
		UpdatedAt: testTimestamp,
		CreatedAt: testTimestamp,
	}

}

func getExpectedCategoryUpdate(body rest.CategoryUpdateBody) data.CategoryUpdate {
	return data.CategoryUpdate{
		Name:  body.Name,
		Color: body.Color,
	}
}

func getExpectedCategoryDeleteMessage(categoryId string) queue.Message {
	return queue.Message{
		Id: testUuid,
		Data: queue.CategoryDeletedMessage{
			CategoryId: categoryId,
		},
	}
}
*/
