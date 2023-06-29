package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/J-Obog/paidoff/data"
)

func ResponseIsError(status int, err error) bool {
	return (err != nil ||
		status == http.StatusForbidden ||
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

func buildServerError(err error) data.RestResponse {
	return data.RestResponse{
		Status: http.StatusInternalServerError,
	}
}

func buildNotFoundError() data.RestResponse {
	return data.RestResponse{
		Status: http.StatusNotFound,
	}
}

func buildOKResponse(d interface{}) data.RestResponse {
	return data.RestResponse{
		Status: http.StatusOK,
		Data:   d,
	}
}

func getAccount(req *data.RestRequest) data.Account {
	return req.Meta["curr_account"].(data.Account)
}
