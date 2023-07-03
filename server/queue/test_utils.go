package queue

import "github.com/J-Obog/paidoff/data"

func testMessage() data.Message {

	return data.Message{
		Id:        "123456",
		Data:      "hello world",
		Timestamp: 123456,
	}
}
