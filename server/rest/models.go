package rest

import "github.com/J-Obog/paidoff/data"

type Request struct {
	Url     string
	Account *data.Account
	Params  Params
	Query   Query
	Body    JsonBody
}

type Response struct {
	Data   any
	Status int
}

type RestError struct {
}
