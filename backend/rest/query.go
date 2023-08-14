package rest

import (
	"reflect"

	"github.com/J-Obog/paidoff/data"
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
	decoder.RegisterConverter(data.Date{}, func(s string) reflect.Value {
		d, err := data.NewDateFromString(s)
		if err != nil {
			return reflect.Value{}
		}

		return reflect.ValueOf(d)
	})

	if err := decoder.Decode(obj, q); err != nil {
		return &RestError{Msg: "invalid value for parsing query"}
	}

	return nil
}
