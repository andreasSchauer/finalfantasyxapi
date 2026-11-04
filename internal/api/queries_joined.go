package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) queryJoined(cfg *Config, r *http.Request, inputIDs []int32, queryNames []typestrings.QueryParamName, wrapperFn QueryWrapJoined) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		allEmpty := true

		for _, queryName := range queryNames {
			queryParam := i.queryLookup[queryName]

			if queryParam.replacingParamsPresent(r, i.queryLookup) {
				return inputIDs, nil
			}
			_, err := queryParam.checkEmpty(r)
			if !queryIsEmpty(err) {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			return inputIDs, nil
		}

		return wrapperFn(cfg, r, ctx)
	}
}
