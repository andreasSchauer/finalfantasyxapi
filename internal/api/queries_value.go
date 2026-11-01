package api

import (
	"context"
	"net/http"
)

func (i handlerInput[T, R, A]) queryValue(r *http.Request, inputIDs []int32, queryName QueryParamName, dbQuery DBQueryStringMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		value, err := parseValueQuery(r, queryParam)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, value)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}
