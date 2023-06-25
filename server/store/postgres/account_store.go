package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresAccountStore struct {
	client *gorm.DB
}

func NewPostgresAccountStore(dbClient *gorm.DB) *PostgresAccountStore {
	return &PostgresAccountStore{
		client: dbClient,
	}
}

func (pg *PostgresAccountStore) Get(id string) (*data.Account, error) {
	account := &data.Account{}

	err := pg.client.First(account).Where(data.Account{Id: id}).Error

	if err == nil {
		return account, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresAccountStore) Insert(account data.Account) error {
	err := pg.client.Create(&account).Error
	return err
}

func (pg *PostgresAccountStore) Update(account data.Account) error {
	err := pg.client.Save(&account).Error
	return err
}

func (pg *PostgresAccountStore) Delete(id string) error {
	err := pg.client.Delete(data.Account{Id: id}).Error
	return err
}
