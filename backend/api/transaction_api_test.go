package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/types"
)

type TransactionApiTestSuite struct {
	ApiTestSuite
	api *TransactionAPI
}

func (s *TransactionApiTestSuite) SetupSuite() {
	s.initDependencies()
	s.api = NewTransactionAPI(
		s.transactionManager,
		s.categoryManager,
	)
}

func (s *TransactionApiTestSuite) SetupTest() {
	err := s.transactionStore.DeleteAll()
	s.NoError(err)

	err = s.categoryStore.DeleteAll()
	s.NoError(err)
}

func (s *TransactionApiTestSuite) TestGets() {
	transactionId := "transaction-123"
	s.transactionStore.Insert(data.Transaction{Id: transactionId, AccountId: testAccountId})
	req := &rest.Request{Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Get(req)
	s.OkResponse(res, data.Transaction{})
}

func (s *TransactionApiTestSuite) TestGetFailsIfNoTransactionExists() {
	transactionId := "transaction-123"
	req := &rest.Request{Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Get(req)
	s.ErrRepsonse(res, rest.ErrInvalidTransactionId)
}

func (s *TransactionApiTestSuite) TestUpdates() {
	transactionId := "transaction-123"
	s.transactionStore.Insert(data.Transaction{Id: transactionId, AccountId: testAccountId})

	reqBody := rest.TransactionUpdateBody{
		Type:   data.BudgetType_Expense,
		Amount: 123.45,
		Month:  10,
		Day:    25,
		Year:   2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Update(req)
	s.OkResponse(res, data.Transaction{})
}

func (s *TransactionApiTestSuite) TestUpdateFailsIfNoTransactionExists() {
	transactionId := "transaction-123"

	reqBody := rest.TransactionUpdateBody{
		Type:   data.BudgetType_Expense,
		Amount: 123.45,
		Month:  10,
		Day:    25,
		Year:   2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidTransactionId)
}

func (s *TransactionApiTestSuite) TestUpdateFailsIfDateIsInvalid() {
	transactionId := "transaction-123"
	s.transactionStore.Insert(data.Transaction{Id: transactionId, AccountId: testAccountId})
	invalidMonth := 89
	invalidDay := -23

	reqBody := rest.TransactionUpdateBody{
		Type:   data.BudgetType_Expense,
		Amount: 123.45,
		Month:  invalidMonth,
		Day:    invalidDay,
		Year:   2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidDate)
}

func (s *TransactionApiTestSuite) TestUpdateFailsIfNoCategoryExists() {
	transactionId := "transaction-123"
	categoryId := "category-1234"
	s.transactionStore.Insert(data.Transaction{Id: transactionId, AccountId: testAccountId})

	reqBody := rest.TransactionUpdateBody{
		CategoryId: types.StringPtr(categoryId),
		Type:       data.BudgetType_Expense,
		Amount:     123.45,
		Month:      10,
		Day:        25,
		Year:       2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidCategoryId)
}

func (s *TransactionApiTestSuite) TestUpdateFailsIfNoteIsTooLong() {
	transactionId := "transaction-123"
	longNote := ""
	s.transactionStore.Insert(data.Transaction{Id: transactionId, AccountId: testAccountId})

	reqBody := rest.TransactionUpdateBody{
		Note:   types.StringPtr(longNote),
		Type:   data.BudgetType_Expense,
		Amount: 123.45,
		Month:  10,
		Day:    25,
		Year:   2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidTransactionNote)
}
func (s *TransactionApiTestSuite) TestCreates() {
	reqBody := rest.TransactionCreateBody{
		Type:   data.BudgetType_Expense,
		Amount: 123.45,
		Month:  10,
		Day:    25,
		Year:   2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.OkResponse(res, data.Transaction{})
}

func (s *TransactionApiTestSuite) TestCreateFailsIfDateIsInvalid() {
	invalidMonth := 89
	invalidDay := -23

	reqBody := rest.TransactionCreateBody{
		Type:   data.BudgetType_Expense,
		Amount: 123.45,
		Month:  invalidMonth,
		Day:    invalidDay,
		Year:   2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.ErrRepsonse(res, rest.ErrInvalidDate)
}

func (s *TransactionApiTestSuite) TestCreateFailsIfNoteIsTooLong() {
	longNote := ""

	reqBody := rest.TransactionCreateBody{
		Note:   types.StringPtr(longNote),
		Type:   data.BudgetType_Expense,
		Amount: 123.45,
		Month:  10,
		Day:    15,
		Year:   2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.ErrRepsonse(res, rest.ErrInvalidTransactionNote)
}

func (s *TransactionApiTestSuite) TestCreateFailsIfNoCategoryExists() {
	categoryId := "foobar-id"

	reqBody := rest.TransactionCreateBody{
		CategoryId: types.StringPtr(categoryId),
		Type:       data.BudgetType_Expense,
		Amount:     123.45,
		Month:      10,
		Day:        15,
		Year:       2029,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.ErrRepsonse(res, rest.ErrInvalidCategoryId)
}

func (s *TransactionApiTestSuite) TestDeletes() {
	transactionId := "transaction-1234"
	s.transactionStore.Insert(data.Transaction{Id: transactionId, AccountId: testAccountId})
	req := &rest.Request{Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Delete(req)
	s.ErrRepsonse(res, rest.ErrInvalidTransactionId)
}

func (s *TransactionApiTestSuite) TestDeleteFailsIfNoTransactionExists() {
	transactionId := "transaction-1234"
	req := &rest.Request{Params: rest.PathParams{"transactionId": transactionId}}
	res := s.api.Delete(req)
	s.ErrRepsonse(res, rest.ErrInvalidTransactionId)
}
