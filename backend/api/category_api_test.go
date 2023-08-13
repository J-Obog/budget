package api

import (
	"testing"
	"time"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

func TestCategoryApi(t *testing.T) {
	suite.Run(t, new(CategoryApiTestSuite))
}

type CategoryApiTestSuite struct {
	ApiTestSuite
	api *CategoryAPI
}

func (s *CategoryApiTestSuite) SetupSuite() {
	s.initDependencies()
	s.api = NewCategoryAPI(
		s.categoryManager,
		s.budgetManager,
	)
}

func (s *CategoryApiTestSuite) SetupTest() {
	err := s.categoryStore.DeleteAll()
	s.NoError(err)

	err = s.budgetStore.DeleteAll()
	s.NoError(err)
}

func (s *CategoryApiTestSuite) TestGets() {
	categoryId := "category-123"
	s.categoryStore.Insert(data.Category{Id: categoryId, AccountId: testAccountId})
	req := &rest.Request{Params: rest.PathParams{"categoryId": categoryId}}
	res := s.api.Get(req)
	s.OkResponse(res, &data.Category{})
}

func (s *CategoryApiTestSuite) TestGetFailsIfNoCategoryExists() {
	categoryId := "category-123"
	req := &rest.Request{Params: rest.PathParams{"categoryId": categoryId}}
	res := s.api.Get(req)
	s.ErrRepsonse(res, rest.ErrInvalidCategoryId)
}

func (s *CategoryApiTestSuite) TestGetsAll() {
	s.categoryStore.Insert(data.Category{Id: "1", AccountId: testAccountId})
	s.categoryStore.Insert(data.Category{Id: "2", AccountId: testAccountId})
	s.categoryStore.Insert(data.Category{Id: "3", AccountId: testAccountId})

	req := &rest.Request{}
	res := s.api.GetAll(req)

	s.OkResponse(res, []data.Category{})
}

func (s *CategoryApiTestSuite) TestUpdates() {
	categoryId := "category-123"
	s.categoryStore.Insert(data.Category{Id: categoryId, AccountId: testAccountId})
	reqBody := rest.CategoryUpdateBody{Name: "some-name", Color: 1011011}
	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"categoryId": categoryId}}
	res := s.api.Update(req)
	s.OkResponse(res, &data.Category{})
}

func (s *CategoryApiTestSuite) TestUpdateFailsIfNoCategoryExists() {
	categoryId := "category-123"
	reqBody := rest.CategoryUpdateBody{Name: "some-name", Color: 1011011}
	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"categoryId": categoryId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidCategoryId)
}

func (s *CategoryApiTestSuite) TestUpdateFailsIfNameIsInvalid() {
	categoryId := "category-123"
	s.categoryStore.Insert(data.Category{Id: categoryId, AccountId: testAccountId})

	invalidNames := []string{}

	for _, invalidName := range invalidNames {
		reqBody := rest.CategoryUpdateBody{Name: invalidName, Color: 1011011}
		req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"categoryId": categoryId}}
		res := s.api.Update(req)
		s.ErrRepsonse(res, rest.ErrInvalidCategoryName)
	}
}

func (s *CategoryApiTestSuite) TestUpdateFailsIfNameIsTaken() {
	categoryId := "category-123"
	name := "some-name"

	s.categoryStore.Insert(data.Category{Id: "another-id", AccountId: testAccountId, Name: name})
	s.categoryStore.Insert(data.Category{Id: categoryId, AccountId: testAccountId})

	reqBody := rest.CategoryUpdateBody{Name: name, Color: 1011011}
	req := &rest.Request{Body: s.getJSONBody(reqBody), Params: rest.PathParams{"categoryId": categoryId}}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrCategoryNameAlreadyExists)
}

func (s *CategoryApiTestSuite) TestCreates() {
	reqBody := rest.CategoryCreateBody{Name: "some-name", Color: 1011011}
	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.OkResponse(res, data.Category{})
}

func (s *CategoryApiTestSuite) TestCreateFailsIfNameIsInvalid() {
	invalidNames := []string{}

	for _, invalidName := range invalidNames {
		reqBody := rest.CategoryCreateBody{Name: invalidName, Color: 1011011}
		req := &rest.Request{Body: s.getJSONBody(reqBody)}
		res := s.api.Create(req)
		s.ErrRepsonse(res, rest.ErrInvalidCategoryName)
	}
}

func (s *CategoryApiTestSuite) TestCreateFailsIfNameIsTaken() {
	name := "some-name"
	s.categoryStore.Insert(data.Category{Id: "some-id", AccountId: testAccountId, Name: name})

	reqBody := rest.CategoryCreateBody{Name: name, Color: 1011011}
	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Create(req)
	s.ErrRepsonse(res, rest.ErrCategoryNameAlreadyExists)
}

func (s *CategoryApiTestSuite) TestDeletes() {
	s.queue.Flush(queue.QueueName_CategoryDeleted)

	categoryId := "category-123"
	s.categoryStore.Insert(data.Category{Id: categoryId, AccountId: testAccountId})
	req := &rest.Request{Params: rest.PathParams{"categoryId": categoryId}}
	res := s.api.Delete(req)
	s.OkResponse(res, rest.Success().Data)

	s.Eventually(func() bool {
		msg, err := s.queue.Pop(queue.QueueName_CategoryDeleted)
		s.NoError(err)
		return (msg != nil) && (err == nil)
	}, 1*time.Minute, 1*time.Second)
}

func (s *CategoryApiTestSuite) TestDeleteFailsIfNoCategoryExists() {
	categoryId := "category-123"
	req := &rest.Request{Params: rest.PathParams{"categoryId": categoryId}}
	res := s.api.Delete(req)
	s.ErrRepsonse(res, rest.ErrInvalidCategoryId)
}
