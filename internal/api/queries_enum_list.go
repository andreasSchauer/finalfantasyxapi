package api

import (
	"context"
	"net/http"
)

func (i handlerInput[T, R, A]) queryEnumList[E, N any](cfg *Config, r *http.Request, et EnumType[E, N], inputIDs []int32, queryName QueryParamName, dbQuery DBQueryEnumList[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		enums, err := parseEnumListQuery(cfg, r, i.Endpoint, queryParam, et)
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
