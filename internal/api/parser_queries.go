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

func parseIDOnlyQueryNul(r *http.Request, queryParam QueryType, maxID int) (*int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	err = checkNoneQuery(query)
	if err != nil {
		return nil, nil
	}

	id, err := parseQueryIdVal(query, queryParam, maxID)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// used for queryParams that thake existing ids or unique names and return a valid id
func parseNameOrIdQuery[P h.HasID](r *http.Request, queryParam QueryType, pResType string, pLookup map[string]P) (int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	return parseQueryNamedVal(query, pResType, queryParam, pLookup)
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
func parseEnumQuery[E, N any](r *http.Request, endpoint string, queryParam QueryType, et EnumType[E, N]) (EnumAPIResource, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return EnumAPIResource{}, err
	}

	return checkEnum(query, endpoint, queryParam, et)
}

// checks for default values, special values, validity, and range validity of an integer-based non-id query. if the query doesn't use defaults, special vals, or ranges, they are simply ignored.
func parseIntQuery(r *http.Request, queryParam QueryType) (int, error) {
	query := r.URL.Query().Get(queryParam.Name)
	// checkEmptyQuery should happen here
	val, err := checkQueryInt(queryParam, query)
	if errors.Is(err, errEmptyQuery) {
		return 0, errEmptyQuery
	}
	if err != nil {
		return 0, err
	}

	return val, nil
}

// converts a list of unique ids into a slice and checks every id's validity. duplicates produce an error.
func parseIdListQuery(r *http.Request, queryParam QueryType, maxID int) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryIDsToSlice(query, queryParam, maxID)
}

func parseNameIdListQuery[P h.HasID](r *http.Request, queryParam QueryType, pResType string, pLookup map[string]P) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	err = checkNoneQuery(query)
	if err != nil {
		return nil, nil
	}

	return queryNamesIDsToSlice(query, queryParam, pResType, pLookup)
}

func parseIntListQuery(r *http.Request, queryParam QueryType) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryIntsToSlice(query, queryParam)
}

func parseEnumListQuery[E, N any](r *http.Request, endpoint string, queryParam QueryType, et EnumType[E, N]) ([]E, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryEnumsToSlice(query, endpoint, queryParam, et)
}

func parseResTypeQuery(r *http.Request, queryParam QueryType) (string, string, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return "", "", err
	}

	resType, unitStr, found := strings.Cut(query, ":")
	if !found {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, query, queryParam.Usage), nil)
	}

	if !slices.Contains(queryParam.AllowedResTypes, resType) {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resource type '%s' for parameter '%s'. supported resource types: %s.", resType, queryParam.Name, h.FormatStringSlice(queryParam.AllowedResTypes)), nil)
	}

	return resType, unitStr, nil
}
