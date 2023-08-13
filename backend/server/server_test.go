package server

import (
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
	testRoutePath  = "/test/foobar"
	testSvrAddress = "localhost"
	testSvrPort    = 8077
)

var (
	httpMethods = []string{http.MethodGet, http.MethodDelete, http.MethodPut, http.MethodPost}
)

func TestServer(t *testing.T) {
	//suite.Run(t, new(ServerTestSuite))
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
	s.server.Start(testSvrAddress, testSvrPort)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", testSvrAddress, testSvrPort))
	s.NoError(err)
	s.NoError(conn.Close())

	err = s.server.Stop()
	s.NoError(err)
}

func (s *ServerTestSuite) TestRegistersRoutesAndGetsResponse() {
	s.server.Start(testSvrAddress, testSvrPort)

	resp := rest.Ok(`some ok response`)

	fakeHandler := new(mocks.RouteHandler)
	fakeHandler.On("Execute", mock.Anything).Return(resp)

	for _, httpMethod := range httpMethods {
		s.server.RegisterRoute(httpMethod, testRoutePath, fakeHandler.Execute)
		url := fmt.Sprintf("http://%s:%d%s", testSvrAddress, testSvrPort, testRoutePath)

		req, err := http.NewRequest(httpMethod, url, nil)
		s.NoError(err)

		res, err := http.DefaultClient.Do(req)
		s.NoError(err)
		s.Equal(res.StatusCode, resp.Status)

		b, err := io.ReadAll(res.Body)
		fmt.Println(string(b))
		s.NoError(err)
		//s.Equal(jsonb, b)
	}

	err := s.server.Stop()
	s.NoError(err)
}
