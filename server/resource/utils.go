package resource

import "net/http"

func mustGetAccountId(req Request) string {
	return req.Meta["accountId"].(string)
}

func serverErrorResponse() *Response {
	return &Response{
		Body:   []byte("Internal server error"),
		Status: http.StatusInternalServerError,
	}
}
