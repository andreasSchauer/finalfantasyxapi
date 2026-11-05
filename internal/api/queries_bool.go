package api

import (
	"context"
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// db query searches for resources with matching boolean db column value
func (i handlerInput[T, R, A]) queryBool(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryBool) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		b, err := queryParam.parseBool(r)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, b)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// db query accumulates all resources that fulfill a certain condition. a false boolean flips these results.
func (i handlerInput[T, R, A]) queryBool2(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryNoInput) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		b, err := queryParam.parseBool(r)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		if !b {
			dbIDs = removeIDs(inputIDs, dbIDs)
		}

		return dbIDs, nil
	}
}

func (i handlerInput[T, R, A]) queryBoolWrapper(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, wrapperFn QueryWrapBool) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		b, err := queryParam.parseBool(r)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(r, ctx, b)
	}
}
