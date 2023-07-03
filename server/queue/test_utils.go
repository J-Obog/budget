package queue

import "github.com/J-Obog/paidoff/data"

const (
	testQueueName = "test-queue"
)

func testMessage() data.Message {

	return data.Message{
		Id:        "123456",
		Data:      "hello world",
		Timestamp: 123456,
	}
}
