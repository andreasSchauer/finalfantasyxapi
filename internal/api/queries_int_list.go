package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func intListQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, dbQuery DbQueryIntList) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		ints, err := parseIntListQuery(cfg, r, queryParam)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, ints)
		if err != nil {
			return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}