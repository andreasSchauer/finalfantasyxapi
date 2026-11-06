package api

import (
	"fmt"
	"net/http"
	"slices"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (q QueryParam) ParseValue(r *http.Request) (string, error) {
	value, err := q.CheckEmpty(r)
	if err != nil {
		return "", err
	}

	err = q.CheckValue(value)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (q QueryParam) CheckValue(value string) error {
	if !slices.Contains(q.AllowedValues, ts.QueryValue(value)) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s", value, q.Name, ts.FormatQvSlice(q.AllowedValues)), nil)
	}

	return nil
}
