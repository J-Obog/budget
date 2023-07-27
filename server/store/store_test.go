package store

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestStores(t *testing.T) {
	suite.Run(t, new(AccountStoreTestSuite))
	suite.Run(t, new(BudgetStoreTestSuite))
	suite.Run(t, new(CategoryStoreTestSuite))
	suite.Run(t, new(TransactionStoreTestSuite))
}
