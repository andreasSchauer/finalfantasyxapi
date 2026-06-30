package api

import (
	"fmt"
	"net/http"
)

// validates enum-queryParam and checks emptiness.
func parseEnumQuery[E, N any](r *http.Request, endpoint EndpointName, queryParam QueryParam, et EnumType[E, N]) (EnumVal, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return EnumVal{}, err
	}

	return checkQueryEnum(query, endpoint, queryParam, et)
}

// checks, if query enum is valid.
func checkQueryEnum[E, N any](val string, endpoint EndpointName, queryParam QueryParam, et EnumType[E, N]) (EnumVal, error) {
	enum, err := CheckEnumVal(val, et)
	switch err {
	case errIdNotFound:
		return EnumVal{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided id '%s' used for parameter '%s' doesn't exist. max id: %d.", val, queryParam.Name, len(et.lookup)), nil)

	case errNoResource:
		return EnumVal{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid enum value '%s' used for parameter '%s'. use /api/%s/parameters to see allowed values.", val, queryParam.Name, endpoint), nil)

	default:
		return enum, nil
	}
}
