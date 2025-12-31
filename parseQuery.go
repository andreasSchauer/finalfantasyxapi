package main

import (
	"fmt"
	"net/http"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func parseBooleanQuery(r *http.Request, queryParam QueryType) (bool, bool, error) {
	query := r.URL.Query().Get(queryParam.Name)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return false, isEmpty, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value. usage: %s", queryParam.Usage), err)
	}

	return b, isEmpty, nil
}

func parseTypeQuery(r *http.Request, queryParam QueryType, lookup map[string]TypedAPIResource) (TypedAPIResource, bool, error) {
	query := r.URL.Query().Get(queryParam.Name)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return TypedAPIResource{}, isEmpty, nil
	}

	enum, err := GetEnumType(query, lookup)
	if err != nil {
		return TypedAPIResource{}, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: '%s', use /api/%s to see valid values", query, queryParam.Name), err)
	}

	return enum, isEmpty, nil
}

func parseUniqueNameQuery[T h.HasID](r *http.Request, queryParam QueryType, lookup map[string]T) (parseResponse, bool, error) {
	query := r.URL.Query().Get(queryParam.Name)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return parseResponse{}, isEmpty, nil
	}

	resource, err := parseSingleSegmentResource(queryParam.Name, query, queryParam.Name, lookup)
	if err != nil {
		return parseResponse{}, false, err
	}

	return resource, isEmpty, nil
}

func parseIDBasedQuery(r *http.Request, queryParam QueryType, maxID int) (int32, bool, error) {
	query := r.URL.Query().Get(queryParam.Name)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return 0, isEmpty, nil
	}

	id, err := strconv.Atoi(query)
	if err != nil {
		return 0, false, newHTTPError(http.StatusBadRequest, "invalid id", err)
	}

	if id > maxID || id <= 0 {
		return 0, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided %s ID %d is out of range. Max ID: %d", queryParam.Name, id, maxID), err)
	}

	return int32(id), isEmpty, nil
}

func queryStrToInt(s string, defaultVal int) (int, error) {
	if s == "" {
		return defaultVal, nil
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return val, nil
}
