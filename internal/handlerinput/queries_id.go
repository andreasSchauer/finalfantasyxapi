package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses an id of another resource type to filter resources

func (i Main[T, R, A]) QueryID[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, fLookup map[string]F, dbQuery DBQueryIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.ParseID(r, fLookup)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, id)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}

// like idOnlyQuery, but with more specialized logic in between (wrapperFn)

func (i Main[T, R, A]) QueryIDWrapper[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, fLookup map[string]F, wrapperFn QueryWrapInt) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.ParseID(r, fLookup)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(r, ctx, id)
	}
}

func (i Main[T, R, A]) QueryIDNul[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, fLookup map[string]F, dbQuery DBQueryNullIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		idPtr, err := queryParam.ParseIDPtr(r, fLookup)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, h.GetNullInt32(idPtr))
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}
