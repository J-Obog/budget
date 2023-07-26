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

type StoreService struct {
	AccountStore     AccountStore
	CategoryStore    CategoryStore
	BudgetStore      BudgetStore
	TransactionStore TransactionStore
}

func GetConfiguredStoreService(cfg *config.AppConfig) *StoreService {
	switch storeImpl {
	case "postgres":
		return getPostgresService(cfg.PostgresUrl)
	default:
		log.Fatal("Not a supported impl for store")
	}

	return nil
}

func getPostgresService(url string) *StoreService {
	pgDb, err := gorm.Open(postgres.Open(url), &gorm.Config{
		AllowGlobalUpdate: true,
		Logger:            logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err)
	}

	return &StoreService{
		AccountStore:     &PostgresAccountStore{db: pgDb},
		CategoryStore:    &PostgresCategoryStore{db: pgDb},
		BudgetStore:      &PostgresBudgetStore{db: pgDb},
		TransactionStore: &PostgresTransactionStore{db: pgDb},
	}
}
