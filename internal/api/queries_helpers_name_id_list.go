package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// ParseNameIDList checks for emptiness of name/id-list-queryParam and converts its input into a slice of valid ids. accepts "none" as input.
func (q QueryParam) ParseNameIDList[P seedtypes.Lookupable](r *http.Request, pResType ts.ResTypeSingle, pLookup map[string]P) ([]int32, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return nil, err
	}

	err = checkNoneQuery(query)
	if err != nil {
		return nil, nil
	}

	return q.sliceNamesIDs(query, pResType, pLookup)
}

// converts a list of unique query ids or single-resource names into a slice of valid ids.
func (q QueryParam) sliceNamesIDs[P seedtypes.Lookupable](query string, pResType ts.ResTypeSingle, pLookup map[string]P) ([]int32, error) {
	queryStrs, err := QueryListSplit(query)
	if err != nil {
		return nil, err
	}
	idMap := make(map[string]int32)

	for _, str := range queryStrs {
		_, ok := idMap[str]
		if ok {
			continue
		}

		id, err := q.CheckNameID(str, pResType, pLookup)
		if err != nil {
			return nil, err
		}

		idMap[str] = id
	}

	ids := queryIntMapToSlice(idMap)

	return ids, nil
}
