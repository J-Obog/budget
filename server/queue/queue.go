package queue

type Queue interface {
	Push(message Message, queueName string) error
	Pop(queueName string) (*Message, error)
	Ack(messageId string) error
}
