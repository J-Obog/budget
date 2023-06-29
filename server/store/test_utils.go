package store

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, store Store) {
	err := store.Flush()
	assert.NoError(t, err)
}

func getStore(t *testing.T) Store {
	cfg, err := config.MakeConfig("local")
	assert.NoError(t, err)

	s, err := MakeStore(cfg)
	assert.NoError(t, err)

	return s
}

func makeAccount() data.Account {
	return data.Account{
		Id:          "test-12345",
		Name:        "John Doe",
		Email:       "jdoe@gmail.com",
		Password:    "foobar",
		IsDeleted:   false,
		IsActivated: true,
		CreatedAt:   1234,
		UpdatedAt:   4567,
	}
}

func makeBudget() data.Budget {
	return data.Budget{
		Id:         "test-123",
		AccountId:  "test-456",
		CategoryId: nil,
		Name:       "eating",
		Type:       data.BudgetType_EXPENSE,
		Month:      8,
		Year:       2023,
		Projected:  115.60,
		Actual:     67.89,
		CreatedAt:  123,
		UpdatedAt:  456,
	}
}

func makeTransaction() data.Transaction {
	return data.Transaction{
		Id:          "t-12354",
		AccountId:   "testing-45678",
		BudgetId:    "testy-123",
		Description: nil,
		Amount:      590.20,
		Month:       9,
		Day:         5,
		Year:        2028,
		CreatedAt:   123,
		UpdatedAt:   456,
	}
}

func makeCategory() data.Category {
	return data.Category{
		Id:        "testing-1234",
		AccountId: "t-3",
		Name:      "personal",
		Color:     1234,
		CreatedAt: 123,
		UpdatedAt: 456,
	}
}
