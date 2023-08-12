package api

import (
	"net/http"

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
	req := &rest.Request{}
	account := data.Account{Id: testAccountId}
	s.accountStore.Insert(account)

	res := s.api.Get(req)

	s.ResponseBodyEquals(res, account)
	s.StatusCodeEquals(res, http.StatusOK)
}

func (s *AccountApiTestSuite) TestGetFailsIfNoAccountExists() {
	req := &rest.Request{}
	res := s.api.Get(req)

	s.ResponseBodyEquals(res, rest.ErrInvalidAccountId)
	s.StatusCodeEquals(res, rest.ErrInvalidAccountId.Status)
}

func (s *AccountApiTestSuite) TestUpdates() {
	account := data.Account{Id: testAccountId, Name: "Old Name"}
	s.accountStore.Insert(account)

	var jsonb rest.JSONBody
	reqBody := rest.AccountUpdateBody{Name: "New Name"}

	err := jsonb.From(&reqBody)
	s.NoError(err)

	req := &rest.Request{Body: jsonb}

	res := s.api.Update(req)
	updatedAccount := res.Data.(data.Account)

	s.Equal(updatedAccount.Name, reqBody.Name)
	s.StatusCodeEquals(res, http.StatusOK)
}

func (s *AccountApiTestSuite) TestUpdateFailsIfNoAccountExists() {
	var jsonb rest.JSONBody
	reqBody := rest.AccountUpdateBody{Name: "New Name"}

	err := jsonb.From(&reqBody)
	s.NoError(err)

	req := &rest.Request{Body: jsonb}

	res := s.api.Update(req)
	s.ResponseBodyEquals(res, rest.ErrInvalidAccountId)
	s.StatusCodeEquals(res, rest.ErrInvalidAccountId.Status)
}

func (s *AccountApiTestSuite) TestUpdateFailsIfAccountNameIsInvalid() {
	account := data.Account{Id: testAccountId, Name: "Old Name"}
	s.accountStore.Insert(account)

	cases := []struct {
		scenario    string
		accountName string
	}{
		{"account name too short", ""},
		{"account name too long", ""},
	}

	for _, testCase := range cases {
		s.Run(testCase.scenario, func() {
			var jsonb rest.JSONBody
			reqBody := rest.AccountUpdateBody{Name: testCase.accountName}

			err := jsonb.From(&reqBody)
			s.NoError(err)

			req := &rest.Request{Body: jsonb}
			res := s.api.Update(req)

			s.ResponseBodyEquals(res, rest.ErrInvalidAccountName)
			s.StatusCodeEquals(res, rest.ErrInvalidAccountName.Status)
		})
	}
}

func (s *AccountApiTestSuite) TestDeletes() {
	account := data.Account{Id: testAccountId}
	s.accountStore.Insert(account)

	req := &rest.Request{}
	res := s.api.Delete(req)

	s.ResponseBodyEquals(res, rest.Success().Data)
	s.StatusCodeEquals(res, http.StatusOK)
}

func (s *AccountApiTestSuite) TestDeleteFailsIfNoAccountExists() {
	req := &rest.Request{}
	res := s.api.Delete(req)

	s.ResponseBodyEquals(res, rest.ErrInvalidAccountId)
	s.StatusCodeEquals(res, rest.ErrInvalidAccountId.Status)
}
