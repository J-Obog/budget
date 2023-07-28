package rest

import (
	"encoding/json"
	"fmt"

	"github.com/J-Obog/paidoff/data"
)

type Request struct {
	Url        string
	ResourceId string
	Account    *data.Account
	Query      map[string][]string
	Params     map[string]string
	Body       []byte
}

// TODO: make any serializable type alias?
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

func ParseQuery[T any](qmap map[string][]string) (T, error) {
	var t T
	return t, nil
}
