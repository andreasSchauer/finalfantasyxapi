package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// query uses the name or id of another resource type to filter resources
func nameIdQuery[T, P seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, pResType ResTypeSingular, pLookup map[string]P, dbQuery DbQueryIntMany) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := parseNameIdQuery(r, queryParam, pResType, pLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, id)
		if err != nil {
			return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

func nameIdQueryWrapper[T, P seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, pResType ResTypeSingular, pLookup map[string]P, wrapperFn func(*Config, *http.Request, context.Context, int32) ([]int32, error)) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := parseNameIdQuery(r, queryParam, pResType, pLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(cfg, r, ctx, id)
	}
}
