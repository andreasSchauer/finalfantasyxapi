package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

// query uses a list of names or ids as database input to filter for resources. alternatively, "none" can be used as input instead.

func (i Main[T, R, A]) QueryNameIDListNul[P seedtypes.Lookupable](r *http.Request, inputIDs []int32, queryName ts.QueryParamName, pResType ts.ResTypeSingle, pLookup map[string]P, dbQuery DBQueryIntList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
			return inputIDs, nil
		}

		queryIDs, err := queryParam.ParseNameIDList(r, pResType, pLookup)
		if errs.QueryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, queryIDs)
		if err != nil {
			return nil, i.ErrorDBFilter(queryParam, err)
		}

		return dbIDs, nil
	}
}
