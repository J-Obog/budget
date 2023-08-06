package server

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/J-Obog/paidoff/config"
	"github.com/stretchr/testify/suite"
)

const (
	testSvrAddress = "localhost"
	testSvrPort    = 8077
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

func (s *ServerTestSuite) TestStartsAndStops() {
	s.server.Start(testSvrAddress, testSvrPort)

	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", testSvrAddress, testSvrPort), timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
	}
	conn.Close()

	err = s.server.Stop()
	s.NoError(err)
}
