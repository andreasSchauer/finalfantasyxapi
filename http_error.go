package main

import "fmt"

type httpError struct {
	code		int
	msg			string
	err			error
}

func (e httpError) Error() string {
    if e.err != nil {
        return fmt.Sprintf("http error: %s: %v", e.msg, e.err)
    }
    return fmt.Sprintf("http error: %s", e.msg)
}

func NewHTTPError(code int, msg string, err error) httpError {
    return httpError{
        code: code,
        msg:  msg,
        err:  err,
    }
}