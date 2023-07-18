package rest

import (
	"errors"
)

// TODO: actually implement errors
var (
	ErrCategoryAlreadyInBudgetPeriod = errors.New("")
	ErrCategoryNameAlreadyExists     = errors.New("")
	ErrBadRequest                    = errors.New("")
	ErrInvalidAccountName            = errors.New("")
	ErrInvalidDate                   = errors.New("")
	ErrCategoryCurrentlyInUse        = errors.New("")
	ErrTransactionNote               = errors.New("")
	ErrInvalidBudgetId               = errors.New("")
	ErrInvalidCategoryId             = errors.New("")
	ErrInvalidTransactionId          = errors.New("")
)

type Response struct {
	Data        any
	Status      int
	InternalErr error
}

func Ok(v any) *Response {
	return &Response{}
}

func Success() *Response {
	return &Response{}
}

func Err(err error) *Response {
	return &Response{}
}
