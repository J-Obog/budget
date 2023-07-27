package manager

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestManagers(t *testing.T) {
	suite.Run(t, new(CategoryManagerTestSuite))
}
