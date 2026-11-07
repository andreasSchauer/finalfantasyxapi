package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses an integer value as input.

func (i HandlerInput[T, R, A]) QueryInt(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		integer, err := queryParam.ParseInt(r)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, int32(integer))
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}

// query uses an integer value as input.

func (i HandlerInput[T, R, A]) QueryIntWrapper(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, wrapperFn QueryWrapInt) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		integer, err := queryParam.ParseInt(r)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := wrapperFn(r, ctx, int32(integer))
		if errors.Is(err, errs.ErrQueryRedirect) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return dbIDs, nil
	}
}
