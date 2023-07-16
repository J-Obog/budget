package queue

const (
	testQueueName = "test-queue"
)

func testMessage() Message {

	return Message{
		Id:   "123456",
		Data: "hello world",
	}
}
