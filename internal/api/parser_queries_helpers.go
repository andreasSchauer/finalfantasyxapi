package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// parses an id or single-segment-resource name and returns a valid id
func parseQueryNamedVal[T h.HasID](query, resourceType string, queryParam QueryType, lookup map[string]T) (int32, error) {
	id, err := checkQueryIDVal(query, queryParam, len(lookup))
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, errNotAnID) {
		return 0, err
	}

	resource, err := checkUniqueName(query, lookup)
	if err == nil {
		return resource.ID, nil
	}

	return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown %s '%s' used for parameter '%s'.", resourceType, query, queryParam.Name), err)
}

// checks if query ID is valid and within range. if it's invalid, it will return an httpError
func parseQueryIdVal(idStr string, queryParam QueryType, maxID int) (int32, error) {
	id, err := checkQueryIDVal(idStr, queryParam, maxID)
	if errors.Is(err, errNotAnID) {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%s' used for parameter '%s'.", idStr, queryParam.Name), err)
	}
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// checks if query ID is valid and within range. if it's invalid, it will return errNotAnID
func checkQueryIDVal(idStr string, queryParam QueryType, maxID int) (int32, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errNotAnID
	}

	if id > maxID || id <= 0 {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided id '%d' used for parameter '%s' is out of range. max id: %d.", id, queryParam.Name, maxID), err)
	}

	return int32(id), nil
}

func checkQueryInt(queryParam QueryType, queryVal string) (int, error) {
	defaultVal, err := checkIntQueryDefaultVal(queryParam, queryVal)
	if errors.Is(err, errEmptyQuery) {
		return 0, errEmptyQuery
	}
	if !errors.Is(err, errNoDefaultVal) {
		return defaultVal, nil
	}

	specialVal, err := checkIntQuerySpecialVals(queryParam, queryVal)
	if !errors.Is(err, errNoSpecialInput) {
		return specialVal, nil
	}

	val, err := checkIntQueryAllowedRange(queryParam, queryVal)
	if err != nil && !errors.Is(err, errNoIntRange) {
		return 0, err
	}

	return val, nil
}

func checkIntQueryDefaultVal(queryParam QueryType, s string) (int, error) {
	if queryParam.DefaultVal == nil {
		if s == "" {
			return 0, errEmptyQuery
		}
		return 0, errNoDefaultVal
	}

	if s == "" {
		return *queryParam.DefaultVal, nil
	}

	return 0, errNoDefaultVal
}

func checkIntQuerySpecialVals(queryParam QueryType, s string) (int, error) {
	if queryParam.SpecialInputs == nil {
		return 0, errNoSpecialInput
	}

	for _, input := range queryParam.SpecialInputs {
		if s == input.Key {
			return input.Val, nil
		}
	}

	return 0, errNoSpecialInput
}

func checkIntQueryAllowedRange(queryParam QueryType, s string) (int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. usage: '%s'.", s, queryParam.Name, queryParam.Usage), err)
	}

	intRange := queryParam.AllowedIntRange
	if intRange == nil {
		return val, errNoIntRange
	}

	min := intRange[0]
	max := intRange[1]

	if val > max || val < min {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%d' used for parameter '%s'. value must be an integer ranging from %d to %d.", val, queryParam.Name, min, max), nil)
	}

	return val, nil
}

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

	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, segment, queryParam.Usage), nil)
	}

	return ints, nil
}
