package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type CategoryManager struct {
	store store.CategoryStore
	clock clock.Clock
	uid   uid.UIDProvider
	queue queue.Queue
}

func NewCategoryManager(store store.CategoryStore, clock clock.Clock, uid uid.UIDProvider, queue queue.Queue) *CategoryManager {
	return &CategoryManager{
		store: store,
		clock: clock,
		uid:   uid,
		queue: queue,
	}
}

func (manager *CategoryManager) Get(id string, accountId string) (*data.Category, error) {
	return manager.store.Get(id, accountId)
}

func (manager *CategoryManager) GetByName(accountId string, name string) (*data.Category, error) {
	return manager.store.GetByName(accountId, name)
}

func (manager *CategoryManager) GetAll(accountId string) ([]data.Category, error) {
	return manager.store.GetAll(accountId)
}

func (manager *CategoryManager) Create(accountId string, body rest.CategoryCreateBody) (data.Category, error) {
	id := manager.uid.GetId()
	timestamp := manager.clock.Now()

	category := data.Category{
		Id:        id,
		AccountId: accountId,
		Name:      body.Name,
		Color:     body.Color,
		UpdatedAt: timestamp,
		CreatedAt: timestamp,
	}

	err := manager.store.Insert(category)

	return category, err
}

func (manager *CategoryManager) Update(existing *data.Category, body rest.CategoryUpdateBody) (bool, error) {
	timestamp := manager.clock.Now()

	update := data.CategoryUpdate{
		Name:  body.Name,
		Color: body.Color,
	}

	existing.Name = update.Name
	existing.Color = update.Color
	existing.UpdatedAt = timestamp

	return manager.store.Update(existing.Id, existing.AccountId, update, timestamp)
}

func (manager *CategoryManager) Delete(id string, accountId string) (bool, error) {
	msgId := manager.uid.GetId()

	msg := queue.Message{
		Id: msgId,
		Data: queue.CategoryDeletedMessage{
			CategoryId: id,
		},
	}

	if err := manager.queue.Push(msg, queue.QueueName_CategoryDeleted); err != nil {
		return false, err
	}

	return manager.store.Delete(id, accountId)
}
