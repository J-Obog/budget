package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"gorm.io/gorm"
)

type PostgresAccountStore struct {
	db *gorm.DB
}

func (pg *PostgresAccountStore) Get(id string) (*data.Account, error) {
	account := new(data.Account)

	err := pg.db.Where(data.Account{Id: id}).First(account).Error
	if err == nil {
		return account, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresAccountStore) Insert(account data.Account) error {
	return pg.db.Create(&account).Error
}

func (pg *PostgresAccountStore) Update(account data.Account) error {
	err := pg.db.UpdateColumns(&account).Error
	return err
}

func (pg *PostgresAccountStore) Delete(id string) error {
	err := pg.db.Delete(data.Account{Id: id}).Error
	return err
}
