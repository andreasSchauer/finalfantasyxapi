package handlerinput

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func (i Main[T, R, A]) QueryIntList(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, dbQuery DBQueryIntList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		if queryName == ts.QpnLimit {
			fmt.Println()
		}
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		ints, err := queryParam.ParseIntList(r)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, ints)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}
