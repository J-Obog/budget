package store

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresStore struct {
	*PostgresAccountStore
	*PostgresBudgetStore
	*PostgresCategoryStore
	*PostgresTransactionStore
}

func NewPostgresStore(url string) *PostgresStore {
	pgDb, err := gorm.Open(postgres.Open(url), &gorm.Config{
		AllowGlobalUpdate: true,
		Logger:            logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{
		PostgresAccountStore:     &PostgresAccountStore{pgDb},
		PostgresTransactionStore: &PostgresTransactionStore{pgDb},
		PostgresCategoryStore:    &PostgresCategoryStore{pgDb},
		PostgresBudgetStore:      &PostgresBudgetStore{pgDb},
	}
}
