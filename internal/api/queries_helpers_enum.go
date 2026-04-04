package api

import (
	"fmt"
	"net/http"
)

// validates enum-queryParam and checks emptiness.
func parseEnumQuery[E, N any](r *http.Request, endpoint string, queryParam QueryType, et EnumType[E, N]) (EnumAPIResource, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return EnumAPIResource{}, err
	}

	return checkQueryEnum(query, endpoint, queryParam, et)
}

// checks, if query enum is valid.
func checkQueryEnum[E, N any](val, endpoint string, queryParam QueryType, et EnumType[E, N]) (EnumAPIResource, error) {
	enum, err := GetEnumAPIResource(val, et)
	switch err {
	case errIdNotFound:
		return EnumAPIResource{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided id '%s' used for parameter '%s' doesn't exist. max id: %d.", val, queryParam.Name, len(et.lookup)), nil)

	case errNoResource:
		return EnumAPIResource{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid enum value '%s' used for parameter '%s'. use /api/%s/parameters to see allowed values.", val, queryParam.Name, endpoint), nil)

	default:
		return enum, nil
	}
}
