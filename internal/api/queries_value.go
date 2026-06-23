package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func valueQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, dbQuery DbQueryStringMany) ([]int32, error) {
	queryParam := i.queryLookup[queryName]
	if replParamsPresent(r, queryParam, i.queryLookup) {
		return inputIDs, nil
	}

	value, err := parseValueQuery(r, queryParam)
	if queryIsEmpty(err) {
		return inputIDs, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), value)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
	}

	return dbIDs, nil
}
