package rest

import "net/http"

const (
	Code_400 = http.StatusBadRequest
	Code_404 = http.StatusNotFound
	Code_500 = http.StatusInternalServerError
)

// TODO: actually implement errors
var (
	// 400 errors
	ErrInvalidJSONBody = &RestError{
		Msg:    "error parsing body",
		Status: Code_400,
	}

	ErrCategoryNameAlreadyExists = &RestError{
		Msg:    "category name already exists",
		Status: Code_400,
	}

	ErrInvalidAccountName = &RestError{
		Msg:    "invalid value for account name",
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
		Msg:    "invalid value for transaction note",
		Status: Code_400,
	}

	ErrInvalidCategoryName = &RestError{
		Msg:    "invalid value for category name",
		Status: Code_400,
	}

	//404 errors
	ErrInvalidBudgetId = &RestError{
		Msg:    "invalid budget id",
		Status: Code_404,
	}

	ErrInvalidCategoryId = &RestError{
		Msg:    "invalid category id",
		Status: Code_404,
	}

	ErrInvalidTransactionId = &RestError{
		Msg:    "invalid transaction id",
		Status: Code_404,
	}

	ErrInvalidAccountId = &RestError{
		Msg:    "invalid account id",
		Status: Code_404,
	}

	//500
	ErrInternalServer = &RestError{
		Msg:    "internal error",
		Status: Code_500,
	}
)

type RestError struct {
	Msg    string `json:"message"`
	Status int
}

func (err *RestError) Error() string {
	return err.Msg
}
