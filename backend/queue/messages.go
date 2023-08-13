package queue

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Id   string `json:"id"`
	Body []byte `json:"body"`
}

type CategoryDeletedMessage struct {
	AccountId  string `json:"accountId"`
	CategoryId string `json:"categoryId"`
}

// TODO: handle unmarshal error?
func ToMessage(id string, obj any) Message {
	bytes, err := json.Marshal(obj)

	if err != nil {
		fmt.Println(err)
	}

	return Message{
		Id:   id,
		Body: bytes,
	}
}
