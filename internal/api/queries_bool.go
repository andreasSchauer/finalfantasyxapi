package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// db query searches for resources with matching boolean db column value
func boolQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, dbQuery DbQueryBool) ([]int32, error) {
	queryParam := i.queryLookup[queryName]
	if replParamsPresent(r, queryParam, i.queryLookup) {
		return inputIDs, nil
	}

	b, err := parseBooleanQuery(r, queryParam)
	if queryIsEmpty(err) {
		return inputIDs, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), b)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
	}

	return dbIDs, nil
}

// db query accumulates all resources that fulfill a certain condition. a false boolean flips these results.
func boolQuery2[T seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, dbQuery DbQueryNoInput) ([]int32, error) {
	queryParam := i.queryLookup[queryName]
	if replParamsPresent(r, queryParam, i.queryLookup) {
		return inputIDs, nil
	}

	b, err := parseBooleanQuery(r, queryParam)
	if queryIsEmpty(err) {
		return inputIDs, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context())
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
	}

	if !b {
		dbIDs = removeIDs(inputIDs, dbIDs)
	}

	return dbIDs, nil
}

func boolQueryWrapper[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, wrapperFn func(*Config, *http.Request, bool) ([]int32, error)) ([]int32, error) {
	queryParam := i.queryLookup[queryName]
	if replParamsPresent(r, queryParam, i.queryLookup) {
		return inputIDs, nil
	}

	b, err := parseBooleanQuery(r, queryParam)
	if queryIsEmpty(err) {
		return inputIDs, nil
	}
	if err != nil {
		return nil, err
	}

	return wrapperFn(cfg, r, b)
}
