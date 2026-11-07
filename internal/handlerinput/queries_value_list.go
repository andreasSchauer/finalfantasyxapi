package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i HandlerInput[T, R, A]) QueryValueList(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryStringList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		values, err := queryParam.ParseValueList(r)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, values)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}
