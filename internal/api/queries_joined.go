package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func joinedQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryNames []string, queryFn func(*Config, *http.Request) ([]int32, error)) ([]A, error) {
	allEmpty := true

	for _, queryName := range queryNames {
		if replParamsPresent(r, i.queryLookup[queryName], i.queryLookup) {
			return inputRes, nil
		}
		_, err := checkEmptyQuery(r, i.queryLookup[queryName])
		if !queryIsEmpty(err) {
			allEmpty = false
			break
		}
	}

	if allEmpty {
		return inputRes, nil
	}

	ids, err := queryFn(cfg, r)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, ids)
	return resources, nil
}