package api

import (
	"context"
	"errors"
	"net/http"
)

// query uses an integer value as input.
func (i handlerInput[T, R, A]) queryInt(r *http.Request, inputIDs []int32, queryName QueryParamName, dbQuery DBQueryIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		integer, err := parseIntQuery(r, queryParam)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, int32(integer))
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// query uses an integer value as input.
func (i handlerInput[T, R, A]) queryIntWrapper(cfg *Config, r *http.Request, inputIDs []int32, queryName QueryParamName, wrapperFn QueryWrapInt) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		integer, err := parseIntQuery(r, queryParam)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := wrapperFn(cfg, r, ctx, int32(integer))
		if errors.Is(err, errQueryRedirect) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return dbIDs, nil
	}
}
