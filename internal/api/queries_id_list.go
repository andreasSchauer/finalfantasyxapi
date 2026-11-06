package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses a list of ids as database input to filter for resources
func (i handlerInput[T, R, A]) queryIDList[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, fLookup map[string]F, dbQuery DBQueryIntList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		queryIDs, err := queryParam.ParseIDList(r, fLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, queryIDs)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// like idListQuery, but with more specialized logic in between (wrapperFn)
func (i handlerInput[T, R, A]) queryIDListWrapper[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, fLookup map[string]F, wrapperFn QueryWrapIntSlice) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]

		queryIDs, err := queryParam.ParseIDList(r, fLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(r, ctx, queryIDs)
	}
}
