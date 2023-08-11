package manager

import (
	"github.com/J-Obog/paidoff/mocks"
	"github.com/stretchr/testify/suite"
)

const (
	testTimestamp = 123456789
	testUuid      = "test-uuid-123"
)

type ManagerTestSuite struct {
	suite.Suite
	clock            *mocks.Clock
	uuidProvider     *mocks.UuidProvider
	queue            *mocks.Queue
	accountStore     *mocks.AccountStore
	budgetStore      *mocks.BudgetStore
	categoryStore    *mocks.CategoryStore
	transactionStore *mocks.TransactionStore
}

func (s *ManagerTestSuite) initMocks() {
	s.clock = new(mocks.Clock)
	s.uuidProvider = new(mocks.UuidProvider)
	s.queue = new(mocks.Queue)
	s.accountStore = new(mocks.AccountStore)
	s.budgetStore = new(mocks.BudgetStore)
	s.categoryStore = new(mocks.CategoryStore)
	s.transactionStore = new(mocks.TransactionStore)
}
