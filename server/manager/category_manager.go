package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type CategoryManager struct {
	store store.CategoryStore
	clock clock.Clock
	uid   uid.UIDProvider
}

func (manager *CategoryManager) NameExists(accountId string, name string) (bool, error) {
	categories, err := manager.store.GetByAccount(accountId)
	if err != nil {
		return false, err
	}

	for _, category := range categories {
		if category.Name == name {
			return true, nil
		}
	}

	return false, nil
}

func (manager *CategoryManager) Exists(id string, accountId string) (bool, error) {
	category, err := manager.store.Get(id)
	if err != nil {
		return false, err
	}

	if category == nil || category.AccountId != accountId {
		return false, nil
	}

	return true, nil
}

func (manager *CategoryManager) Get(id string, accountId string) (*data.Category, error) {
	category, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}
	if category == nil || category.AccountId != accountId {
		return nil, nil
	}

	return category, nil
}

func (manager *CategoryManager) GetByAccount(accountId string) ([]data.Category, error) {
	return manager.store.GetByAccount(accountId)
}

func (manager *CategoryManager) Create(accountId string, req rest.CategoryCreateBody) error {
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

func (manager *CategoryManager) Update(existing *data.Category, req rest.CategoryUpdateBody) error {
	existing.Color = req.Color
	existing.Name = req.Name
	existing.UpdatedAt = manager.clock.Now()

	return manager.store.Update(*existing)
}

func (manager *CategoryManager) Delete(id string) error {
	return manager.store.Delete(id)
}
