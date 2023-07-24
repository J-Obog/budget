package rest

import (
	"errors"
)

// TODO: actually implement errors
var (
	ErrCategoryNameAlreadyExists     = errors.New("")
	ErrBadRequest                    = errors.New("")
	ErrInvalidAccountName            = errors.New("")
	ErrCategoryAlreadyInBudgetPeriod = errors.New("")
	ErrInvalidDate                   = errors.New("")
	ErrCategoryCurrentlyInUse        = errors.New("")
	ErrInvalidTransactionNote        = errors.New("")
	ErrInvalidBudgetId               = errors.New("")
	ErrInvalidCategoryId             = errors.New("")
	ErrInvalidTransactionId          = errors.New("")
	ErrInvalidCategoryName           = errors.New("")
)

type Response struct {
	Data  any
	Error error
}

func Ok(v any) *Response {
	return &Response{
		Data: v,
	}
}

func Success() *Response {
	return &Response{}
}

func Err(err error) *Response {
	return &Response{
		Error: err,
	}
}
