package api

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/suite"
)

func TestAccountApi(t *testing.T) {
	suite.Run(t, new(AccountApiTestSuite))
}

type AccountApiTestSuite struct {
	ApiTestSuite
	api *AccountAPI
}

func (s *AccountApiTestSuite) SetupSuite() {
	s.initDependencies()
	s.api = NewAccountAPI(s.accountManager)
}

func (s *AccountApiTestSuite) SetupTest() {
	err := s.accountStore.DeleteAll()
	s.NoError(err)
}

func (s *AccountApiTestSuite) TestGets() {
	s.accountStore.Insert(data.Account{Id: testAccountId})
	req := &rest.Request{}
	res := s.api.Get(req)
	s.OkResponse(res, &data.Account{})
}

func (s *AccountApiTestSuite) TestGetFailsIfNoAccountExists() {
	req := &rest.Request{}
	res := s.api.Get(req)
	s.ErrRepsonse(res, rest.ErrInvalidAccountId)
}

func (s *AccountApiTestSuite) TestUpdates() {
	s.accountStore.Insert(data.Account{Id: testAccountId})
	reqBody := rest.AccountUpdateBody{Name: "New Name"}
	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Update(req)
	s.OkResponse(res, &data.Account{})
}

func (s *AccountApiTestSuite) TestUpdateFailsIfNoAccountExists() {
	reqBody := rest.AccountUpdateBody{Name: "New Name"}
	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidAccountId)
}

func (s *AccountApiTestSuite) TestUpdateFailsIfAccountNameIsInvalid() {
	s.accountStore.Insert(data.Account{Id: testAccountId})
	invalidNames := []string{}

	for _, name := range invalidNames {
		reqBody := rest.AccountUpdateBody{Name: name}
		req := &rest.Request{Body: s.getJSONBody(reqBody)}
		res := s.api.Update(req)
		s.ErrRepsonse(res, rest.ErrInvalidAccountName)
	}
}

func (s *AccountApiTestSuite) TestDeletes() {
	s.accountStore.Insert(data.Account{Id: testAccountId})
	req := &rest.Request{}
	res := s.api.Delete(req)
	s.OkResponse(res, rest.Success().Data)
}

func (s *AccountApiTestSuite) TestDeleteFailsIfNoAccountExists() {
	req := &rest.Request{}
	res := s.api.Delete(req)
	s.ErrRepsonse(res, rest.ErrInvalidAccountId)
}
