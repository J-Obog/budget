package api

import (
	"net/http"
	"reflect"

	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
	"github.com/stretchr/testify/suite"
)

var (
	veryLongAccountName     = genString(config.LimitMaxAccountNameChars + 1)
	veryShortAccountName    = genString(config.LimitMinAccountNameChars - 1)
	veryLongCategoryName    = genString(config.LimitMaxCategoryNameChars + 1)
	veryShortCategoryName   = genString(config.LimitMinCategoryNameChars - 1)
	veryLongTransactionNote = genString(config.LimitMaxTransactionNoteChars + 1)
)

func genString(length int) string {
	var s string

	for i := 0; i < length; i++ {
		s += "F"
	}

	return s
}

type ApiTestSuite struct {
	suite.Suite
	accountStore       store.AccountStore
	budgetStore        store.BudgetStore
	categoryStore      store.CategoryStore
	transactionStore   store.TransactionStore
	queue              queue.Queue
	accountManager     *manager.AccountManager
	budgetManager      *manager.BudgetManager
	categoryManager    *manager.CategoryManager
	transactionManager *manager.TransactionManager
}

func (s *ApiTestSuite) initDependencies() {
	cfg := config.Get()
	clock := clock.NewClock(cfg)
	uuidProvider := uuid.NewUuidProvider(cfg)

	storeSvc := store.NewStoreService(cfg)
	s.accountStore = storeSvc.AccountStore
	s.budgetStore = storeSvc.BudgetStore
	s.categoryStore = storeSvc.CategoryStore
	s.transactionStore = storeSvc.TransactionStore
	s.queue = queue.NewQueue(cfg)

	s.accountManager = manager.NewAccountManager(s.accountStore, clock)
	s.budgetManager = manager.NewBudgetManager(s.budgetStore, uuidProvider, clock)
	s.categoryManager = manager.NewCategoryManager(s.categoryStore, uuidProvider, clock, s.queue)
	s.transactionManager = manager.NewTransactionManager(s.transactionStore, uuidProvider, clock)
}

func (s *ApiTestSuite) OkResponse(res *rest.Response, expectedShema any) {
	s.Equal(reflect.TypeOf(expectedShema), reflect.TypeOf(res.Data))
	s.Equal(http.StatusOK, res.Status)
}

func (s *ApiTestSuite) ErrRepsonse(res *rest.Response, expectedError *rest.RestError) {
	s.Equal(expectedError, res.Data)
	s.Equal(expectedError.Status, res.Status)
}

func (s *ApiTestSuite) getJSONBody(obj any) *rest.JSONBody {
	jsonb := &rest.JSONBody{}

	err := jsonb.From(obj)
	s.NoError(err)

	return jsonb
}
