package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// query uses an enum type (id or string possible) that needs to be checked for validity and then returns all resources matching that type
func enumQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList, E, N any](r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputIDs []int32, queryName QueryParamName, dbQuery DbQueryEnum[E]) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if replParamsPresent(r, queryParam, i.queryLookup) {
			return inputIDs, nil
		}

		enum, err := parseEnumQuery(r, i.endpoint, queryParam, et)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		typedStr := et.convFunc(enum.Name)

		dbIDs, err := dbQuery(ctx, typedStr)
		if err != nil {
			return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// like enum query, but with more specialized logic in between (wrapperFn). For example, if types are grouped together (ctbIconType)
func enumQueryWrapper[T seeding.Lookupable, R any, A APIResource, L APIResourceList, E, N any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputIDs []int32, queryName QueryParamName, wrapperFn QueryWrapEnum[E]) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		enum, err := parseEnumQuery(r, i.endpoint, queryParam, et)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		typedStr := et.convFunc(enum.Name)

		return wrapperFn(cfg, r, ctx, typedStr)
	}
}
