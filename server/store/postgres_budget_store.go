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

func (pg *PostgresBudgetStore) GetBudget(id string, accountId string) (*data.Budget, error) {
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

func (pg *PostgresBudgetStore) GetBudgetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error) {
	var budget data.Budget

	q := data.Budget{
		AccountId:  accountId,
		CategoryId: categoryId,
		Month:      month,
		Year:       year,
	}

	err := pg.db.Where(q).First(&budget).Error
	if err == nil {
		return types.Ptr[data.Budget](budget), nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) GetBudgetsByCategory(accountId string, categoryId string) ([]data.Budget, error) {
	budgets := make([]data.Budget, 0)

	query := data.Budget{
		AccountId:  accountId,
		CategoryId: categoryId,
	}

	err := pg.db.Where(query).Find(&budgets).Error
	if err == nil {
		return budgets, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) GetBudgetsByFilter(accountId string, filter data.BudgetFilter) ([]data.Budget, error) {
	budgets := make([]data.Budget, 0)

	query := data.Budget{
		AccountId: accountId,
		Month:     filter.Month,
		Year:      filter.Year,
	}

	err := pg.db.Where(query).Find(&budgets).Error
	if err == nil {
		return budgets, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) InsertBudget(budget data.Budget) error {
	return pg.db.Create(&budget).Error
}

func (pg *PostgresBudgetStore) UpdateBudget(id string, accountId string, update data.BudgetUpdate, timestamp int64) (bool, error) {
	q := pg.db.Where("id = ?", id)
	q = q.Where("account_id = ?", accountId)

	res := q.UpdateColumns(&data.Budget{
		CategoryId: update.CategoryId,
		Projected:  update.Projected,
		UpdatedAt:  timestamp,
	})

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresBudgetStore) DeleteBudget(id string, accountId string) (bool, error) {
	res := pg.db.Delete(data.Budget{Id: id, AccountId: accountId})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresBudgetStore) DeleteAllBudgets() error {
	err := pg.db.Delete(data.Budget{}).Error
	return err
}
