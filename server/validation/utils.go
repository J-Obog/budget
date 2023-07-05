package validation

import "errors"

func StrLen(str string, minLen int, maxLen int) func() error {
	return func() error {
		ln := len(str)

		//errors.New()

		if ln < minLen {
			return errors.New("min len error")
		}

		if (maxLen != -1) && (ln > maxLen) {
			return errors.New("max len error")
		}

		return nil
	}
}
