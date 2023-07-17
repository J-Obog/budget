package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresBudgetStore struct {
	db *gorm.DB
}

func (pg *PostgresBudgetStore) Get(id string, accountId string) (*data.Budget, error) {
	budget := new(data.Budget)

	err := pg.db.Where(data.Budget{Id: id, AccountId: accountId}).First(budget).Error
	if err == nil {
		return budget, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) GetBy(filter data.BudgetFilter) (data.BudgetList, error) {
	budgets := make(data.BudgetList, 0)

	q := pg.db

	if filter.AccountId != nil {
		q = q.Where(data.Budget{AccountId: *filter.AccountId})
	}

	if filter.CategoryId != nil {
		q = q.Where(data.Budget{CategoryId: *filter.CategoryId})
	}

	if filter.Month != nil {
		q = q.Where(data.Budget{Month: *filter.Month})
	}

	if filter.Year != nil {
		q = q.Where(data.Budget{Year: *filter.Year})
	}

	err := q.Find(&budgets).Error
	if err == nil {
		return budgets, nil
	}

	return nil, err
}

func (pg *PostgresBudgetStore) Insert(budget data.Budget) error {
	return pg.db.Create(&budget).Error
}

func (pg *PostgresBudgetStore) Update(id string, budget data.Budget) (bool, error) {
	res := pg.db.UpdateColumns(&budget)
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresBudgetStore) Delete(id string, accountId string) (bool, error) {
	res := pg.db.Delete(data.Budget{Id: id})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresBudgetStore) DeleteAll() error {
	err := pg.db.Delete(data.Budget{}).Error
	return err
}
