package manager

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/assert"
)

func TestCategoryManagerGetsByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryGet()
		category := testCategory()

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(category, nil)

		res := manager.GetByRequest(req)

		assert.Equal(t, res.Data, category)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryGet()

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.GetByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})
}

func TestCategoryManagerGetsAllByRequest(t *testing.T) {
	manager := categoryManagerMock()
	req := categoryGet()

	c1 := testCategory()
	c1.Id = "test-id-1"

	c2 := testCategory()
	c2.Id = "test-id-2"

	c3 := testCategory()
	c3.Id = "test-id-3"

	expected := []data.Category{*c1, *c2, *c3}

	manager.MockCategoryStore.On("GetAll", req.Account.Id).Return(expected, nil)

	res := manager.GetAllByRequest(req)

	assert.Equal(t, res.Data, expected)
	assert.NoError(t, res.Error)
}

func TestCategoryManagerCreatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()

		id := "some-gen-uuid"
		timestamp := int64(123454321)

		req := categoryCreate()
		body := req.Body.(rest.CategoryCreateBody)

		expected := data.Category{
			Id:        id,
			AccountId: req.Account.Id,
			Name:      body.Name,
			Color:     body.Color,
			UpdatedAt: timestamp,
			CreatedAt: timestamp,
		}

		manager.MockClock.On("Now").Return(timestamp)
		manager.MockUid.On("GetId").Return(id)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(nil, nil)
		manager.MockCategoryStore.On("Insert", expected).Return(nil)

		res := manager.CreateByRequest(req)

		assert.NoError(t, res.Error)
	})

	t.Run("it fails if name is too long", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryCreateNameTooLong()
		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name is too short", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryCreateNameTooShort()
		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name already exists", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryCreate()
		body := req.Body.(rest.CategoryCreateBody)

		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(testCategory(), nil)

		res := manager.CreateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrCategoryNameAlreadyExists)
	})
}

func TestCategoryManagerUpdatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()
		timestamp := int64(123454321)
		req := categoryUpdate()
		body := req.Body.(rest.CategoryUpdateBody)
		existing := testCategory()

		update := data.CategoryUpdate{
			Name:  body.Name,
			Color: body.Color,
		}

		manager.MockClock.On("Now").Return(timestamp)
		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(nil, nil)
		manager.MockCategoryStore.On("Update", req.ResourceId, req.Account.Id, update, timestamp).Return(true, nil)

		res := manager.UpdateByRequest(req)

		assert.NoError(t, res.Error)
	})

	t.Run("it fails if category wasn't updated", func(t *testing.T) {
		manager := categoryManagerMock()
		timestamp := int64(123454321)
		req := categoryUpdate()
		body := req.Body.(rest.CategoryUpdateBody)
		existing := testCategory()

		update := data.CategoryUpdate{
			Name:  body.Name,
			Color: body.Color,
		}

		manager.MockClock.On("Now").Return(timestamp)
		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(nil, nil)
		manager.MockCategoryStore.On("Update", req.ResourceId, req.Account.Id, update, timestamp).Return(false, nil)

		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryUpdate()

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

	t.Run("it fails if name is too long", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryUpdateNameTooLong()
		existing := testCategory()

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)

		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name is too short", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryUpdateNameTooShort()
		existing := testCategory()

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)

		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryName)
	})

	t.Run("it fails if name already exists", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryUpdate()
		body := req.Body.(rest.CategoryUpdateBody)
		existing := testCategory()
		categoryThatHasName := testCategory()

		manager.MockCategoryStore.On("Get", req.ResourceId, req.Account.Id).Return(existing, nil)
		manager.MockCategoryStore.On("GetByName", req.Account.Id, body.Name).Return(categoryThatHasName, nil)

		res := manager.UpdateByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrCategoryNameAlreadyExists)
	})
}

func TestCategoryManagerDeletesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryDelete()
		id := "some-generate-msg-id"

		expectedMsg := queue.Message{
			Id: id,
			Data: queue.CategoryDeletedMessage{
				CategoryId: req.ResourceId,
			},
		}

		manager.MockUid.On("GetId").Return(id)
		manager.MockBudgetStore.On("GetByCategory", req.Account.Id, req.ResourceId).Return([]data.Budget{}, nil)
		manager.MockQueue.On("Push", expectedMsg, queue.QueueName_CategoryDeleted).Return(nil)
		manager.MockCategoryStore.On("Delete", req.ResourceId, req.Account.Id).Return(true, nil)

		res := manager.DeleteByRequest(req)

		assert.NoError(t, res.Error)
	})

	t.Run("it fails if category is being used", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryDelete()

		budgets := []data.Budget{*testBudget()}

		manager.MockBudgetStore.On("GetByCategory", req.Account.Id, req.ResourceId).Return(budgets, nil)

		res := manager.DeleteByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrCategoryCurrentlyInUse)
	})

	t.Run("it fails if category wasn't deleted", func(t *testing.T) {
		manager := categoryManagerMock()
		req := categoryDelete()
		id := "some-generate-msg-id"

		expectedMsg := queue.Message{
			Id: id,
			Data: queue.CategoryDeletedMessage{
				CategoryId: req.ResourceId,
			},
		}

		manager.MockUid.On("GetId").Return(id)
		manager.MockBudgetStore.On("GetByCategory", req.Account.Id, req.ResourceId).Return([]data.Budget{}, nil)
		manager.MockQueue.On("Push", expectedMsg, queue.QueueName_CategoryDeleted).Return(nil)
		manager.MockCategoryStore.On("Delete", req.ResourceId, req.Account.Id).Return(false, nil)

		res := manager.DeleteByRequest(req)

		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

}
