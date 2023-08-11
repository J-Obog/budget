package rest

import (
	"encoding/json"
	"fmt"
)

type JSONBody []byte

func (j JSONBody) From() error {
	return nil
}

func (j JSONBody) To(obj any) error {
	err := json.Unmarshal(j, obj)

	if err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			return ErrInvalidJSONBody
		case *json.UnmarshalTypeError:
			return &RestError{Msg: fmt.Sprintf("invalid value for %s", jsonErr.Field)}
		default:
			return ErrInternalServer
		}
	}

	return nil
}
