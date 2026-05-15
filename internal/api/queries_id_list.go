package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// query uses a list of ids as database input to filter for resources
func idListQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, dbQuery DbQueryIntList) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	if replParamsPresent(r, queryParam, i.queryLookup) {
		return inputRes, nil
	}

	queryIDs, err := parseIdListQuery(cfg, r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), queryIDs)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// like idListQuery, but with more specialized logic in between (wrapperFn)
func idListQueryWrapper[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, wrapperFn func(*Config, *http.Request, []int32) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	queryIDs, err := parseIdListQuery(cfg, r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := wrapperFn(cfg, r, queryIDs)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}
