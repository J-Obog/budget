package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
	"gorm.io/gorm"
)

type PostgresCategoryStore struct {
	db *gorm.DB
}

func (pg *PostgresCategoryStore) GetCategory(id string, accountId string) (*data.Category, error) {
	var category data.Category

	err := pg.db.Where(data.Category{Id: id, AccountId: accountId}).First(&category).Error
	if err == nil {
		return types.Ptr[data.Category](category), nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) GetAllCategories(accountId string) ([]data.Category, error) {
	categories := make([]data.Category, 0)

	err := pg.db.Where(data.Category{AccountId: accountId}).Find(&categories).Error
	if err == nil {
		return categories, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) GetCategoryByName(accountId string, name string) (*data.Category, error) {
	var category data.Category

	err := pg.db.Where(data.Category{AccountId: accountId, Name: name}).First(&category).Error
	if err == nil {
		return types.Ptr[data.Category](category), nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) InsertCategory(category data.Category) error {
	return pg.db.Create(&category).Error
}

func (pg *PostgresCategoryStore) UpdateCategory(id string, accountId string, update data.CategoryUpdate, timestamp int64) (bool, error) {
	q := pg.db.Where("id = ?", id)
	q = q.Where("account_id = ?", accountId)

	res := q.UpdateColumns(&data.Category{
		Name:      update.Name,
		Color:     update.Color,
		UpdatedAt: timestamp,
	})

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresCategoryStore) DeleteCategory(id string, accountId string) (bool, error) {
	res := pg.db.Delete(data.Category{Id: id, AccountId: accountId})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresCategoryStore) DeleteAllCategories() error {
	err := pg.db.Delete(data.Category{}).Error
	return err
}
