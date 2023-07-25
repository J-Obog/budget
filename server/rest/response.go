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

func getErrBody(errMsg string) []byte {
	return []byte(fmt.Sprintf(`{"error": {"message": %s}}`, errMsg))
}

func (res *Response) ToJSON() ([]byte, int) {
	if res.Error != nil {
		restErr, isHandledError := res.Error.(*RestError)

		if isHandledError {
			return getErrBody(restErr.Error()), restErr.Status
		}

		return getErrBody(ErrInternalServor.Error()), ErrInternalServor.Status
	}

	b, err := json.Marshal(res.Data)

	if err != nil {
		return getErrBody(ErrInternalServor.Error()), ErrInternalServor.Status
	}

	return b, 200
}
