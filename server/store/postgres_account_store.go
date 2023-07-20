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

func (pg *PostgresAccountStore) Get(id string) (types.Optional[data.Account], error) {
	account := types.OptionalOf[data.Account](nil)

	err := pg.db.Where(data.Account{Id: id}).First(account).Error
	if err == nil {
		return account, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return account, nil
	}

	return account, err
}

func (pg *PostgresAccountStore) Insert(account data.Account) error {
	return pg.db.Create(&account).Error
}

func (pg *PostgresAccountStore) Update(id string, update data.AccountUpdate, timestamp int64) (bool, error) {
	/*res := pg.db.UpdateColumns(&data.Account{
		Name:      update.Name,
		UpdatedAt: timestamp,
	})

	return res.Error*/
	return true, nil
}

func (pg *PostgresAccountStore) Delete(id string) (bool, error) {
	res := pg.db.Delete(data.Account{Id: id})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresAccountStore) DeleteAll() error {
	err := pg.db.Delete(data.Account{}).Error
	return err
}
