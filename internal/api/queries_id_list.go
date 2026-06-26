package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// query uses a list of ids as database input to filter for resources
func idListQuery[T, F seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, fLookup map[string]F, dbQuery DbQueryIntList) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		queryIDs, err := parseIdListQuery(cfg, r, queryParam, fLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, queryIDs)
		if err != nil {
			return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// like idListQuery, but with more specialized logic in between (wrapperFn)
func idListQueryWrapper[T, F seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, fLookup map[string]F, wrapperFn func(*Config, *http.Request, context.Context, []int32) ([]int32, error)) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]

		queryIDs, err := parseIdListQuery(cfg, r, queryParam, fLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(cfg, r, ctx, queryIDs)
	}
}
