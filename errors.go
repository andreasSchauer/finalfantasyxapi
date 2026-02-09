package main


import (
	"errors"
	"fmt"
	"net/http"
)

var errNotAnID = errors.New("not an id.")
var errNoResource = errors.New("no resources found.")
var errEmptyQuery = errors.New("query parameter is empty.")
var errNoDefaultVal = errors.New("query parameter doesn't have a default value, or default value is unused.")
var errNoSpecialInput = errors.New("query parameter doesn't have special inputs, or no special input was found.")
var errNoIntRange = errors.New("query parameter doesn't have integer range.")
var errCorrect = errors.New("test got the expected error.")
var errIgnoredField = errors.New("this field of the test struct is ignored")

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

	respondWithError(w, http.StatusInternalServerError, "an unexpected error occurred.", err)
	return true
}
