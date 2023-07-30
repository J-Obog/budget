package rest

import (
	"encoding/json"
	"fmt"

	"github.com/J-Obog/paidoff/data"
	"github.com/gorilla/schema"
)

type Request struct {
	Url        string
	ResourceId string
	Account    *data.Account
	Query      map[string][]string
	Params     map[string]string
	Body       []byte
}

func ParseBody[T any](jsonb []byte) (T, error) {
	var t T

	err := json.Unmarshal(jsonb, &t)

	if err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			return t, ErrInvalidJSONBody
		case *json.UnmarshalTypeError:
			return t, &RestError{Msg: fmt.Sprintf("invalid value for %s", jsonErr.Field)}
		default:
			return t, ErrInternalServer
		}
	}

	return t, nil
}

func ParseQuery[T any](queryMap map[string][]string) (T, error) {
	var t T

	decoder := schema.NewDecoder()

	if err := decoder.Decode(&t, queryMap); err != nil {
		return t, &RestError{Msg: "invalid value for parsing query"}
	}

	return t, nil
}
