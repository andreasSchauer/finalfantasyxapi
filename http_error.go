package main

import (
	"fmt"
	"net/http"
)

type httpError struct {
	code int
	msg  string
	err  error
}


func (e httpError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("http error: %s: %v", e.msg, e.err)
	}
	return fmt.Sprintf("http error: %s", e.msg)
}


func newHTTPError(code int, msg string, err error) httpError {
	return httpError{
		code: code,
		msg:  msg,
		err:  err,
	}
}


func handleHTTPError(w http.ResponseWriter, err error) bool {
    if err == nil {
        return false
    }

    if httpErr, ok := err.(httpError); ok {
        respondWithError(w, httpErr.code, httpErr.msg, httpErr.err)
        return true
    }

    respondWithError(w, http.StatusInternalServerError, "An unexpected error occurred", err)
    return true
}
