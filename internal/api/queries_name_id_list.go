package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// query uses a list of names or ids as database input to filter for resources. alternatively, "none" can be used as input instead.
func nameIdListQueryNul[T, P seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName, pResType string, pLookup map[string]P, dbQuery DbQueryIntList) ([]int32, error) {
	queryParam := i.queryLookup[queryName]
	if replParamsPresent(r, queryParam, i.queryLookup) {
		return inputIDs, nil
	}

	queryIDs, err := parseNameIdListQuery(cfg, r, queryParam, pResType, pLookup)
	if queryIsEmpty(err) {
		return inputIDs, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), queryIDs)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	return dbIDs, nil
}
