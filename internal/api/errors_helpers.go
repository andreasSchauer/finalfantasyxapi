package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func HandleHTTPError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	httpErr, ok := err.(errs.HTTPError)
	if ok {
		respondWithError(w, httpErr.Code, httpErr.Msg, httpErr.Err)
		return true
	}

	respondWithError(w, http.StatusInternalServerError, "an unexpected error occurred", err)
	return true
}

func cascadeHTTPError(msg string, err error) error {
	if err == nil {
		return nil
	}

	httpErr, ok := err.(errs.HTTPError)
	if ok {
		httpErr.Msg = fmt.Sprintf("%s: %s", msg, httpErr.Msg)
		return httpErr
	}

	return nil
}

func fieldErrorMsg(fieldName ts.FieldName, msg string) string {
	return fmt.Sprintf("field '%s': %s", fieldName, msg)
}

func getFieldRangeErrStart(val int32, fieldName ts.FieldName, intType string, isArray bool) string {
	errStart := fmt.Sprintf("%s '%d'", intType, val)

	if !isArray {
		errStart += fmt.Sprintf(" of field '%s'", fieldName)
	}

	errStart += " is out of range"

	return errStart
}
