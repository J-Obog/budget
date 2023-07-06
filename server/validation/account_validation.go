package validation

func ValidateAccountUpdateReq(m map[string]interface{}) error {
	v := NewValidator()
	v.Field("name", Required(), As[string](), Length(0, 1))

	return v.Validate(m)
}
