package queue

import (
	"fmt"
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/stretchr/testify/suite"
)

const (
	testQueueName = "test-queue"
)

func TestQueue(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}

type QueueTestSuite struct {
	suite.Suite
	queue Queue
}

func (s *QueueTestSuite) SetupSuite() {
	cfg := config.Get()
	fmt.Println(cfg)
	s.queue = NewQueue(cfg)
}

func (s *QueueTestSuite) SetupTest() {
	err := s.queue.Flush(testQueueName)
	s.NoError(err)
}

func (s *QueueTestSuite) TestPushAndPop() {
	msg := Message{Id: "some-id", Body: []byte(`some payload`)}

	err := s.queue.Push(msg, testQueueName)
	s.NoError(err)

	m, err := s.queue.Pop(testQueueName)
	s.NoError(err)
	s.Equal(msg, *m)
}

func (s *QueueTestSuite) TestAck() {
	msg := Message{Id: "some-id", Body: []byte(`some payload`)}

	err := s.queue.Push(msg, testQueueName)
	s.NoError(err)

	_, err = s.queue.Pop(testQueueName)
	s.NoError(err)

	err = s.queue.Ack(msg.Id)
	s.NoError(err)

	m, err := s.queue.Pop(testQueueName)
	s.NoError(err)
	s.Nil(m)
}
