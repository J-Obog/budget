package rest

import (
	"encoding/json"
	"fmt"

	"github.com/J-Obog/paidoff/data"
)

type JSONSerializable struct {
}

func (j *JSONSerializable) ToBytes() []byte {
	return
}

type AccountUpdateBody struct {
	Name string `json:"name"`
}

type CategoryUpdateBody struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type CategoryCreateBody struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type TransactionUpdateBody struct {
	CategoryId *string         `json:"categoryId"`
	Note       *string         `json:"note"`
	Type       data.BudgetType `json:"budgetType"`
	Amount     float64         `json:"amount"`
	Month      int             `json:"month"`
	Day        int             `json:"day"`
	Year       int             `json:"year"`
}

type TransactionCreateBody struct {
	CategoryId *string         `json:"categoryId"`
	Note       *string         `json:"note"`
	Type       data.BudgetType `json:"budgetType"`
	Amount     float64         `json:"amount"`
	Month      int             `json:"month"`
	Day        int             `json:"day"`
	Year       int             `json:"year"`
}

type BudgetCreateBody struct {
	CategoryId string  `json:"categoryId"`
	Projected  float64 `json:"projected"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
}

type BudgetUpdateBody struct {
	CategoryId string  `json:"categoryId"`
	Projected  float64 `json:"projected"`
}

func ParseBody[T any](jsonb []byte) (T, error) {
	var t T

	err := json.Unmarshal(jsonb, &t)

	if err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			return t, ErrInvalidJSONBody
		case *json.UnmarshalTypeError:
			return t, &RestError{Msg: fmt.Sprintf("invalid value for %s", jsonErr.Field)}
		default:
			return t, ErrInternalServer
		}
	}

	return t, nil
}
