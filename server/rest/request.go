package rest

import "github.com/J-Obog/paidoff/data"

//TODO: change any to 'serializable' type alias?

type Request struct {
	Url        string
	ResourceId string
	Account    *data.Account
	Query      any
	Body       any
}
