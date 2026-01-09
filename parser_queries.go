package main

import (
	"errors"
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


// checks for default values, special values, validity, and range validity of an integer-based non-id query. if the query doesn't use defaults, special vals, or ranges, they are simply ignored.
func parseIntQuery(r *http.Request, queryParam QueryType) (int, error) {
	query := r.URL.Query().Get(queryParam.Name)
	
	defaultVal, err := checkDefaultVal(queryParam, query)
	if errors.Is(err, errEmptyQuery) {
		return 0, errEmptyQuery
	}
	if !errors.Is(err, errNoDefaultVal) {
		return defaultVal, nil
	}

	specialVal, err := checkQuerySpecialVals(queryParam, query)
	if !errors.Is(err, errNoSpecialInput) {
		return specialVal, nil
	}

	val, err := checkQueryIntRange(queryParam, query)
	if err != nil && !errors.Is(err, errNoIntRange) {
		return 0, err
	}

	return val, nil
}


// converts a list of ids into a slice and checks every id's validity
func parseIdListQuery(r *http.Request, queryParam QueryType, maxID int) ([]int32, error) {
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return nil, errEmptyQuery
	}

	ids, err := queryIDsToSlice(query, queryParam, maxID)
	if err != nil {
		return nil, err
	}

	return ids, nil
}