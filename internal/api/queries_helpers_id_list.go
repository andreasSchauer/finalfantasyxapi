package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

// checks for emptiness of id-list-queryParam and converts its input into a slice of valid ids.
func parseIdListQuery(cfg *Config, r *http.Request, queryParam QueryParam, maxID int) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryIDsToSlice(cfg, r, query, queryParam, maxID)
}

// converts a list of unique query ids into a slice of valid ids.
func queryIDsToSlice(cfg *Config, r *http.Request, query string, queryParam QueryParam, maxID int) ([]int32, error) {
	idSegments, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}
	ids := []int32{}

	for _, segment := range idSegments {
		idsNew, err := checkQueryIdRange(queryParam, segment, maxID)
		if err != nil {
			return nil, err
		}
		ids = slices.Concat(ids, idsNew)

		if len(ids) > cfg.fetchLimit {
			return nil, newHTTPErrorFetchLimit(cfg.fetchLimit)
		}
	}

	ids, err = cleanUpIntList(cfg, r, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func checkQueryIdRange(queryParam QueryParam, segment string, maxID int) ([]int32, error) {
	idStrs := strings.Split(segment, "-")
	ids := []int32{}

	switch len(idStrs) {
	case 0:
		return nil, nil

	case 1:
		id, err := parseQueryID(idStrs[0], queryParam, maxID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)

	case 2:
		newIDs, err := idRangeToSlice(queryParam, idStrs, maxID)
		if err != nil {
			return nil, err
		}
		ids = slices.Concat(ids, newIDs)

	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, segment, queryParam.Usage), nil)
	}

	return ids, nil
}

func idRangeToSlice(queryParam QueryParam, idStrs []string, maxID int) ([]int32, error) {
	minId, err := parseQueryID(idStrs[0], queryParam, maxID)
	if err != nil {
		return nil, err
	}

	maxId, err := parseQueryID(idStrs[1], queryParam, maxID)
	if err != nil {
		return nil, err
	}

	ids := sliceFromIntRange(minId, maxId)

	return ids, nil
}
