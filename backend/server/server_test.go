package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	testRoutePath       = "/foobar"
	queryParamsTestPath = "/query.params.test"
	pathParamsTestPath  = "/path.params.test"
	bodyTestPath        = "/body.test"

	testSvrAddress = "localhost"
	testSvrPort    = 8077
)

var (
	baseTestUrl = fmt.Sprintf("http://%s:%d", testSvrAddress, testSvrPort)
)

var (
	httpMethods = []string{http.MethodGet, http.MethodDelete, http.MethodPut, http.MethodPost}
)

func TestServer(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
	server Server
}

func (s *ServerTestSuite) SetupSuite() {
	cfg := config.Get()
	s.server = NewServer(cfg)
}

// TODO: check if server has been shut down properly
func (s *ServerTestSuite) TestStartsAndStops() {
	go s.server.Start(testSvrAddress, testSvrPort)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", testSvrAddress, testSvrPort))
	s.NoError(err)
	s.NoError(conn.Close())

	err = s.server.Stop()
	s.NoError(err)
}

func (s *ServerTestSuite) TestMapsParamsToRequest() {
	go s.server.Start(testSvrAddress, testSvrPort)

	expectedParams := rest.PathParams{}
	routePath := pathParamsTestPath
	url := baseTestUrl + routePath

	for i := 0; i < 10; i++ {
		pName := fmt.Sprintf("p%d", i)
		pVal := fmt.Sprintf("foobar%d", i)
		expectedParams[pName] = pVal
		routePath += "/:" + pName
		url += "/" + pVal
	}

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.MatchedBy(func(req *rest.Request) bool {
		return s.Equal(expectedParams, req.Params)
	})).Return(rest.Ok(`ok`))

	s.server.RegisterRoute(http.MethodGet, routePath, fakeHandler.Execute)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	s.NoError(err)

	_, err = http.DefaultClient.Do(req)
	s.NoError(err)

	err = s.server.Stop()
	s.NoError(err)
}

func (s *ServerTestSuite) TestMapsQueryToRequest() {
	go s.server.Start(testSvrAddress, testSvrPort)

	q1 := "foo"
	q2 := "bar"
	q3 := "baz"

	queryParams := fmt.Sprintf("&q1=%s?q2=%s?q3=%s", q1, q2, q3)

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.MatchedBy(func(req *rest.Request) bool {
		expected := rest.Query{
			"q1": {q1},
			"q2": {q2},
			"q3": {q3},
		}

		return s.Equal(expected, req.Query)
	})).Return(rest.Ok(`ok`))

	s.server.RegisterRoute(http.MethodGet, testRoutePath+"boop", fakeHandler.Execute)

	url := fmt.Sprintf("http://%s:%d%s", testSvrAddress, testSvrPort, testRoutePath+"boop"+queryParams)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	s.NoError(err)

	_, err = http.DefaultClient.Do(req)
	s.NoError(err)

	err = s.server.Stop()
	s.NoError(err)
}

func (s *ServerTestSuite) TestMapsBodyToRequest() {
	go s.server.Start(testSvrAddress, testSvrPort)

	body := []byte(`{"foo": "bar"}`)

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.MatchedBy(func(req *rest.Request) bool {
		return s.JSONEq(string(req.Body.Bytes()), string(body))
	})).Return(rest.Ok(`ok`))

	s.server.RegisterRoute(http.MethodPost, testRoutePath+"foo", fakeHandler.Execute)

	url := fmt.Sprintf("http://%s:%d%s", testSvrAddress, testSvrPort, testRoutePath+"foo")

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	s.NoError(err)

	_, err = http.DefaultClient.Do(req)
	s.NoError(err)

	err = s.server.Stop()
	s.NoError(err)
}

func (s *ServerTestSuite) TestRegistersRoutesAndGetsResponse() {
	go s.server.Start(testSvrAddress, testSvrPort)

	resp := rest.Ok(`some ok response`)

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.Anything).Return(resp)

	for _, httpMethod := range httpMethods {
		s.server.RegisterRoute(httpMethod, testRoutePath, fakeHandler.Execute)
		url := fmt.Sprintf("http://%s:%d%s", testSvrAddress, testSvrPort, testRoutePath)

		req, err := http.NewRequest(httpMethod, url, nil)
		s.NoError(err)

		res, err := http.DefaultClient.Do(req)
		s.NoError(err)
		s.Equal(res.StatusCode, resp.Status)

		b, err := io.ReadAll(res.Body)
		s.NoError(err)

		respJSONBody := rest.JSONBody{}
		err = respJSONBody.From(&resp.Data)
		s.NoError(err)

		s.JSONEq(string(b), string(respJSONBody.Bytes()))
	}

	err := s.server.Stop()
	s.NoError(err)
}
