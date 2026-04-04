package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// checks for emptiness of name/id-list-queryParam and converts its input into a slice of valid ids. accepts "none" as input.
func parseNameIdListQuery[P h.HasID](r *http.Request, queryParam QueryParam, pResType string, pLookup map[string]P) ([]int32, error) {
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

// converts a list of unique query ids or single-resource names into a slice of valid ids.
func queryNamesIDsToSlice[P h.HasID](query string, queryParam QueryParam, pResType string, pLookup map[string]P) ([]int32, error) {
	queryStrs := querySplit(query, ",")
	ids := []int32{}

	for _, str := range queryStrs {
		id, err := checkQueryNameID(str, pResType, queryParam, pLookup)
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
