package api

import (
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"strings"
)

// checks for emptiness of int-list-queryParam and converts its input into a slice of valid integers
func parseIntListQuery(cfg *Config, r *http.Request, queryParam QueryParam) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryIntsToSlice(cfg, r, query, queryParam)
}

// converts a list of unique query ints into a slice of valid integers. also deals with ranged inputs.
func queryIntsToSlice(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
	intSegments, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}
	ints := []int32{}

	for _, segment := range intSegments {
		intsNew, err := checkQueryIntRange(queryParam, segment)
		if err != nil {
			return nil, err
		}
		ints = slices.Concat(ints, intsNew)

		if len(ints) > cfg.fetchLimit {
			return nil, newHTTPErrorFetchLimit(cfg.fetchLimit)
		}
	}

	ints, err = cleanUpIntList(cfg, r, ints)
	if err != nil {
		return nil, err
	}

	return ints, nil
}

// parses a single item of a query id-list.
func checkQueryIntRange(queryParam QueryParam, segment string) ([]int32, error) {
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
func intRangeToSlice(queryParam QueryParam, intStrs []string) ([]int32, error) {
	minInt, err := checkQueryInt(queryParam, intStrs[0])
	if err != nil {
		return nil, err
	}

	maxInt, err := checkQueryInt(queryParam, intStrs[1])
	if err != nil {
		return nil, err
	}

	ints := sliceFromIntRange(int32(minInt), int32(maxInt))

	return ints, nil
}

func sliceFromIntRange(min, max int32) []int32 {
	ints := []int32{}

	if min > max {
		temp := min
		min = max
		max = temp
	}

	for i := min; i <= max; i++ {
		ints = append(ints, int32(i))
	}

	return ints
}

func cleanUpIntList(cfg *Config, r *http.Request, ints []int32) ([]int32, error) {
	ints = removeIntListDuplicates(ints)

	limit, err := getQueryLimit(cfg, r)
	if err != nil {
		return nil, err
	}

	if limit <= len(ints) {
		return ints[0:limit], nil
	}

	return ints, nil
}


func removeIntListDuplicates(ints []int32) []int32 {
	intMap := make(map[int32]bool)
	newInts := []int32{}

	for _, integer := range ints {
		if intMap[integer] {
			continue
		}
		intMap[integer] = true
		newInts = append(newInts, integer)
	}

	slices.SortStableFunc(newInts, func(a, b int32) int {
		return cmp.Compare(a, b)
	})

	return newInts
}