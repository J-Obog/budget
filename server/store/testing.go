package store

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

func NewStoreIntegrationTest() *StoreService {
	cfg := config.Get()
	storeSvc := GetConfiguredStoreService(cfg)

	if err := storeSvc.AccountStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := storeSvc.CategoryStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := storeSvc.BudgetStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := storeSvc.TransactionStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	return storeSvc
}
