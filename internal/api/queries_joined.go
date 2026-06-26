package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func joinedQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryNames []QueryParamName, queryFn func(*Config, *http.Request, context.Context) ([]int32, error)) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
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

		return queryFn(cfg, r, ctx)
	}
}