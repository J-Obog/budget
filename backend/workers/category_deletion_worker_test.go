package workers

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/types"
	"github.com/stretchr/testify/suite"
)

func TestCategoryDeletionWorker(t *testing.T) {
	suite.Run(t, new(CategoryDeletionWorkerTestSuite))
}

type CategoryDeletionWorkerTestSuite struct {
	WorkerTestSuite
	worker *CategoryDeletionWorker
}

func (s *CategoryDeletionWorkerTestSuite) SetupSuite() {
	s.initDeps()
	s.worker = &CategoryDeletionWorker{
		queue:              s.queue,
		transactionManager: s.transactionManager,
		categoryManager:    s.categoryManager,
	}
}

func (s *CategoryDeletionWorkerTestSuite) SetupTest() {
	err := s.queue.Flush(queue.QueueName_CategoryDeleted)
	s.NoError(err)

	err = s.categoryStore.DeleteAll()
	s.NoError(err)

	err = s.transactionStore.DeleteAll()
	s.NoError(err)
}

func (s *CategoryDeletionWorkerTestSuite) TestNullsCategoryIds() {
	transactionId := "some-cool-id"
	accountId := "acct-1234"
	categoryId := "some-cat-id"

	err := s.transactionStore.Insert(data.Transaction{
		Id:         transactionId,
		AccountId:  accountId,
		CategoryId: types.StringPtr(categoryId),
	})

	s.NoError(err)

	msg := queue.CategoryDeletedMessage{CategoryId: categoryId, AccountId: accountId}

	err = s.queue.Push(queue.ToMessage("some-uuid", &msg), queue.QueueName_CategoryDeleted)
	s.NoError(err)

	err = s.worker.Work()
	s.NoError(err)

	transaction, err := s.transactionStore.Get(transactionId, accountId)
	s.NoError(err)
	s.NotNil(transaction)
	s.Nil(transaction.CategoryId)
}
