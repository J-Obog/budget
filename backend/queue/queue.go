package queue

const (
	QueueName_CategoryDeleted string = "category.deleted"
)

type Queue interface {
	Push(message Message, queueName string) error
	Pop(queueName string) (*Message, error)
	Ack(messageId string) error
	Flush(queueName string) error
}
