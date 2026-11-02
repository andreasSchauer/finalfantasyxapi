package api

import (
	"context"
	"net/http"
)

// query uses an enum type (id or string possible) that needs to be checked for validity and then returns all resources matching that type
func (i handlerInput[T, R, A]) queryEnum[E, N any](r *http.Request, et EnumType[E, N], inputIDs []int32, queryName QueryParamName, dbQuery DBQueryEnum[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		enum, err := queryParam.parseEnum(r, et)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		typedStr := et.convFunc(enum.Name)

		dbIDs, err := dbQuery(ctx, typedStr)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// like enum query, but with more specialized logic in between (wrapperFn). For example, if types are grouped together (ctbIconType)
func (i handlerInput[T, R, A]) queryEnumWrapper[E, N any](cfg *Config, r *http.Request, et EnumType[E, N], inputIDs []int32, queryName QueryParamName, wrapperFn QueryWrapEnum[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		enum, err := queryParam.parseEnum(r, et)
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
