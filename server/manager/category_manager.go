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

func (manager *CategoryManager) Create(req data.CategoryCreateRequest) error {
	return nil
}

func (manager *CategoryManager) Update(category *data.Category, req data.CategoryUpdateRequest) error {
	return nil
}

func (manager *CategoryManager) Delete(id string) error {
	return manager.store.Delete(id)
}
