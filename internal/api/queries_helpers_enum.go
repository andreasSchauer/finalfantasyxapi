package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

// validates enum-queryParam and checks emptiness.
func (q QueryParam) parseEnum[E, N any](r *http.Request, et enums.EnumType[E, N]) (enums.EnumVal, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return enums.EnumVal{}, err
	}

	return q.checkEnum(query, et)
}

// checks, if query enum is valid.
func (q QueryParam) checkEnum[E, N any](val string, et enums.EnumType[E, N]) (enums.EnumVal, error) {
	enum, err := et.CheckVal(val)
	switch err {
	case errs.ErrIDNotFound:
		return enums.EnumVal{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided id '%s' used for parameter '%s' doesn't exist. max id: %d", val, q.Name, len(et.Lookup)), nil)

	case errs.ErrNoResource:
		return enums.EnumVal{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid enum value '%s' used for parameter '%s'. use /api/%s/parameters to see allowed values", val, q.Name, q.Endpoint), nil)

	default:
		return enum, nil
	}
}
