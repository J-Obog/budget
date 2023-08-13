package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
	"gorm.io/gorm"
)

type PostgresBudgetStore struct {
	db *gorm.DB
}

func (pg *PostgresBudgetStore) Get(id string, accountId string) (*data.Budget, error) {
	var budget data.Budget

	err := pg.db.Where(data.Budget{Id: id, AccountId: accountId}).First(&budget).Error
	if err == nil {
		return types.Ptr[data.Budget](budget), nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) GetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error) {
	var budget data.Budget

	err := pg.db.Where(&data.Budget{
		AccountId:  accountId,
		CategoryId: categoryId,
		Month:      month,
		Year:       year,
	}).First(&budget).Error

	if err == nil {
		return types.Ptr[data.Budget](budget), nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) GetByCategory(accountId string, categoryId string) ([]data.Budget, error) {
	budgets := make([]data.Budget, 0)

	err := pg.db.Where(&data.Budget{
		AccountId:  accountId,
		CategoryId: categoryId,
	}).Find(&budgets).Error

	if err == nil {
		return budgets, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) GetBy(accountId string, filter data.BudgetFilter) ([]data.Budget, error) {
	budgets := make([]data.Budget, 0)

	err := pg.db.Where(&data.Budget{
		AccountId: accountId,
		Month:     filter.Month,
		Year:      filter.Year,
	}).Find(&budgets).Error

	if err == nil {
		return budgets, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) Insert(budget data.Budget) error {
	return pg.db.Create(&budget).Error
}

func (pg *PostgresBudgetStore) Update(updated data.Budget) (bool, error) {
	res := pg.db.Where(
		"id = ? AND account_id = ?",
		updated.Id,
		updated.AccountId).UpdateColumns(&updated)

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresBudgetStore) Delete(id string, accountId string) (bool, error) {
	res := pg.db.Delete(data.Budget{Id: id, AccountId: accountId})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresBudgetStore) DeleteAll() error {
	err := pg.db.Delete(data.Budget{}).Error
	return err
}
