package handlerinput

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i HandlerInput[T, R, A]) QueryJoined(r *http.Request, inputIDs []int32, queryNames []ts.QueryParamName, wrapperFn QueryWrapJoined) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		allEmpty := true

		for _, queryName := range queryNames {
			queryParam := i.QueryLookup[queryName]

			if queryParam.ReplacingParamsPresent(r, i.QueryLookup) {
				return inputIDs, nil
			}
			_, err := queryParam.CheckEmpty(r)
			if !errs.QueryIsEmpty(err) {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			return inputIDs, nil
		}

		return wrapperFn(r, ctx)
	}
}
