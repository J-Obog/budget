package rest

import (
	"encoding/json"
	"fmt"
)

type JSONBody struct {
	bytes []byte
}

func NewJSONBody(b []byte) *JSONBody {
	return &JSONBody{bytes: b}
}

func (j *JSONBody) getRestError(err error) *RestError {
	switch jsonErr := err.(type) {
	case *json.SyntaxError:
		return ErrInvalidJSONBody
	case *json.UnmarshalTypeError:
		return &RestError{Msg: fmt.Sprintf("invalid value for %s", jsonErr.Field)}
	default:
		return ErrInternalServer
	}
}
func (j *JSONBody) Bytes() []byte {
	return j.bytes
}

func (j *JSONBody) From(obj any) error {
	bytes, err := json.Marshal(obj)

	if err != nil {
		return j.getRestError(err)
	}

	j.bytes = bytes

	return nil
}

func (j *JSONBody) To(obj any) error {
	err := json.Unmarshal(j.bytes, obj)

	if err != nil {
		return j.getRestError(err)
	}

	return nil
}
