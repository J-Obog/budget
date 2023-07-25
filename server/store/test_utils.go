package store

import (
	"log"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
)

type StoreIntegrationTest struct {
	StoreConfig
}

func NewStoreIntegrationTest() *StoreIntegrationTest {
	cfg := config.Get()
	storeCfg := MakeStoreConfig(cfg)

	return &StoreIntegrationTest{
		StoreConfig: *storeCfg,
	}
}

func (it *StoreIntegrationTest) Setup() {
	if err := it.AccountStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := it.CategoryStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := it.BudgetStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}

	if err := it.TransactionStore.DeleteAll(); err != nil {
		log.Fatal(err)
	}
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
		AccountId: "test-acct-12345678",
		UpdatedAt: 1234,
		CreatedAt: 1234,
	}
}
