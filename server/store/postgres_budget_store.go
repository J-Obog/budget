package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresBudgetStore struct {
	db *gorm.DB
}

func (pg *PostgresBudgetStore) Get(id string) (*data.Budget, error) {
	budget := new(data.Budget)

	err := pg.db.Where(data.Budget{Id: id}).First(budget).Error
	if err == nil {
		return budget, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) GetByAccount(accountId string) ([]data.Budget, error) {
	budgets := make([]data.Budget, 0)

	err := pg.db.Where(data.Budget{AccountId: accountId}).Find(&budgets).Error
	if err == nil {
		return budgets, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) Insert(budget data.Budget) error {
	return pg.db.Create(&budget).Error
}

func (pg *PostgresBudgetStore) Update(budget data.Budget) error {
	err := pg.db.UpdateColumns(&budget).Error
	return err
}

func (pg *PostgresBudgetStore) Delete(id string) error {
	err := pg.db.Delete(data.Budget{Id: id}).Error
	return err
}
