package errs

import "errors"

func ErrExceptEmptyQuery(err error) bool {
	return err != nil && !QueryIsEmpty(err)
}

func QueryIsEmpty(err error) bool {
	return errors.Is(err, ErrEmptyQuery)
}
