package queue

type Message struct {
	Id        string
	Data      []byte
	Timestamp int64
}

type Queue interface {
	Push(message Message) error
	Pull() ([]Message, error)
}
