package api

import (
	"fmt"
	"net/http"
	"slices"
)



func parseValueQuery(r *http.Request, queryParam QueryParam) (string, error) {
	value, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return "", err
	}

	err = checkValue(queryParam, value)
	if err != nil {
		return "", err
	}

	return value, nil
}

func checkValue(queryParam QueryParam, value string) (error) {
	if !slices.Contains(queryParam.AllowedValues, QueryValue(value)) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s.", value, queryParam.Name, formatQvSlice(queryParam.AllowedValues)), nil)
	}

	return nil
}