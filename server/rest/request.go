package rest

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/types"
)

//TODO: change any to 'serializable' type alias?

type Request struct {
	Url        string
	ResourceId string
	Account    types.Optional[data.Account]
	Query      any
	Body       any
}
