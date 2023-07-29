package manager

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestManagers(t *testing.T) {
	suite.Run(t, new(AccountManagerTestSuite))
	suite.Run(t, new(BudgetManagerTestSuite))
	suite.Run(t, new(CategoryManagerTestSuite))
	suite.Run(t, new(TransactionManagerTestSuite))
}
