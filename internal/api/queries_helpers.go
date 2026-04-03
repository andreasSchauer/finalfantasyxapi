package api

import (
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func queryMapToSlice(lookup map[string]QueryType) []QueryType {
	queryParams := []QueryType{}

	for key := range lookup {
		queryParams = append(queryParams, lookup[key])
	}

	slices.SortStableFunc(queryParams, func(a, b QueryType) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return queryParams
}

func queryMapToString(lookup map[string]QueryType) string {
	params := queryMapToSlice(lookup)
	names := []string{}

	for _, param := range params {
		names = append(names, param.Name)
	}

	return h.FormatStringSlice(names)
}

func queryIDsToSliceNoDupes(query string, queryParam QueryType, maxID int) ([]int32, error) {
	idStrs := querySplit(query, ",")
	ids := []int32{}

	for _, idStr := range idStrs {
		id, err := parseQueryIdVal(idStr, queryParam, maxID)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	err := checkDuplicateIDs(queryParam, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func queryNamesIDsToSlice[P h.HasID](query string, queryParam QueryType, pResType string, pLookup map[string]P) ([]int32, error) {
	queryStrs := querySplit(query, ",")
	ids := []int32{}

	for _, str := range queryStrs {
		id, err := parseQueryNamedVal(str, pResType, queryParam, pLookup)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	err := checkDuplicateIDs(queryParam, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func queryIntsToSlice(query string, queryParam QueryType) ([]int32, error) {
	intSegments := querySplit(query, ",")
	ints := []int32{}

	for _, segment := range intSegments {
		intsNew, err := checkQueryIntRange(queryParam, segment)
		if err != nil {
			return nil, err
		}
		ints = slices.Concat(ints, intsNew)
	}

	err := checkDuplicateInts(queryParam, ints)
	if err != nil {
		return nil, err
	}

	return ints, nil
}

func checkEnum[E, N any](val, endpoint string, queryParam QueryType, et EnumType[E, N]) (EnumAPIResource, error) {
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

func queryEnumsToSlice[E, N any](query, endpoint string, queryParam QueryType, et EnumType[E, N]) ([]E, error) {
	enumStrs := querySplit(query, ",")
	enums := []E{}

	for _, enumStr := range enumStrs {
		enum, err := checkEnum(enumStr, endpoint, queryParam, et)
		if err != nil {
			return nil, err
		}
		typedStr := et.convFunc(enum.Name)
		enums = append(enums, typedStr)
	}

	err := checkDuplicateEnums(queryParam, enums)
	if err != nil {
		return nil, err
	}

	return enums, nil
}

func checkEmptyQuery(r *http.Request, queryParam QueryType) (string, error) {
	query := r.URL.Query().Get(queryParam.Name)
	if query == "" {
		return "", errEmptyQuery
	}

	return strings.ToLower(query), nil
}

func checkNoneQuery(query string) error {
	if query == "none" {
		return errQueryNone
	}

	return nil
}

func checkDuplicateIDs(queryParam QueryType, ids []int32) error {
	idMap := make(map[int32]bool)

	for _, id := range ids {
		if idMap[id] {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of id '%d' for parameter '%s'. each id can only be used once.", id, queryParam.Name), nil)
		}
		idMap[id] = true
	}

	return nil
}

func checkDuplicateInts(queryParam QueryType, ints []int32) error {
	intMap := make(map[int32]bool)

	for _, int := range ints {
		if intMap[int] {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of value '%d' for parameter '%s'. each value can only be used once.", int, queryParam.Name), nil)
		}
		intMap[int] = true
	}

	return nil
}

func checkDuplicateEnums[E any](queryParam QueryType, enums []E) error {
	enumMap := make(map[any]bool)

	for _, enum := range enums {
		if enumMap[enum] {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of enum '%v' for parameter '%s'. each enum can only be used once.", enum, queryParam.Name), nil)
		}
		enumMap[enum] = true
	}

	return nil
}

func idStrsToUniqueIDs(idStrs []string, resourceType string, maxID int) ([]int32, error) {
	idMap := make(map[int32]bool)
	ids := []int32{}

	for _, idStr := range idStrs {
		resp, err := parseID(idStr, resourceType, maxID)
		if err != nil {
			return nil, err
		}
		id := resp.ID

		if idMap[id] {
			continue
		}

		idMap[id] = true
		ids = append(ids, id)
	}

	return ids, nil
}

func querySplit(query, sep string) []string {
	queryTrimmed := strings.TrimSuffix(query, sep)
	return strings.Split(queryTrimmed, sep)
}
