package api

import (
	"fmt"
	"net/http"
)

func errSingleResParam(paramName QueryParamName) error {
	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with single-resource-endpoints.", paramName, paramName), nil)
}

func errListResParam(paramName QueryParamName) error {
	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with list-endpoints.", paramName, paramName), nil)
}


