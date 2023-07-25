package rest

import "net/http"

const (
	Code_400 = http.StatusBadRequest
	Code_404 = http.StatusNotFound
)

// TODO: actually implement errors
var (
	// 400 errors
	ErrCategoryNameAlreadyExists = &RestError{
		Msg:    "category name already exists",
		Status: Code_400,
	}

	ErrInvalidAccountName = &RestError{
		Msg:    "invalid value for name",
		Status: Code_400,
	}

	ErrCategoryAlreadyInBudgetPeriod = &RestError{
		Msg:    "category has already been assigned to a budget",
		Status: Code_400,
	}

	ErrInvalidDate = &RestError{
		Msg:    "invalid date",
		Status: Code_400,
	}

	ErrCategoryCurrentlyInUse = &RestError{
		Msg:    "category is currently being referenced",
		Status: Code_400,
	}

	ErrInvalidTransactionNote = &RestError{
		Msg:    "invalid value for note",
		Status: Code_400,
	}

	ErrInvalidCategoryName = &RestError{
		Msg:    "invalid value for name",
		Status: Code_400,
	}

	//404 errors
	ErrInvalidBudgetId = &RestError{
		Msg:    "category name already exists",
		Status: Code_404,
	}

	ErrInvalidCategoryId = &RestError{
		Msg:    "category name already exists",
		Status: Code_404,
	}

	ErrInvalidTransactionId = &RestError{
		Msg:    "category name already exists",
		Status: Code_404,
	}
)

type RestError struct {
	Msg    string
	Status int
}

func (err *RestError) Error() string {
	return err.Msg
}
