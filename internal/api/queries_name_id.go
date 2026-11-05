package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses the name or id of another resource type to filter resources
func (i handlerInput[T, R, A]) queryNameID[P seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, pResType ts.ResTypeSingle, pLookup map[string]P, dbQuery DBQueryIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.parseNameID(r, pResType, pLookup)
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

func (i handlerInput[T, R, A]) queryNameIDWrapper[P seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, pResType ts.ResTypeSingle, pLookup map[string]P, wrapperFn QueryWrapInt) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.parseNameID(r, pResType, pLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(r, ctx, id)
	}
}
