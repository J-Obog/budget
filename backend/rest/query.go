package rest

import (
	"github.com/gorilla/schema"
)

type Query map[string][]string

func (q Query) From(obj any) error {
	encoder := schema.NewEncoder()

	if err := encoder.Encode(obj, q); err != nil {
		return err
	}

	return nil
}

func (q Query) To(obj any) error {
	decoder := schema.NewDecoder()

	if err := decoder.Decode(obj, q); err != nil {
		return &RestError{Msg: "invalid value for parsing query"}
	}

	return nil
}
