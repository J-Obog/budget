package validation

func ValidateAccountUpdateReq(jsonb []byte) error {
	v := NewValidator()
	return v.Validate(jsonb)
}

func ValidateBudgetUpdateReq(jsonb []byte) error {
	v := NewValidator()
	return v.Validate(jsonb)
}

func ValidateBudgetCreateReq(jsonb []byte) error {
	v := NewValidator()
	return v.Validate(jsonb)
}
