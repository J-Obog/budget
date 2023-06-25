package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresStore struct {
	client *gorm.DB
}

func NewPostgresStore(dbClient *gorm.DB) *PostgresStore {
	return &PostgresStore{
		client: dbClient,
	}
}

func (pg *PostgresStore) Flush() error {
	if err := pg.client.Delete(data.Account{}).Error; err != nil {
		return err
	}

	if err := pg.client.Delete(data.Budget{}).Error; err != nil {
		return err
	}

	if err := pg.client.Delete(data.Transaction{}).Error; err != nil {
		return err
	}

	if err := pg.client.Delete(data.Category{}).Error; err != nil {
		return err
	}

	return nil
}

// Account queries

func (pg *PostgresStore) GetAccount(id string) (account *data.Account, e error) {
	err := pg.client.Where(data.Account{Id: id}).First(account).Error

	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	e = err
	return
}

func (pg *PostgresStore) InsertAccount(account data.Account) error {
	return pg.client.Create(&account).Error
}

func (pg *PostgresStore) UpdateAccount(account data.Account) error {
	err := pg.client.Save(&account).Error
	return err
}

func (pg *PostgresStore) DeleteAccount(id string) error {
	err := pg.client.Delete(data.Account{Id: id}).Error
	return err
}

// Category queries
func (pg *PostgresStore) GetCategory(id string) (*data.Category, error) {
	return nil, nil
}

func (pg *PostgresStore) InsertCategory(category data.Category) error {
	return nil
}

func (pg *PostgresStore) UpdateCategory(category data.Category) error {
	return nil
}

func (pg *PostgresStore) DeleteCategory(id string) error {
	return nil
}

// Budget queries

func (pg *PostgresStore) GetBudget(id string) (*data.Budget, error) {
	return nil, nil
}

func (pg *PostgresStore) InsertBudget(budget data.Budget) error {
	return nil
}

func (pg *PostgresStore) UpdateBudget(budget data.Budget) error {
	return nil
}

func (pg *PostgresStore) DeleteBudget(id string) error {
	return nil
}

// Transaction queries

func (pg *PostgresStore) GetTransaction(id string) (*data.Transaction, error) {
	return nil, nil
}

func (pg *PostgresStore) InsertTransaction(transaction data.Transaction) error {
	return nil
}

func (pg *PostgresStore) UpdateTransaction(transaction data.Transaction) error {
	return nil
}

func (pg *PostgresStore) DeleteTransaction(id string) error {
	return nil
}
