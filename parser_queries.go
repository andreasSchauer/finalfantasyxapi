package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// used, if a queryParam only takes existing ids and returns a valid id
func parseIDOnlyQuery(r *http.Request, queryParam QueryType, maxID int) (int32, error) {
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return 0, errEmptyQuery
	}

	id, err := parseQueryIdVal(query, queryParam, maxID)
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// used for boolean queryParams
func parseBooleanQuery(r *http.Request, queryParam QueryType) (bool, error) {
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return false, errEmptyQuery
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid boolean value. usage: %s", queryParam.Usage), err)
	}

	return b, nil
}

// used, if a queryParam is looking up an enum entry
func parseTypeQuery(r *http.Request, queryParam QueryType, lookup map[string]TypedAPIResource) (TypedAPIResource, error) {
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return TypedAPIResource{}, errEmptyQuery
	}

	enum, err := GetEnumType(query, lookup)
	if err != nil {
		return TypedAPIResource{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: '%s', use /api/%s to see valid values", query, queryParam.Name), err)
	}

	return enum, nil
}
