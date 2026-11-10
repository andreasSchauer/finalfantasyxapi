package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
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
