package api

import (
	"net/http"

	"github.com/J-Obog/paidoff/data"
)

func buildServerError(err error) *data.RestResponse {
	return &data.RestResponse{
		Status: http.StatusInternalServerError,
	}
}

func buildNotFoundError() *data.RestResponse {
	return &data.RestResponse{
		Status: http.StatusNotFound,
	}
}

func buildForbiddenError() *data.RestResponse {
	return &data.RestResponse{
		Status: http.StatusForbidden,
	}
}

func buildOKResponse(d interface{}) *data.RestResponse {
	return &data.RestResponse{
		Status: http.StatusOK,
		Data:   d,
	}
}
