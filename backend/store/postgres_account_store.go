package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
	"gorm.io/gorm"
)

type PostgresAccountStore struct {
	db *gorm.DB
}

func (pg *PostgresAccountStore) Get(id string) (*data.Account, error) {
	var account data.Account

	err := pg.db.Where(data.Account{Id: id}).First(&account).Error
	if err == nil {
		return types.Ptr[data.Account](account), nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresAccountStore) Insert(account data.Account) error {
	return pg.db.Create(&account).Error
}

func (pg *PostgresAccountStore) Update(updated data.Account) (bool, error) {
	res := pg.db.UpdateColumns(&updated)
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresAccountStore) SoftDelete(id string) (bool, error) {
	res := pg.db.Where("id = ?", id).UpdateColumns(&data.Account{IsDeleted: true})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresAccountStore) Delete(id string) (bool, error) {
	res := pg.db.Delete(data.Account{Id: id})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresAccountStore) DeleteAll() error {
	err := pg.db.Delete(data.Account{}).Error
	return err
}
