package store

import (
	"log"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
)

func NewStoreIntegrationTest() *StoreConfig {
	cfg, err := config.MakeConfig(config.EnvType_LOCAL)

	if err != nil {
		log.Fatal(err)
	}

	storeCfg := MakeStoreConfig(cfg)

	if err := storeCfg.AccountStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	/*if err := storeCfg.TransactionStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := storeCfg.BudgetStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := storeCfg.CategoryStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}*/

	return storeCfg
}

func testAccount() data.Account {
	return data.Account{
		Id:        "test-12345",
		IsDeleted: false,
		UpdatedAt: 1234,
		CreatedAt: 1234,
	}
}

func testBudget() data.Budget {
	return data.Budget{
		Id:        "test-123",
		UpdatedAt: 1234,
		CreatedAt: 1234,
	}
}

func testTransaction() data.Transaction {
	return data.Transaction{
		Id:        "t-12354",
		UpdatedAt: 1234,
		CreatedAt: 1234,
	}
}

func testCategory() data.Category {
	return data.Category{
		Id:        "testing-1234",
		UpdatedAt: 1234,
		CreatedAt: 1234,
	}
}
