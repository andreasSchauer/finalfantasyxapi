package queryparam

import (
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

// ParseIntList checks for emptiness of int-list-queryParam and converts its input into a slice of valid integers
func (q QueryParam) ParseIntList(r *http.Request) ([]int32, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return nil, err
	}

	return q.sliceInts(r, query)
}

// converts a list of unique query ints into a slice of valid integers. also deals with ranged inputs.
func (q QueryParam) sliceInts(r *http.Request, query string) ([]int32, error) {
	intSegments, err := QueryListSplit(query)
	if err != nil {
		return nil, err
	}
	ints := []int32{}

	for _, segment := range intSegments {
		intsNew, err := q.checkIntRange(segment)
		if err != nil {
			return nil, err
		}
		ints = slices.Concat(ints, intsNew)

		if len(ints) > config.Cfg.FetchLimit {
			return nil, errs.NewHTTPErrorFetchLimit(config.Cfg.FetchLimit)
		}
	}

	ints, err = cleanUpIntList(r, ints)
	if err != nil {
		return nil, err
	}

	return ints, nil
}

// parses a single item of a query id-list.
func (q QueryParam) checkIntRange(segment string) ([]int32, error) {
	intStrs := strings.Split(segment, "-")
	ints := []int32{}

	switch len(intStrs) {
	case 0:
		return nil, nil

	case 1:
		integer, err := q.CheckInt(intStrs[0])
		if err != nil {
			return nil, err
		}
		ints = append(ints, int32(integer))

	case 2:
		newInts, err := q.sliceIntRange(intStrs)
		if err != nil {
			return nil, err
		}
		ints = slices.Concat(ints, newInts)

	default:
		return nil, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'", q.Name, segment, q.Usage), nil)
	}

	return ints, nil
}

// converts the two values of a ranged int query input into a slice of integers.
func (q QueryParam) sliceIntRange(intStrs []string) ([]int32, error) {
	minInt, err := q.CheckInt(intStrs[0])
	if err != nil {
		return nil, err
	}

	maxInt, err := q.CheckInt(intStrs[1])
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

func cleanUpIntList(r *http.Request, ints []int32) ([]int32, error) {
	ints = removeIntListDuplicates(ints)

	limit, err := getRawQueryLimit(r)
	if errs.QueryIsEmpty(err) {
		return ints, nil
	}
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
