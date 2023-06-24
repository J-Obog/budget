package queue

import "github.com/J-Obog/paidoff/data"

type Queue interface {
	Push(serializable interface{}) error
	Pull() ([]data.Message, error)
}
