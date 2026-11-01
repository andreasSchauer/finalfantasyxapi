package api

import (
	"context"
	"net/http"
)

func (i handlerInput[T, R, A]) queryJoined(cfg *Config, r *http.Request, inputIDs []int32, queryNames []QueryParamName, wrapperFn QueryWrapJoined) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		allEmpty := true

		for _, queryName := range queryNames {
			if replParamsPresent(r, i.queryLookup[queryName], i.queryLookup) {
				return inputIDs, nil
			}
			_, err := checkEmptyQuery(r, i.queryLookup[queryName])
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
