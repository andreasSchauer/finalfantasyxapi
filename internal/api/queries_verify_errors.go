package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func errSingleResParam(paramName ts.QueryParamName) error {
	return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with single-resource-endpoints", paramName, paramName), nil)
}

func errListResParam(paramName ts.QueryParamName) error {
	return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with list-endpoints", paramName, paramName), nil)
}
