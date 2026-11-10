package queryparam

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
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
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s", value, q.Name, ts.FormatQvSlice(q.AllowedValues)), nil)
	}

	return nil
}
