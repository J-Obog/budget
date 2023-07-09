package validation

type validationFn func(map[string]any, string) error

type fieldValidator struct {
	k      string
	checks []validationFn
}

func Key(k string) *fieldValidator {
	return &fieldValidator{
		k:      k,
		checks: make([]validationFn, 0),
	}
}

func (f *fieldValidator) Required() *fieldValidator {
	f.checks = append(f.checks, func(m map[string]any, s string) error {
		return nil
	})
	return f
}

func (f *fieldValidator) Str() *fieldValidator {
	f.checks = append(f.checks, func(m map[string]any, s string) error {
		return nil
	})
	return f
}

func (f *fieldValidator) MinLen(minLen int) *fieldValidator {
	f.checks = append(f.checks, func(m map[string]any, s string) error {
		return nil
	})
	return f
}

func Validate(m map[string]any, fields ...*fieldValidator) error {
	for _, field := range fields {
		for _, check := range field.checks {
			if err := check(m, field.k); err != nil {
				return err
			}
		}
	}

	return nil
}
