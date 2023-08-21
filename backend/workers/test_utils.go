package workers

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
	"github.com/stretchr/testify/suite"
)

type WorkerTestSuite struct {
	suite.Suite
	queue              queue.Queue
	categoryStore      store.CategoryStore
	transactionStore   store.TransactionStore
	categoryManager    *manager.CategoryManager
	transactionManager *manager.TransactionManager
}

func (s *WorkerTestSuite) initDeps() {
	cfg := config.Get()
	clock := clock.NewClock(cfg)
	uuidProvider := uuid.NewUuidProvider(cfg)
	storeSvc := store.NewStoreService(cfg)
	s.categoryStore = storeSvc.CategoryStore
	s.transactionStore = storeSvc.TransactionStore
	s.queue = queue.NewQueue(cfg)
	s.categoryManager = manager.NewCategoryManager(s.categoryStore, uuidProvider, clock, s.queue)
	s.transactionManager = manager.NewTransactionManager(s.transactionStore, uuidProvider, clock)
}
