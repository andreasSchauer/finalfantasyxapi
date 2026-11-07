package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i HandlerInput[T, R, A]) QueryEnumList[E, N any](r *http.Request, et enums.EnumType[E, N], inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryEnumList[E]) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		enums, err := queryParam.ParseEnumList(r, et)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, enums)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}
