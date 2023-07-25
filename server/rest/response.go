package rest

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Data  any
	Error error
}

func Ok(v any) *Response {
	return &Response{
		Data: v,
	}
}

func Success() *Response {
	return &Response{}
}

func Err(err error) *Response {
	return &Response{
		Error: err,
	}
}

func (res *Response) ToJSON() ([]byte, int) {
	if res.Error != nil {
		restErr, isHandledError := res.Error.(*RestError)

		if isHandledError {
			d := []byte(fmt.Sprintf(`{"error": {"message": %s}}`, restErr.Error()))
			return d, restErr.Status
		}

		return []byte(`{"error": {"message": "internal error"}}`), 500
	}

	b, err := json.Marshal(res.Data)

	if err != nil {
		return []byte(`{"error": {"message": "internal error"}}`), 500
	}

	return b, 200
}
