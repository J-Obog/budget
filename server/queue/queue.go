package queue

type Message struct {
	Id        string `json:"id"`
	Data      []byte `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

type Queue interface {
	Push(message Message) error
	Pull() ([]Message, error)
}
