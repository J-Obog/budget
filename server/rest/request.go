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
