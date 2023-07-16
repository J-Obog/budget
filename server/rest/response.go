package rest

import "net/http"

type Response struct {
	Data        any
	Status      int
	InternalErr error
}

const (
	statusOk          = http.StatusOK
	statusBadReq      = http.StatusBadRequest
	statusInternalErr = http.StatusInternalServerError
)

func (r *Response) IsErr() bool {
	return r.Status == statusBadReq || r.Status == statusInternalErr
}

func (r *Response) Ok(v any) {
	r.Data = v
	r.Status = http.StatusOK
}

func (r *Response) ErrInternal(err error) {
	r.Status = statusInternalErr
	r.InternalErr = err
}

func (r *Response) ErrAccountNotFound() {
	r.Status = statusBadReq
}

func (r *Response) ErrCategoryNotFound() {
	r.Status = statusBadReq
}

func (r *Response) ErrBudgetNotFound() {
	r.Status = statusBadReq
}

func (r *Response) ErrTransactionNotFound() {
	r.Status = statusBadReq
}

func (r *Response) ErrCategoryInBudgetPeriod() {
	r.Status = statusBadReq
}

func (r *Response) ErrBadRequest() {
	r.Status = statusBadReq
}
