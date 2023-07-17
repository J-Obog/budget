package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresCategoryStore struct {
	db *gorm.DB
}

func (pg *PostgresCategoryStore) Get(id string, accountId string) (*data.Category, error) {
	category := new(data.Category)

	err := pg.db.Where(data.Category{Id: id, AccountId: accountId}).First(category).Error
	if err == nil {
		return category, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) GetBy(filter data.CategoryFilter) (data.CategoryList, error) {
	categories := make(data.CategoryList, 0)

	q := pg.db

	if filter.AccountId != nil {
		q = q.Where(data.Category{AccountId: *filter.AccountId})
	}

	if filter.Name != nil {
		q = q.Where(data.Category{Name: *filter.Name})
	}

	err := q.Find(&categories).Error
	if err == nil {
		return categories, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) Insert(category data.Category) error {
	return pg.db.Create(&category).Error
}

func (pg *PostgresCategoryStore) Update(id string, category data.Category) (bool, error) {
	res := pg.db.UpdateColumns(&category)
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresCategoryStore) Delete(id string) (bool, error) {
	res := pg.db.Delete(data.Category{Id: id})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresCategoryStore) DeleteAll() error {
	err := pg.db.Delete(data.Category{}).Error
	return err
}
