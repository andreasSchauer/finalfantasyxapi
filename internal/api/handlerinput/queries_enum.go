package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

// query uses an enum type (id or string possible) that needs to be checked for validity and then returns all resources matching that type

func (i Main[T, R, A]) QueryEnum[E, N any](r *http.Request, et enums.EnumType[E, N], inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryEnum[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
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
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}

// like enum query, but with more specialized logic in between (wrapperFn). For example, if types are grouped together (ctbIconType)

func (i Main[T, R, A]) QueryEnumWrapper[E, N any](r *http.Request, et enums.EnumType[E, N], inputIDs []int32, queryName ts.QueryParamName, wrapperFn QueryWrapEnum[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
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
