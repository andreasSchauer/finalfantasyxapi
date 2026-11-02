package api

import (
	"fmt"
	"net/http"
)

// validates enum-queryParam and checks emptiness.
func (q QueryParam) parseEnum[E, N any](r *http.Request, et EnumType[E, N]) (EnumVal, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return EnumVal{}, err
	}

	return q.checkEnum(query, et)
}

// checks, if query enum is valid.
func (q QueryParam) checkEnum[E, N any](val string, et EnumType[E, N]) (EnumVal, error) {
	enum, err := et.CheckVal(val)
	switch err {
	case errIDNotFound:
		return EnumVal{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided id '%s' used for parameter '%s' doesn't exist. max id: %d", val, q.Name, len(et.lookup)), nil)

	case errNoResource:
		return EnumVal{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid enum value '%s' used for parameter '%s'. use /api/%s/parameters to see allowed values", val, q.Name, q.Endpoint), nil)

	default:
		return enum, nil
	}
}
