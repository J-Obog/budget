package api

import (
	"errors"

	"github.com/J-Obog/paidoff/data"
)

const (
	accountIdHeaderKey = "Authorization"
)

func getCurrentAccountId(req *data.RestRequest) (error, string) {
	val, ok := req.Headers[accountIdHeaderKey].(string)

	if ok {
		return nil, val
	}

	return errors.New("something went wrong"), ""
}
