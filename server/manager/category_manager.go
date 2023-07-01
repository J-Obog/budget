package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type CategoryManager struct {
	store store.CategoryStore
	clock clock.Clock
	uid   uid.UIDProvider
}

func (manager *CategoryManager) Get(id string) (*data.Category, error) {
	return manager.store.Get(id)
}

func (manager *CategoryManager) GetByAccount(accountId string) ([]data.Category, error) {
	return manager.store.GetByAccount(accountId)
}

func (manager *CategoryManager) Create(accountId string, req data.CategoryCreateRequest) error {
	now := manager.clock.Now()

	newCategory := data.Category{
		Id:        manager.uid.GetId(),
		AccountId: accountId,
		Name:      req.Name,
		Color:     req.Color,
		UpdatedAt: now,
		CreatedAt: now,
	}
	return manager.store.Insert(newCategory)
}

func (manager *CategoryManager) Update(existing *data.Category, req data.CategoryUpdateRequest) error {
	existing.Color = req.Color
	existing.Name = req.Name
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *CategoryManager) Delete(id string) error {
	return manager.store.Delete(id)
}
