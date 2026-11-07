package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses the name or id of another resource type to filter resources

func (i HandlerInput[T, R, A]) QueryNameID[P seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, pResType ts.ResTypeSingle, pLookup map[string]P, dbQuery DBQueryIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.ParseNameID(r, pResType, pLookup)
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

func (i HandlerInput[T, R, A]) QueryNameIDWrapper[P seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, pResType ts.ResTypeSingle, pLookup map[string]P, wrapperFn QueryWrapInt) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		id, err := queryParam.ParseNameID(r, pResType, pLookup)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(r, ctx, id)
	}
}
