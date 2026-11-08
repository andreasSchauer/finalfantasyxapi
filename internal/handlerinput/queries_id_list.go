package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses a list of ids as database input to filter for resources

func (i Main[T, R, A]) QueryIDList[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, fLookup map[string]F, dbQuery DBQueryIntList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		queryIDs, err := queryParam.ParseIDList(r, fLookup)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, queryIDs)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}

// like idListQuery, but with more specialized logic in between (wrapperFn)

func (i Main[T, R, A]) QueryIDListWrapper[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, fLookup map[string]F, wrapperFn QueryWrapIntSlice) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]

		queryIDs, err := queryParam.ParseIDList(r, fLookup)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(r, ctx, queryIDs)
	}
}
