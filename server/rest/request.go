package rest

import (
	"github.com/J-Obog/paidoff/data"
)

type Request struct {
	Url        string
	ResourceId string
	Account    *data.Account
	Query      map[string]string
	Body       []byte
}

// TODO: make any serializable type alias?
func ParseBody[T any](jsonb []byte) (T, error) {
	var t T
	return t, nil
}

func ParseQuery[T any](qmap map[string]string) (T, error) {
	var t T
	return t, nil
}
