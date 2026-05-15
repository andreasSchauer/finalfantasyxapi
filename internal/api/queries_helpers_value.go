package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
	if !slices.Contains(queryParam.AllowedValues, value) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s.", value, queryParam.Name, h.FormatStringSlice(queryParam.AllowedValues)), nil)
	}

	return nil
}