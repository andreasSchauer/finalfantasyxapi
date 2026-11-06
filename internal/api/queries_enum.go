package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// query uses an enum type (id or string possible) that needs to be checked for validity and then returns all resources matching that type
func (i handlerInput[T, R, A]) queryEnum[E, N any](r *http.Request, et enums.EnumType[E, N], inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryEnum[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		enum, err := queryParam.ParseEnum(r, et)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		typedStr := et.ConvFunc(enum.Name)

		dbIDs, err := dbQuery(ctx, typedStr)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}

// like enum query, but with more specialized logic in between (wrapperFn). For example, if types are grouped together (ctbIconType)
func (i handlerInput[T, R, A]) queryEnumWrapper[E, N any](r *http.Request, et enums.EnumType[E, N], inputIDs []int32, queryName ts.QueryParamName, wrapperFn QueryWrapEnum[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		enum, err := queryParam.ParseEnum(r, et)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		typedStr := et.ConvFunc(enum.Name)

		return wrapperFn(r, ctx, typedStr)
	}
}
