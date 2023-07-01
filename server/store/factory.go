package store

import (
	"log"

	"github.com/J-Obog/paidoff/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	storeImpl = "postgres"
)

type DBConfig struct {
	AccountStore     AccountStore
	CategoryStore    CategoryStore
	BudgetStore      BudgetStore
	TransactionStore TransactionStore
}

func MakeDBConfig(cfg *config.AppConfig) *DBConfig {
	switch storeImpl {
	case "postgres":
		pgDb, err := gorm.Open(postgres.Open(cfg.PostgresUrl), &gorm.Config{
			AllowGlobalUpdate: true,
			NowFunc:           nil,
			Logger:            logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			log.Fatal(err)
		}

		return &DBConfig{
			AccountStore:     &PostgresAccountStore{db: pgDb},
			CategoryStore:    &PostgresCategoryStore{db: pgDb},
			BudgetStore:      &PostgresBudgetStore{db: pgDb},
			TransactionStore: &PostgresTransactionStore{db: pgDb},
		}

	default:
		log.Fatal("Not a supported impl for store")
	}

	return nil
}
