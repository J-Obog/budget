package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresTransactionStore struct {
	db *gorm.DB
}

func (pg *PostgresTransactionStore) Get(id string) (*data.Transaction, error) {
	transaction := new(data.Transaction)

	err := pg.db.Where(data.Transaction{Id: id}).First(transaction).Error
	if err == nil {
		return transaction, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresTransactionStore) GetByAccount(accountId string) ([]data.Transaction, error) {
	transactions := make([]data.Transaction, 0)

	err := pg.db.Where(data.Transaction{AccountId: accountId}).Find(&transactions).Error
	if err == nil {
		return transactions, nil
	}

	return nil, err
}

func (pg *PostgresTransactionStore) Insert(transaction data.Transaction) error {
	return pg.db.Create(&transaction).Error
}

func (pg *PostgresTransactionStore) Update(transaction data.Transaction) error {
	err := pg.db.UpdateColumns(&transaction).Error
	return err
}

func (pg *PostgresTransactionStore) Delete(id string) error {
	err := pg.db.Delete(data.Transaction{Id: id}).Error
	return err
}

func (pg *PostgresTransactionStore) DeleteAll() error {
	err := pg.db.Delete(data.Transaction{}).Error
	return err
}
