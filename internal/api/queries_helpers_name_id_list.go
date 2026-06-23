package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// checks for emptiness of name/id-list-queryParam and converts its input into a slice of valid ids. accepts "none" as input.
func parseNameIdListQuery[P seeding.Lookupable](cfg *Config, r *http.Request, queryParam QueryParam, pResType ResTypeSingular, pLookup map[string]P) ([]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	err = checkNoneQuery(query)
	if err != nil {
		return nil, nil
	}

	return queryNamesIDsToSlice(cfg, query, queryParam, pResType, pLookup)
}

// converts a list of unique query ids or single-resource names into a slice of valid ids.
func queryNamesIDsToSlice[P seeding.Lookupable](cfg *Config, query string, queryParam QueryParam, pResType ResTypeSingular, pLookup map[string]P) ([]int32, error) {
	queryStrs, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}
	idMap := make(map[string]int32)

	for _, str := range queryStrs {
		_, ok := idMap[str]
		if ok {
			continue
		}

		id, err := checkQueryNameID(str, pResType, queryParam, pLookup)
		if err != nil {
			return nil, err
		}

		idMap[str] = id
	}

	ids := queryIntMapToSlice(idMap)

	return ids, nil
}
