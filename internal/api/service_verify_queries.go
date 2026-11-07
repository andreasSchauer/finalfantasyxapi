package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

// VerifyQueryParamsService verifies, if no query params are in use for an endpoint that receives a post request
func VerifyQueryParamsService(r *http.Request) error {
	if len(r.URL.Query()) > 0 {
		return errs.NewHTTPError(http.StatusBadRequest, "computational endpoints don't have any query parameters", nil)
	}

	return nil
}