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

func (pg *PostgresTransactionStore) GetBy(accountId string, filter data.TransactionFilter) ([]data.Transaction, error) {
	transactions := make([]data.Transaction, 0)

	q := pg.db.Where("amount >= ?", filter.GreaterThan)
	q = q.Where("amount <= ?", filter.LessThan)
	q = q.Where("make_date(year, month, day) <= ?", dateToSQL(filter.Before))
	q = q.Where("make_date(year, month, day) >= ?", dateToSQL(filter.After))

	err := q.Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (pg *PostgresTransactionStore) GetByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error) {
	transactions := make([]data.Transaction, 0)

	query := data.Transaction{
		AccountId:  accountId,
		CategoryId: types.StringPtr(categoryId),
		Month:      month,
		Year:       year,
	}

	err := pg.db.Where(query).Find(&transactions).Error
	if err == nil {
		return transactions, nil
	}

	return nil, err
}

func (pg *PostgresTransactionStore) Insert(transaction data.Transaction) error {
	return pg.db.Create(&transaction).Error
}

func (pg *PostgresTransactionStore) Update(id string, accountId string, update data.TransactionUpdate, timestamp int64) (bool, error) {
	q := pg.db.Where("id = ?", id)
	q = q.Where("account_id = ?", accountId)

	res := q.UpdateColumns(&data.Transaction{
		CategoryId: update.CategoryId,
		Note:       update.Note,
		Type:       update.Type,
		Amount:     update.Amount,
		Month:      update.Month,
		Day:        update.Day,
		Year:       update.Year,
	})

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
