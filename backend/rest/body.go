package rest

import (
	"encoding/json"
	"fmt"
)

type JSONBody []byte

func (j JSONBody) getRestError(err error) *RestError {
	switch jsonErr := err.(type) {
	case *json.SyntaxError:
		return ErrInvalidJSONBody
	case *json.UnmarshalTypeError:
		return &RestError{Msg: fmt.Sprintf("invalid value for %s", jsonErr.Field)}
	default:
		return ErrInternalServer
	}
}

func (j JSONBody) From(obj any) error {
	bytes, err := json.Marshal(obj)

	if err != nil {
		return j.getRestError(err)
	}

	j = bytes

	return nil
}

func (j JSONBody) To(obj any) error {
	err := json.Unmarshal(j, obj)

	if err != nil {
		return j.getRestError(err)
	}

	return nil
}
