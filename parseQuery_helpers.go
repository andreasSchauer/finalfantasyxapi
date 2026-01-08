package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// parses an id or single-segment-resource name and returns a valid id
func parseQueryNamedVal[T h.HasID](query, resourceType string, queryParam QueryType, lookup map[string]T) (int32, error) {
	id, err := parseQueryIDVal(query, queryParam, len(lookup))
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, errNotAnID) {
		return 0, err
	}

	resource, err := parseUniqueName(query, lookup)
	if err == nil {
		return resource.ID, nil
	}

	return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown %s '%s' in '%s'.", resourceType, query, queryParam.Name), err)
}

// checks if query ID is valid and within range. if it's invalid, it will return errNotAnID
func parseQueryIDVal(idStr string, queryParam QueryType, maxID int) (int32, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errNotAnID
	}

	if id > maxID || id <= 0 {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided ID %d in '%s' is out of range. Max ID: %d", id, queryParam.Name, maxID), err)
	}

	return int32(id), nil
}

// checks if query ID is valid and within range. if it's invalid, it will return an httpError
func parseQueryIdValStrict(idStr string, queryParam QueryType, maxID int) (int32, error) {
	id, err := parseQueryIDVal(idStr, queryParam, maxID)
	if errors.Is(err, errNotAnID) {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%s' used for parameter '%s'", idStr, queryParam.Name), err)
	}
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// converts an int string from a query to an int and uses the default value, if the string is empty
func AtoiOrDefault(s string, defaultVal int) (int, error) {
	if s == "" {
		return defaultVal, nil
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return val, nil
}
