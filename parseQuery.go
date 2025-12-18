package main

import (
	"fmt"
	"net/http"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func parseBooleanQuery(r *http.Request, queryParam string) (bool, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return false, isEmpty, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value. usage: %s={boolean}", queryParam), err)
	}

	return b, isEmpty, nil
}

func parseTypeQuery(r *http.Request, queryParam string, lookup map[string]TypedAPIResource) (TypedAPIResource, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return TypedAPIResource{}, isEmpty, nil
	}

	enum, err := GetEnumType(query, lookup)
	if err != nil {
		return TypedAPIResource{}, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/%s to see valid values", query, queryParam), err)
	}

	return enum, isEmpty, nil
}

func parseUniqueNameQuery[T h.HasID](r *http.Request, queryParam string, lookup map[string]T) (parseResponse, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return parseResponse{}, isEmpty, nil
	}

	resource, err := parseSingleSegmentResource(queryParam, query, lookup)
	if err != nil {
		return parseResponse{}, false, err
	}

	return resource, isEmpty, nil
}

func parseIDBasedQuery(r *http.Request, queryParam string, maxID int) (int32, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return 0, isEmpty, nil
	}

	id, err := strconv.Atoi(query)
	if err != nil {
		return 0, false, newHTTPError(http.StatusBadRequest, "invalid id", err)
	}

	if id > maxID {
		return 0, false, newHTTPError(http.StatusNotFound, fmt.Sprintf("provided %s ID is out of range. Max ID: %d", query, maxID), err)
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
