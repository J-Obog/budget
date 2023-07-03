package queue

import "github.com/J-Obog/paidoff/data"

type Queue interface {
	Push(message data.Message) error
	Pop() (*data.Message, error)
}
