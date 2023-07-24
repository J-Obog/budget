package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/assert"
)

func TestCategoryManagerGetsByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		category := data.Category{Id: "category-id-1"}

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(category, nil)

		res := manager.GetByRequest(req)
		assert.Equal(t, res.Data, category)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := categoryManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.GetByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})
}

func TestCategoryManagerGetsAllByRequest(t *testing.T) {
	manager := categoryManagerMock()
	req := testRequest()
	req.ResourceId = testResourceId

	expected := []data.Category{{Id: "some-id"}}

	manager.MockCategoryStore.On("GetAll", req.Account.Id).Return(expected, nil)

	res := manager.GetAllByRequest(req)
	assert.Equal(t, res.Data, expected)
	assert.NoError(t, res.Error)
}

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
