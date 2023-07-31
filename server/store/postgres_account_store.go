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

func (pg *PostgresAccountStore) GetAccount(id string) (*data.Account, error) {
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

func (pg *PostgresAccountStore) InsertAccount(account data.Account) error {
	return pg.db.Create(&account).Error
}

func (pg *PostgresAccountStore) UpdateAccount(id string, update data.AccountUpdate, timestamp int64) (bool, error) {
	q := pg.db.Where("id = ?", id)

	res := q.UpdateColumns(&data.Account{
		Name:      update.Name,
		UpdatedAt: timestamp,
	})

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresAccountStore) SoftDeleteAccount(id string) (bool, error) {
	q := pg.db.Where("id = ?", id)

	res := q.UpdateColumns(&data.Account{
		IsDeleted: true,
	})

	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresAccountStore) DeleteAccount(id string) (bool, error) {
	res := pg.db.Delete(data.Account{Id: id})
	return (res.RowsAffected == 1), res.Error
}

func (pg *PostgresAccountStore) DeleteAllAccounts() error {
	err := pg.db.Delete(data.Account{}).Error
	return err
}
