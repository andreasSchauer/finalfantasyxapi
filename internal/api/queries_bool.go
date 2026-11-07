package api

import (
	"context"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// db query searches for resources with matching boolean db column value

func (i HandlerInput[T, R, A]) QueryBool(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryBool) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		b, err := queryParam.ParseBool(r)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, b)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}

// db query accumulates all resources that fulfill a certain condition. a false boolean flips these results.

func (i HandlerInput[T, R, A]) QueryBool2(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryNoInput) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		b, err := queryParam.ParseBool(r)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		if !b {
			dbIDs = ah.RemoveIDs(inputIDs, dbIDs)
		}

		return dbIDs, nil
	}
}

func (i HandlerInput[T, R, A]) QueryBoolWrapper(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, wrapperFn QueryWrapBool) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		b, err := queryParam.ParseBool(r)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(r, ctx, b)
	}
}
