package store

import (
	"log"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
)

func setup(cfg *DBConfig) {
	cfg.AccountStore.DeleteAll()
	cfg.BudgetStore.DeleteAll()
	cfg.CategoryStore.DeleteAll()
	cfg.TransactionStore.DeleteAll()
}

func dbIntegrationTest() *DBConfig {
	cfg, err := config.MakeConfig(config.EnvType_LOCAL)

	if err != nil {
		log.Fatal(err)
	}

	return MakeDBConfig(cfg)
}

func testAccount() data.Account {
	return data.Account{
		Id:        "test-12345",
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
