package api

import (
	"fmt"
	"net/http"
	"slices"
)

func (q QueryParam) parseValue(r *http.Request) (string, error) {
	value, err := q.checkEmpty(r)
	if err != nil {
		return "", err
	}

	err = q.checkValue(value)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (q QueryParam) checkValue(value string) error {
	if !slices.Contains(q.AllowedValues, QueryValue(value)) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s", value, q.Name, formatQvSlice(q.AllowedValues)), nil)
	}

	return nil
}
