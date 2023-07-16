package rest

import "github.com/J-Obog/paidoff/data"

//TODO: change any to 'serializable' type alias?

type Request struct {
	Url     string
	Account *data.Account
	Params  Params
	Query   any
	Body    any
}
