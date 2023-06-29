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

func (pg *PostgresStore) GetAccount(id string) (*data.Account, error) {
	account := new(data.Account)

	err := pg.client.Where(data.Account{Id: id}).First(account).Error
	if err == nil {
		return account, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresStore) InsertAccount(account data.Account) error {
	return pg.client.Create(&account).Error
}

func (pg *PostgresStore) UpdateAccount(account data.Account) error {
	err := pg.client.UpdateColumns(&account).Error
	return err
}

func (pg *PostgresStore) DeleteAccount(id string) error {
	err := pg.client.Delete(data.Account{Id: id}).Error
	return err
}

// Category queries
func (pg *PostgresStore) GetCategory(id string) (*data.Category, error) {
	category := new(data.Category)

	err := pg.client.Where(data.Category{Id: id}).First(category).Error
	if err == nil {
		return category, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresStore) GetCategories(accountId string) ([]data.Category, error) {
	categories := make([]data.Category, 0)

	err := pg.client.Where(data.Category{AccountId: accountId}).Find(&categories).Error
	if err == nil {
		return categories, nil
	}

	return nil, err
}

func (pg *PostgresStore) InsertCategory(category data.Category) error {
	return pg.client.Create(&category).Error
}

func (pg *PostgresStore) UpdateCategory(category data.Category) error {
	err := pg.client.UpdateColumns(&category).Error
	return err
}

func (pg *PostgresStore) DeleteCategory(id string) error {
	err := pg.client.Delete(data.Category{Id: id}).Error
	return err
}

// Budget queries

func (pg *PostgresStore) GetBudget(id string) (*data.Budget, error) {
	budget := new(data.Budget)

	err := pg.client.Where(data.Budget{Id: id}).First(budget).Error
	if err == nil {
		return budget, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresStore) InsertBudget(budget data.Budget) error {
	return pg.client.Create(&budget).Error
}

func (pg *PostgresStore) UpdateBudget(budget data.Budget) error {
	err := pg.client.UpdateColumns(&budget).Error
	return err
}

func (pg *PostgresStore) DeleteBudget(id string) error {
	err := pg.client.Delete(data.Budget{Id: id}).Error
	return err
}

// Transaction queries

func (pg *PostgresStore) GetTransaction(id string) (*data.Transaction, error) {
	transaction := new(data.Transaction)

	err := pg.client.Where(data.Transaction{Id: id}).First(transaction).Error
	if err == nil {
		return transaction, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func (pg *PostgresStore) InsertTransaction(transaction data.Transaction) error {
	return pg.client.Create(&transaction).Error
}

func (pg *PostgresStore) UpdateTransaction(transaction data.Transaction) error {
	err := pg.client.UpdateColumns(&transaction).Error
	return err
}

func (pg *PostgresStore) DeleteTransaction(id string) error {
	err := pg.client.Delete(data.Transaction{Id: id}).Error
	return err
}
