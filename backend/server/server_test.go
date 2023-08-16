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
	testSvrAddress      = "localhost"
	testSvrPort         = 8077
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

func (s *ServerTestSuite) SetupTest() {
	go s.server.Start(testSvrAddress, testSvrPort)
}

func (s *ServerTestSuite) TearDownTest() {
	s.NoError(s.server.Stop())
}

func (s *ServerTestSuite) clientDo(method string, url string, body []byte) *http.Response {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	s.NoError(err)

	res, err := http.DefaultClient.Do(req)
	s.NoError(err)

	return res
}

// TODO: check if server has been shut down properly
func (s *ServerTestSuite) TestStarts() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", testSvrAddress, testSvrPort))
	s.NoError(err)
	s.NoError(conn.Close())
}

func (s *ServerTestSuite) TestMapsParamsToRequest() {
	expected := rest.PathParams{"p1": "foo", "p2": "bar", "p3": "baz"}
	routePath := pathParamsTestPath
	url := baseTestUrl + routePath

	for k, v := range expected {
		routePath += "/:" + k
		url += "/" + v
	}

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.MatchedBy(func(req *rest.Request) bool {
		return s.Equal(expected, req.Params)
	})).Return(rest.Ok(`ok`))

	s.server.RegisterRoute(http.MethodGet, routePath, fakeHandler.Execute)

	s.clientDo(http.MethodGet, url, nil)
}

func (s *ServerTestSuite) TestMapsQueryToRequest() {
	expected := rest.Query{"q1": {"foo"}, "q2": {"bar"}, "q3": {"baz"}}
	routePath := queryParamsTestPath
	url := baseTestUrl + routePath

	i := 0
	for k, v := range expected {
		if i == 0 {
			url += fmt.Sprintf("&%s=%s", k, v[0])
		} else {
			url += fmt.Sprintf("?%s=%s", k, v[0])
		}
		i++
	}

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.MatchedBy(func(req *rest.Request) bool {
		return s.Equal(expected, req.Query)
	})).Return(rest.Ok(`ok`))

	s.server.RegisterRoute(http.MethodGet, routePath, fakeHandler.Execute)

	s.clientDo(http.MethodGet, url, nil)
}

func (s *ServerTestSuite) TestMapsBodyToRequest() {
	body := []byte(`{"key1": "foo", "key2": "bar", "key3": "baz"}`)

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.MatchedBy(func(req *rest.Request) bool {
		return s.JSONEq(string(req.Body.Bytes()), string(body))
	})).Return(rest.Ok(`ok`))

	s.server.RegisterRoute(http.MethodPost, bodyTestPath, fakeHandler.Execute)

	s.clientDo(http.MethodGet, baseTestUrl+bodyTestPath, body)
}

func (s *ServerTestSuite) TestRegistersRoutesAndGetsResponse() {
	resp := rest.Ok(`some ok response`)

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.EXPECT().Execute(mock.Anything).Return(resp)

	for _, httpMethod := range httpMethods {
		s.server.RegisterRoute(httpMethod, testRoutePath, fakeHandler.Execute)
		url := fmt.Sprintf("http://%s:%d%s", testSvrAddress, testSvrPort, testRoutePath)

		res := s.clientDo(httpMethod, url, nil)

		s.Equal(res.StatusCode, resp.Status)

		b, err := io.ReadAll(res.Body)
		s.NoError(err)

		respJSONBody := rest.JSONBody{}
		err = respJSONBody.From(&resp.Data)
		s.NoError(err)

		s.JSONEq(string(b), string(respJSONBody.Bytes()))
	}
}
