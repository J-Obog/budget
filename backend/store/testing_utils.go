package store

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/stretchr/testify/suite"
)

const (
	testTimestamp = int64(123454321)
)

type StoreTestSuite struct {
	suite.Suite
	accountStore     AccountStore
	budgetStore      BudgetStore
	categoryStore    CategoryStore
	transactionStore TransactionStore
}

func (s *StoreTestSuite) SetupSuite() {
	cfg := config.Get()
	svc := NewStoreService(cfg)
	s.accountStore = svc.AccountStore
	s.budgetStore = svc.BudgetStore
	s.categoryStore = svc.CategoryStore
	s.transactionStore = svc.TransactionStore
}
