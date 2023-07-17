package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresTransactionStore struct {
	db *gorm.DB
}

func (pg *PostgresTransactionStore) Get(id string, accountId string) (*data.Transaction, error) {
	transaction := new(data.Transaction)
	err := pg.db.Where(data.Transaction{Id: id, AccountId: accountId}).First(transaction).Error
	if err == nil {
		return transaction, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

// TODO: implement
func (pg *PostgresTransactionStore) GetBy(filter data.TransactionFilter) (data.TransactionList, error) {
	transactions := make(data.TransactionList, 0)

	q := pg.db

	if filter.AccountId != nil {
		q = q.Where(data.Category{AccountId: *filter.AccountId})
	}

	err := q.Find(&transactions).Error
	if err == nil {
		return transactions, nil
	}

	return nil, err
}

func (pg *PostgresTransactionStore) Insert(transaction data.Transaction) error {
	return pg.db.Create(&transaction).Error
}

func (pg *PostgresTransactionStore) Update(id string, transaction data.Transaction) (bool, error) {
	res := pg.db.UpdateColumns(&transaction)
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresTransactionStore) Delete(id string) (bool, error) {
	res := pg.db.Delete(data.Transaction{Id: id})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresTransactionStore) DeleteAll() error {
	err := pg.db.Delete(data.Transaction{}).Error
	return err
}
