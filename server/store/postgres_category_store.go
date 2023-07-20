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

func (pg *PostgresCategoryStore) Get(id string, accountId string) (types.Optional[data.Category], error) {
	category := types.OptionalOf[data.Category](nil)

	err := pg.db.Where(data.Category{Id: id, AccountId: accountId}).First(category).Error
	if err == nil {
		return category, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return category, nil
	}

	return category, err
}

func (pg *PostgresCategoryStore) GetAll(accountId string) ([]data.Category, error) {
	categories := make([]data.Category, 0)

	err := pg.db.Where(data.Category{AccountId: accountId}).Find(&categories).Error
	if err == nil {
		return categories, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) GetByName(accountId string, name string) (types.Optional[data.Category], error) {
	category := types.OptionalOf[data.Category](nil)

	err := pg.db.Where(data.Category{AccountId: accountId, Name: name}).First(category).Error
	if err == nil {
		return category, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return category, nil
	}

	return category, err
}

func (pg *PostgresCategoryStore) Insert(category data.Category) error {
	return pg.db.Create(&category).Error
}

func (pg *PostgresCategoryStore) Update(id string, accountId string, update data.CategoryUpdate, timestamp int64) (bool, error) {
	q := pg.db.Where("id = ?", id)
	q = q.Where("accountId = ?", accountId)

	res := q.UpdateColumns(&data.Category{
		Name:  update.Name,
		Color: update.Color,
	})

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresCategoryStore) Delete(id string, accountId string) (bool, error) {
	res := pg.db.Delete(data.Category{Id: id, AccountId: accountId})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresCategoryStore) DeleteAll() error {
	err := pg.db.Delete(data.Category{}).Error
	return err
}
