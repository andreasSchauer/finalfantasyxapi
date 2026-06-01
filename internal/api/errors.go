package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

var errNotAnID = errors.New("not an id.")
var errIdNotFound = errors.New("id doesn't exit.")
var errContinue = errors.New("loop should be continued.")
var errNoResource = errors.New("no resources found.")
var errEmptyQuery = errors.New("query parameter is empty.")
var errQueryNone = errors.New("'none' input in query.")
var errQueryRedirect = errors.New("query parameter is not empty, but will be dealt with else where.")
var errNoDefaultVal = errors.New("query parameter doesn't have a default value, or default value is unused.")
var errNoSpecialInput = errors.New("query parameter doesn't have special inputs, or no special input was found.")
var errNoIntRange = errors.New("query parameter doesn't have integer range.")
var errCorrect = errors.New("test got the expected error.")
var errIgnoredField = errors.New("this field of the test struct is ignored")

func errExceptEmptyQuery(err error) bool {
	return err != nil && !queryIsEmpty(err)
}

func queryIsEmpty(err error) bool {
	return errors.Is(err, errEmptyQuery)
}

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

func newHTTPErrorDB(fetchType string, parentItem seeding.Lookupable, err error) httpError {
	return newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %ss of %s.", fetchType, parentItem), err)
}

func newHTTPErrorDbPairs(childResType, parentResType string, err error) httpError {
	return newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s + %s pairs.", childResType, parentResType), err)
}

func newHTTPErrorDbOne(fetchType string, parentItem seeding.Lookupable, err error) httpError {
	return newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %s of %s.", fetchType, parentItem), err)
}

func newHTTPErrorDbFilter(fetchType string, queryParam QueryParam, err error) httpError {
	return newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", fetchType, queryParam.Name), err)
}

func newHTTPErrorFetchLimit(fetchLimit int) httpError {
	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("fetch limit exceeded. the maximum amount of inputs is %d.", fetchLimit), nil)
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
