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
