package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresCategoryStore struct {
	db *gorm.DB
}

func (pg *PostgresCategoryStore) Get(id string) (*data.Category, error) {
	category := new(data.Category)

	err := pg.db.Where(data.Category{Id: id}).First(category).Error
	if err == nil {
		return category, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) GetByAccount(accountId string) ([]data.Category, error) {
	categories := make([]data.Category, 0)

	err := pg.db.Where(data.Category{AccountId: accountId}).Find(&categories).Error
	if err == nil {
		return categories, nil
	}

	return nil, err
}

func (pg *PostgresCategoryStore) Insert(category data.Category) error {
	return pg.db.Create(&category).Error
}

func (pg *PostgresCategoryStore) Update(category data.Category) error {
	err := pg.db.UpdateColumns(&category).Error
	return err
}

func (pg *PostgresCategoryStore) Delete(id string) error {
	err := pg.db.Delete(data.Category{Id: id}).Error
	return err
}
