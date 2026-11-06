package errs

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code int
	Msg  string
	Err  error
}

var _ error = HTTPError{}

func (e HTTPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("http error: %s: %v", e.Msg, e.Err)
	}
	return fmt.Sprintf("http error: %s", e.Msg)
}

func NewHTTPError(code int, msg string, err error) HTTPError {
	return HTTPError{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

func NewHTTPErrorFetchLimit(fetchLimit int) HTTPError {
	return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("fetch limit exceeded. the maximum amount of inputs is %d", fetchLimit), nil)
}