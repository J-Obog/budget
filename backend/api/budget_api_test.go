package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
)

type BudgetApiTestSuite struct {
	ApiTestSuite
	api *BudgetAPI
}

func (s *BudgetApiTestSuite) SetupSuite() {
	s.initDependencies()
	s.api = NewBudgetAPI(
		s.budgetManager,
		s.transactionManager,
		s.categoryManager,
	)
}

func (s *BudgetApiTestSuite) SetupTest() {
	err := s.budgetStore.DeleteAll()
	s.NoError(err)
}

func (s *BudgetApiTestSuite) TestGets() {
	budgetId := "budget-123"
	s.budgetStore.Insert(data.Budget{Id: budgetId, AccountId: testAccountId})

	req := &rest.Request{Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Get(req)
	s.OkResponse(res, data.BudgetMaterialized{})
}

func (s *BudgetApiTestSuite) TestGetFailsIfNoBudgetExists() {
	budgetId := "budget-123"
	req := &rest.Request{Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Get(req)
	s.ErrRepsonse(res, rest.ErrInvalidBudgetId)
}

func (s *BudgetApiTestSuite) TestUpdates() {
	budgetId := "budget-123"
	categoryId := "category-123"
	s.categoryStore.Insert(data.Category{Id: categoryId, AccountId: testAccountId})
	s.budgetStore.Insert(data.Budget{Id: budgetId, AccountId: testAccountId})

	reqBody := rest.BudgetUpdateBody{CategoryId: categoryId, Projected: 12.34}
	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Update(req)
	s.OkResponse(res, data.Budget{})
}

func (s *BudgetApiTestSuite) TestUpdateFailsIfNoBudgetExists() {
	budgetId := "budget-123"
	categoryId := "category-12345"

	reqBody := rest.BudgetUpdateBody{CategoryId: categoryId, Projected: 12.34}
	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidBudgetId)
}

func (s *BudgetApiTestSuite) TestUpdateFailsIfNoCategoryExists() {
	budgetId := "budget-123"
	categoryId := "category-123"
	s.budgetStore.Insert(data.Budget{Id: budgetId, AccountId: testAccountId})

	reqBody := rest.BudgetUpdateBody{CategoryId: categoryId, Projected: 12.34}
	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidCategoryId)
}

func (s *BudgetApiTestSuite) TestUpdateFailsIfCategoryIsTaken() {
	budgetId := "budget-123"
	anotherId := "budget-456"
	categoryId := "category-123"
	month := 10
	year := 2023

	s.budgetStore.Insert(data.Budget{
		Id:        budgetId,
		AccountId: testAccountId,
		Month:     month,
		Year:      year,
	})

	s.budgetStore.Insert(data.Budget{
		Id:         anotherId,
		AccountId:  testAccountId,
		CategoryId: categoryId,
		Month:      month,
		Year:       year,
	})

	reqBody := rest.BudgetUpdateBody{CategoryId: categoryId, Projected: 12.34}
	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrCategoryAlreadyInBudgetPeriod)
}
func (s *BudgetApiTestSuite) TestCreates() {
	categoryId := "category-123"
	s.categoryStore.Insert(data.Category{Id: categoryId, AccountId: testAccountId})

	reqBody := rest.BudgetCreateBody{CategoryId: categoryId, Projected: 12.34, Month: 10, Year: 2023}
	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.OkResponse(res, data.Budget{})
}

func (s *BudgetApiTestSuite) TestCreateFailsIfPeriodIsInvalid() {
	categoryId := "category-123"
	invalidMonth := 55

	reqBody := rest.BudgetCreateBody{
		CategoryId: categoryId,
		Projected:  12.34,
		Month:      invalidMonth,
		Year:       2023,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.ErrRepsonse(res, rest.ErrInvalidDate)
}

func (s *BudgetApiTestSuite) TestCreateFailsIfCategoryIsTaken() {
	categoryId := "category-123"
	month := 10
	year := 2023

	s.budgetStore.Insert(data.Budget{
		Id:         "some-budget-id",
		AccountId:  testAccountId,
		CategoryId: categoryId,
		Month:      month,
		Year:       year,
	})

	reqBody := rest.BudgetCreateBody{
		CategoryId: categoryId,
		Projected:  12.34,
		Month:      month,
		Year:       year,
	}

	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.ErrRepsonse(res, rest.ErrCategoryAlreadyInBudgetPeriod)
}

func (s *BudgetApiTestSuite) TestCreateFailsIfNoCategoryExists() {
	categoryId := "category-123"
	reqBody := rest.BudgetCreateBody{CategoryId: categoryId, Projected: 12.34, Month: 10, Year: 2023}
	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.ErrRepsonse(res, rest.ErrInvalidCategoryId)
}

func (s *BudgetApiTestSuite) TestDeletes() {
	budgetId := "budget-1234"
	s.budgetStore.Insert(data.Budget{Id: budgetId, AccountId: testAccountId})
	req := &rest.Request{Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Delete(req)
	s.OkResponse(res, rest.Success().Data)
}

func (s *BudgetApiTestSuite) TestDeleteFailsIfNoBudgetExists() {
	budgetId := "budget-1234"
	req := &rest.Request{Params: rest.PathParams{"budgetId": budgetId}}
	res := s.api.Delete(req)
	s.OkResponse(res, rest.ErrInvalidBudgetId)
}
