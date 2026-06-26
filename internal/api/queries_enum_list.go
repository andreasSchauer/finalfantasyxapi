package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func enumListQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList, E, N any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputIDs []int32, queryName QueryParamName, dbQuery DbQueryEnumList[E]) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		enums, err := parseEnumListQuery(cfg, r, i.endpoint, queryParam, et)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, enums)
		if err != nil {
			return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}
