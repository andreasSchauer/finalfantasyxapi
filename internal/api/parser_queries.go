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
func parseEnumQuery[E, N any](r *http.Request, endpoint string, queryParam QueryType, et EnumType[E, N]) (EnumAPIResource, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return EnumAPIResource{}, err
	}

	return parseEnum(query, endpoint, queryParam, et)
}

func parseEnumSliceQuery[E, N any](r *http.Request, endpoint string, queryParam QueryType, et EnumType[E, N]) ([]E, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	enumStrs := strings.Split(query, ",")
	enums := []E{}

	for _, str := range enumStrs {
		_, err := parseEnum(str, endpoint, queryParam, et)
		if err != nil {
			return nil, err
		}

		enum := et.convFunc(str)
		enums = append(enums, enum)
	}

	return enums, nil
}

func parseEnum[E, N any](val, endpoint string, queryParam QueryType, et EnumType[E, N]) (EnumAPIResource, error) {
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

// checks for default values, special values, validity, and range validity of an integer-based non-id query. if the query doesn't use defaults, special vals, or ranges, they are simply ignored.
func parseIntQuery(r *http.Request, queryParam QueryType) (int, error) {
	query := r.URL.Query().Get(queryParam.Name)

	val, err := checkQueryInt(queryParam, query)
	if errors.Is(err, errEmptyQuery) {
		return 0, errEmptyQuery
	}
	if err != nil {
		return 0, err
	}

	return val, nil
}

// converts a list of ids into a slice and checks every id's validity. duplicates are simply filtered out.
func parseIdListQuery[T h.HasID, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], queryName string, fetchLimit int) ([]int32, error) {
	queryParam := i.queryLookup[queryName]
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "parameter 'ids' can't be empty.", err)
	}

	idStrs := querySplit(query, ",")
	if len(idStrs) > fetchLimit {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("fetch limit exceeded. the maximum amount of resources that can be fetched is %d.", fetchLimit), nil)
	}

	ids, err := idStrsToUniqueIDs(idStrs, i.resourceType, len(i.objLookupID))
	if err != nil {
		return nil, err
	}

	return ids, nil
}

// converts a list of unique ids into a slice and checks every id's validity. duplicates produce an error.
func parseIdListQueryNoDupes(r *http.Request, queryParam QueryType, maxID int) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	ids, err := queryIDsToSliceNoDupes(query, queryParam, maxID)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func parseIntListQuery(r *http.Request, queryParam QueryType) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	ints, err := queryIntsToSlice(query, queryParam)
	if err != nil {
		return nil, err
	}

	return ints, nil
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
