package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) queryEnumList[E, N any](server *Server, r *http.Request, et enums.EnumType[E, N], inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryEnumList[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		enums, err := queryParam.parseEnumList(server, r, et)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, enums)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}
