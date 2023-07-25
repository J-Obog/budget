package rest

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
