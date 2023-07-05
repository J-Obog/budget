package validation

type Validator struct {
	rules []func() error
}

func NewValidator() *Validator {
	return &Validator{
		rules: make([]func() error, 0),
	}
}

func (v *Validator) AddRule(rule func() error) {
	v.rules = append(v.rules, rule)
}

func (v *Validator) Validate() error {
	for _, rule := range v.rules {
		if err := rule(); err != nil {
			return err
		}
	}

	return nil
}
