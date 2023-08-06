package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
)

type CategoryManager struct {
	store        store.CategoryStore
	uuidProvider uuid.UuidProvider
	clock        clock.Clock
}

func NewCategoryManager(
	store store.CategoryStore,
	uuidProvider uuid.UuidProvider,
	clock clock.Clock,
) *CategoryManager {
	return &CategoryManager{
		store:        store,
		uuidProvider: uuidProvider,
		clock:        clock,
	}
}

func (manager *CategoryManager) Get(id string, accountId string) (*data.Category, error) {
	return nil, nil
}

func (manager *CategoryManager) GetAll(accountId string) ([]data.Category, error) {
	return nil, nil
}

func (manager *CategoryManager) Create(accountId string, createReq rest.CategoryCreateBody) (data.Category, error) {
	return data.Category{}, nil
}

func (manager *CategoryManager) Update(updated *data.Category, updateReq rest.CategoryUpdateBody) error {
	return nil
}

func (manager *CategoryManager) Delete(id string, accountId string) error {
	return nil
}

func (manager *CategoryManager) CheckNameNotTaken(accountId string, name string) error {
	return nil
}
