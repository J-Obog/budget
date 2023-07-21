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

type StoreConfig struct {
	AccountStore     AccountStore
	CategoryStore    CategoryStore
	BudgetStore      BudgetStore
	TransactionStore TransactionStore
}

func MakeStoreConfig(cfg *config.AppConfig) *StoreConfig {
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

		return &StoreConfig{
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
