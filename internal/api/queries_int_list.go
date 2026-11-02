package api

import (
	"context"
	"net/http"
)

func (i handlerInput[T, R, A]) queryIntList(cfg *Config, r *http.Request, inputIDs []int32, queryName QueryParamName, dbQuery DBQueryIntList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		ints, err := queryParam.parseIntList(cfg, r)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, ints)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}
