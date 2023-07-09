package api

import (
	"net/http"

	"github.com/J-Obog/paidoff/rest"
)

func buildBadRequestError() *rest.Response {
	return &rest.Response{
		Status: http.StatusBadRequest,
	}
}

func buildServerError(err error) *rest.Response {
	return &rest.Response{
		Status: http.StatusInternalServerError,
	}
}

func buildOKResponse(d interface{}) *rest.Response {
	return &rest.Response{
		Status: http.StatusOK,
		Data:   d,
	}
}
