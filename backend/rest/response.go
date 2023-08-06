package rest

import "net/http"

type Response struct {
	Data           any
	Status         int
	InternalErrMsg string
}

func Ok(v any) *Response {
	return &Response{
		Data:   v,
		Status: http.StatusOK,
	}
}

func Success() *Response {
	return &Response{}
}

func Err(err error) *Response {
	res := new(Response)

	restErr, ok := err.(*RestError)

	if ok {
		res.Data = restErr
		res.Status = restErr.Status
	} else {
		res.Data = ErrInternalServer
		res.Status = restErr.Status
		res.InternalErrMsg = err.Error()
	}

	return res
}
