package store

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
	"gorm.io/gorm"
)

type PostgresTransactionStore struct {
	db *gorm.DB
}

func (pg *PostgresTransactionStore) Get(id string, accountId string) (types.Optional[data.Transaction], error) {
	transaction := types.OptionalOf[data.Transaction](nil)
	err := pg.db.Where(data.Transaction{Id: id, AccountId: accountId}).First(transaction).Error
	if err == nil {
		return transaction, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return transaction, nil
	}

	return transaction, err
}

// TODO: implement
func (pg *PostgresTransactionStore) GetBy(accountId string, filter data.TransactionFilter) ([]data.Transaction, error) {
	return make([]data.Transaction, 0), nil
}

func (pg *PostgresTransactionStore) GetByPeriodCategory(accountId string, categoryId string, month int, year int) ([]data.Transaction, error) {
	transactions := make([]data.Transaction, 0)

	query := data.Transaction{
		AccountId:  accountId,
		CategoryId: types.OptionalString(categoryId),
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
	q = q.Where("accountId = ?", accountId)

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
