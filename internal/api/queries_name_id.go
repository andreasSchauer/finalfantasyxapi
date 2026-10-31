package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// query uses the name or id of another resource type to filter resources
func nameIDQuery[T, P seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, pResType ResTypeSingle, pLookup map[string]P, dbQuery DBQueryIntMany) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := parseNameIDQuery(r, queryParam, pResType, pLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, id)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

func nameIDQueryWrapper[T, P seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, pResType ResTypeSingle, pLookup map[string]P, wrapperFn QueryWrapInt) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		id, err := parseNameIDQuery(r, queryParam, pResType, pLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return wrapperFn(cfg, r, ctx, id)
	}
}
