package api

import (
	"context"
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) queryIntList(server *Server, r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryIntList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		ints, err := queryParam.parseIntList(server, r)
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
