package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/J-Obog/paidoff/data"
)

func isErrorResponse(status int) bool {
	return (status == http.StatusForbidden ||
		status == http.StatusNotFound ||
		status == http.StatusInternalServerError)
}

func ICast[T interface{}](v interface{}) (T, error) {
	val, ok := v.(T)

	if !ok {
		return val, errors.New("failed to convert")
	}

	return val, nil
}

func FromJSON[T interface{}](body []byte) (T, error) {
	var d T
	err := json.Unmarshal(body, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func buildServerError(res *data.RestResponse, err error) {
	res.Status = http.StatusInternalServerError
}

func buildNotFoundError(res *data.RestResponse) {
	res.Status = http.StatusNotFound
}

func buildForbiddenError(res *data.RestResponse) {
	res.Status = http.StatusForbidden
}

func buildOKResponse(res *data.RestResponse, d interface{}) {
	res.Status = http.StatusOK
	res.Data = d
}
