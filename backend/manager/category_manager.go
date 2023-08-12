package manager

import (
	"encoding/json"

	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type CategoryManager struct {
	store        store.CategoryStore
	uuidProvider uuid.UuidProvider
	clock        clock.Clock
	queue        queue.Queue
}

func NewCategoryManager(
	store store.CategoryStore,
	uuidProvider uuid.UuidProvider,
	clock clock.Clock,
	queue queue.Queue,
) *CategoryManager {
	return &CategoryManager{
		store:        store,
		uuidProvider: uuidProvider,
		clock:        clock,
		queue:        queue,
	}
}

func (manager *CategoryManager) Get(id string, accountId string) (*data.Category, error) {
	return manager.store.Get(id, accountId)
}

func (manager *CategoryManager) GetAll(accountId string) ([]data.Category, error) {
	return manager.store.GetAll(accountId)
}

func (manager *CategoryManager) Create(accountId string, reqBody rest.CategoryCreateBody) (data.Category, error) {
	timestamp := manager.clock.Now()
	uuid := manager.uuidProvider.GetUuid()

	newCategory := data.Category{
		Id:        uuid,
		AccountId: accountId,
		Name:      reqBody.Name,
		Color:     reqBody.Color,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	if err := manager.store.Insert(newCategory); err != nil {
		return data.Category{}, err
	}

	return newCategory, nil
}

func (manager *CategoryManager) Update(existing *data.Category, body rest.CategoryUpdateBody) (bool, error) {
	existing.Name = body.Name
	existing.Color = body.Color
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *CategoryManager) Delete(id string, accountId string) (bool, error) {
	if err := manager.enqueueCategoryDeleteMsg(id, accountId); err != nil {
		return false, err
	}

	return manager.Delete(id, accountId)
}

func (manager *CategoryManager) Exists(id string, accountId string) (bool, error) {
	category, err := manager.Get(id, accountId)
	if err != nil {
		return false, nil
	}

	return category != nil, nil
}

func (manager *CategoryManager) NameIsUnique(accountId string, name string) (bool, error) {
	category, err := manager.store.GetByName(accountId, name)
	if err != nil {
		return false, err
	}

	return category == nil, nil
}

func (manager *CategoryManager) enqueueCategoryDeleteMsg(categoryId string, accountId string) error {
	categoryDeleteMsg := queue.CategoryDeletedMessage{
		CategoryId: categoryId,
		AccountId:  accountId,
	}

	bytes, err := json.Marshal(&categoryDeleteMsg)

	if err != nil {
		return err
	}

	queueMsg := queue.Message{
		Id:   manager.uuidProvider.GetUuid(),
		Body: bytes,
	}

	return manager.queue.Push(queueMsg, queue.QueueName_CategoryDeleted)
}
