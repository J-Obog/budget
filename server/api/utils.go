package api

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/paidoff/data"
)

func isErrorResponse(status int) bool {
	return (status == http.StatusForbidden ||
		status == http.StatusNotFound ||
		status == http.StatusInternalServerError)
}

func FromJSON[T interface{}](body []byte) (T, error) {
	var d T
	err := json.Unmarshal(body, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func ToJSON(serializable interface{}) ([]byte, error) {
	return json.Marshal(serializable)
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
