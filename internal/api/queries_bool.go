package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// db query searches for resources with matching boolean db column value
func boolQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryBool) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), b)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// db query accumulates all resources that fulfill a certain condition. a false boolean flips these results.
func boolQuery2[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryNoInput) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context())
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	if !b {
		resources = removeResources(inputRes, resources)
	}

	return resources, nil
}

func boolQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, bool) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := wrapperFn(cfg, r, b)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}