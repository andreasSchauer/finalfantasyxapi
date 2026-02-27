package api

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// used, if a queryParam only takes existing ids and returns a valid id
func parseIDOnlyQuery(r *http.Request, queryParam QueryType, maxID int) (int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	id, err := parseQueryIdVal(query, queryParam, maxID)
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// used for queryParams that thake existing ids or unique names and return a valid id
func parseNameOrIdQuery[T h.HasID](r *http.Request, queryParam QueryType, resourceType string, lookup map[string]T) (int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	return parseQueryNamedVal(query, resourceType, queryParam, lookup)
}

// used for boolean queryParams
func parseBooleanQuery(r *http.Request, queryParam QueryType) (bool, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid boolean value '%s' used for parameter '%s'. usage: '%s'.", query, queryParam.Name, queryParam.Usage), err)
	}

	return b, nil
}

// used, if a queryParam is looking up an enum entry
func parseTypeQuery[T, N any](r *http.Request, endpoint string, queryParam QueryType, et EnumType[T, N]) (TypedAPIResource, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return TypedAPIResource{}, err
	}

	enum, err := GetTypedAPIResource(query, et)
	if err != nil {
		return TypedAPIResource{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid enum value '%s' used for parameter '%s'. use /api/%s/parameters to see allowed values.", query, queryParam.Name, endpoint), err)
	}

	return enum, nil
}

// checks for default values, special values, validity, and range validity of an integer-based non-id query. if the query doesn't use defaults, special vals, or ranges, they are simply ignored.
func parseIntQuery(r *http.Request, queryParam QueryType) (int, error) {
	query := r.URL.Query().Get(queryParam.Name)

	defaultVal, err := checkQueryIntDefaultVal(queryParam, query)
	if errors.Is(err, errEmptyQuery) {
		return 0, errEmptyQuery
	}
	if !errors.Is(err, errNoDefaultVal) {
		return defaultVal, nil
	}

	specialVal, err := checkQueryIntSpecialVals(queryParam, query)
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
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	ids, err := queryIDsToSlice(query, queryParam, maxID)
	if err != nil {
		return nil, err
	}

	return ids, nil
}




func parseResTypeQuery(r *http.Request, queryParam QueryType) (string, string, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return "", "", err
	}

	resType, idStr, found := strings.Cut(query, ":")
	if !found {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, query, queryParam.Usage), nil)
	}

	if !slices.Contains(queryParam.AllowedResTypes, resType) {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resource type '%s' for parameter '%s'. supported resource types: %s.", resType, queryParam.Name, h.FormatStringSlice(queryParam.AllowedResTypes)), nil)
	}

	return resType, idStr, nil
}

func checkEmptyQuery(r *http.Request, queryParam QueryType) (string, error) {
	query := r.URL.Query().Get(queryParam.Name)
	if query == "" {
		return "", errEmptyQuery
	}

	return strings.ToLower(query), nil
}
