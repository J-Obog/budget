package validation

import "errors"

type fieldValidation func(field interface{}) error

type check struct {
	key string
	fn  fieldValidation
}

type Validator struct {
	checks []check
}

func NewValidator() *Validator {
	return &Validator{
		checks: make([]check, 0),
	}
}

func (v *Validator) Field(fieldName string, validations ...fieldValidation) {
	for _, fn := range validations {
		v.checks = append(v.checks, check{
			key: fieldName,
			fn:  fn,
		})
	}
}

func (v *Validator) Validate(m map[string]interface{}) error {
	for _, check := range v.checks {
		if err := check.fn(m[check.key]); err != nil {
			return err
		}
	}

	return nil
}

func Required() fieldValidation {
	return func(field interface{}) error {
		if field == nil {
			return errors.New("some required error")
		}

		return nil
	}
}

func As[T interface{}]() fieldValidation {
	return func(field interface{}) error {
		if field != nil {
			if _, ok := field.(T); !ok {
				return errors.New("invalid type error")
			}

			return nil
		}
		return nil
	}
}

func Length(min int, max int) fieldValidation {
	return func(field interface{}) error {
		if field != nil {
			s := field.(string)

			if len(s) < min {
				return errors.New("some min error")
			}

			if len(s) > max {
				return errors.New("some max error")
			}

			return nil
		}
		return nil
	}
}
