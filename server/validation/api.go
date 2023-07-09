package validation

import "math"

func ValidateAccountUpdateReq(jsonb []byte) error {
	v := NewValidator()
	v.Field("name", Required(), As[string](), Length(1, math.MaxInt))
	return v.Validate(jsonb)
}

func ValidateBudgetCreateReq(jsonb []byte) error {
	v := NewValidator()
	v.Field("categoryId", Required(), As[string]())
	v.Field("name", Required(), As[string](), Length(1, 255))
	v.Field("month", Required(), As[int]())
	v.Field("year", Required(), As[int]())
	v.Field("projected", Required(), As[int]())
	return v.Validate(jsonb)
}

func ValidateBudgetUpdateReq(jsonb []byte) error {
	return ValidateBudgetCreateReq(jsonb)
}

func ValidateTransactionCreateReq(jsonb []byte) error {
	v := NewValidator()
	v.Field("categoryId", Required(), As[string]())
	v.Field("name", Required(), As[string](), Length(1, 255))
	v.Field("month", Required(), As[int]())
	v.Field("year", Required(), As[int]())
	v.Field("projected", Required(), As[int]())
	return v.Validate(jsonb)
}

func ValidateTransactionUpdateReq(jsonb []byte) error {
	v := NewValidator()
	return v.Validate(jsonb)
}
