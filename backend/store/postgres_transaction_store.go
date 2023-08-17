package store

import (
	"errors"
	"fmt"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
	"gorm.io/gorm"
)

type PostgresTransactionStore struct {
	db *gorm.DB
}

func (pg *PostgresTransactionStore) Get(id string, accountId string) (*data.Transaction, error) {
	var transaction data.Transaction

	err := pg.db.Where(data.Transaction{Id: id, AccountId: accountId}).First(&transaction).Error
	if err == nil {
		return types.Ptr[data.Transaction](transaction), nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresTransactionStore) GetByFilter(accountId string, filter data.TransactionFilter) ([]data.Transaction, error) {
	transactions := make([]data.Transaction, 0)
	query := pg.db

	if filter.MinAmount != nil {
		query.Where("amount >= ?", filter.MinAmount)
	}

	if filter.MaxAmount != nil {
		query.Where("amount <= ?", filter.MaxAmount)
	}

	if filter.StartDate != nil {
		query.Where("make_date(year, month, day) >= ?", dateToSQL(*filter.StartDate))
	}

	if filter.EndDate != nil {
		query.Where("make_date(year, month, day) <= ?", dateToSQL(*filter.EndDate))
	}

	err := query.Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (pg *PostgresTransactionStore) GetByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error) {
	transactions := make([]data.Transaction, 0)

	err := pg.db.Where(&data.Transaction{
		AccountId:  accountId,
		CategoryId: types.StringPtr(categoryId),
		Month:      month,
		Year:       year,
	}).Find(&transactions).Error

	if err == nil {
		return transactions, nil
	}

	return nil, err
}

func (pg *PostgresTransactionStore) Insert(transaction data.Transaction) error {
	return pg.db.Create(&transaction).Error
}

func (pg *PostgresTransactionStore) Update(updated data.Transaction) (bool, error) {
	res := pg.db.Where(
		"id = ? AND account_id = ?",
		updated.Id,
		updated.AccountId).UpdateColumns(&updated)

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresTransactionStore) NullCategoryId(id string, accountId string) (bool, error) {
	res := pg.db.Model(&data.Transaction{}).Where(
		"id = ? AND account_id = ?",
		id,
		accountId,
	).UpdateColumn("category_id", nil)

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresTransactionStore) Delete(id string, accountId string) (bool, error) {
	res := pg.db.Delete(data.Transaction{Id: id, AccountId: accountId})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresTransactionStore) DeleteAll() error {
	err := pg.db.Delete(data.Transaction{}).Error
	return err
}

func dateToSQL(date data.Date) string {
	return fmt.Sprintf("%d-%d-%d", date.Year, date.Month, date.Day)
}
