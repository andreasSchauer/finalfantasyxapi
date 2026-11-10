package handlerinput

import (
	"context"
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func (i Main[T, R, A]) QueryWrapper(r *http.Request, inputIDs []int32, queryName ts.QueryParamName, wrapperFn QueryWrapBasic) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.QueryLookup[queryName]
		query, err := queryParam.BasicChecks(r, i.QueryLookup)
		if err != nil {
			return inputIDs, nil
		}

		dbIDs, err := wrapperFn(r, ctx, query, queryParam)
		if errors.Is(err, errs.ErrQueryRedirect) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return dbIDs, nil
	}
}
