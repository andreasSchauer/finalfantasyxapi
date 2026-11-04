package api

import (
	"context"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses an id of another resource type to filter resources
func (i handlerInput[T, R, A]) queryID[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName typestrings.QueryParamName, fLookup map[string]F, dbQuery DBQueryIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.parseID(r, fLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, id)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// like idOnlyQuery, but with more specialized logic in between (wrapperFn)
func (i handlerInput[T, R, A]) queryIDWrapper[F seedtypes.Lookupable](cfg *Config, r *http.Request, inputIDs []int32, queryName typestrings.QueryParamName, fLookup map[string]F, wrapperFn QueryWrapInt) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.parseID(r, fLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(cfg, r, ctx, id)
	}
}

func (i handlerInput[T, R, A]) queryIDNul[F seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName typestrings.QueryParamName, fLookup map[string]F, dbQuery DBQueryNullIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		idPtr, err := queryParam.parseIDPtr(r, fLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, h.GetNullInt32(idPtr))
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}
