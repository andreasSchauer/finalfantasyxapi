package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

// checks for emptiness of int-list-queryParam and converts its input into a slice of valid integers
func parseIntListQuery(r *http.Request, queryParam QueryType) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryIntsToSlice(query, queryParam)
}

// converts a list of unique query ints into a slice of valid integers. also deals with ranged inputs.
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

// parses a single item of a query id-list.
func checkQueryIntRange(queryParam QueryType, segment string) ([]int32, error) {
	intStrs := strings.Split(segment, "-")
	ints := []int32{}

	switch len(intStrs) {
	case 0:
		return nil, nil

	case 1:
		integer, err := checkQueryInt(queryParam, intStrs[0])
		if err != nil {
			return nil, err
		}
		ints = append(ints, int32(integer))

	case 2:
		newInts, err := intRangeToSlice(queryParam, intStrs)
		if err != nil {
			return nil, err
		}
		ints = slices.Concat(ints, newInts)

	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, segment, queryParam.Usage), nil)
	}

	return ints, nil
}

// converts the two values of a ranged int query input into a slice of integers.
func intRangeToSlice(queryParam QueryType, intStrs []string) ([]int32, error) {
	ints := []int32{}
	
	minInt, err := checkQueryInt(queryParam, intStrs[0])
	if err != nil {
		return nil, err
	}

	maxInt, err := checkQueryInt(queryParam, intStrs[1])
	if err != nil {
		return nil, err
	}

	if minInt > maxInt {
		temp := minInt
		minInt = maxInt
		maxInt = temp
	}

	for i := minInt; i <= maxInt; i++ {
		ints = append(ints, int32(i))
	}

	return ints, nil
}


// checks, if there are duplicate ints in a slice.
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
