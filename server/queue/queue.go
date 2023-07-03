package queue

import "github.com/J-Obog/paidoff/data"

type Queue interface {
	Push(message data.Message, queueName string) error
	Pop(queueName string) (*data.Message, error)
	Ack(messageId string) error
}
