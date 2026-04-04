package api

import (
	"fmt"
	"net/http"
)

// checks for emptiness of id-list-queryParam and converts its input into a slice of valid ids.
func parseIdListQuery(r *http.Request, queryParam QueryParam, maxID int) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryIDsToSlice(query, queryParam, maxID)
}

// converts a list of unique query ids into a slice of valid ids.
func queryIDsToSlice(query string, queryParam QueryParam, maxID int) ([]int32, error) {
	idStrs := querySplit(query, ",")
	ids := []int32{}
	const fetchLimit = 50

	if len(idStrs) > fetchLimit {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("fetch limit exceeded. the maximum amount of resources that can be fetched is %d.", fetchLimit), nil)
	}

	for _, idStr := range idStrs {
		id, err := parseQueryID(idStr, queryParam, maxID)
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

// checks, if there are duplicate ids in a slice.
func checkDuplicateIDs(queryParam QueryParam, ids []int32) error {
	idMap := make(map[int32]bool)

	for _, id := range ids {
		if idMap[id] {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of id '%d' for parameter '%s'. each id can only be used once.", id, queryParam.Name), nil)
		}
		idMap[id] = true
	}

	return nil
}
