package main

import (
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"strings"
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

func queryIDsToSlice(query string, queryParam QueryType, maxID int) ([]int32, error) {
	idStrs := strings.Split(query, ",")
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

func checkDuplicateIDs(queryParam QueryType, ids []int32) error {
	idMap := make(map[int32]bool)

	for _, id := range ids {
		if idMap[id] {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of id %d in %s. each id can only be used once.", id, queryParam.Name), nil)
		}
		idMap[id] = true
	}

	return nil
}
