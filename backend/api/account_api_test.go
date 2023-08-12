package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
)

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
	s.OkResponse(res, data.Account{})
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
	s.OkResponse(res, data.Account{})
}

func (s *AccountApiTestSuite) TestUpdateFailsIfNoAccountExists() {
	reqBody := rest.AccountUpdateBody{Name: "New Name"}
	req := &rest.Request{Body: s.getJSONBody(reqBody)}
	res := s.api.Update(req)
	s.ErrRepsonse(res, rest.ErrInvalidAccountId)
}

func (s *AccountApiTestSuite) TestUpdateFailsIfAccountNameIsInvalid() {
	s.accountStore.Insert(data.Account{Id: testAccountId})

	cases := []struct {
		scenario    string
		accountName string
	}{
		{"account name too short", ""},
		{"account name too long", ""},
	}

	for _, testCase := range cases {
		s.Run(testCase.scenario, func() {
			reqBody := rest.AccountUpdateBody{Name: testCase.accountName}
			req := &rest.Request{Body: s.getJSONBody(reqBody)}
			res := s.api.Update(req)
			s.ErrRepsonse(res, rest.ErrInvalidAccountName)
		})
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
