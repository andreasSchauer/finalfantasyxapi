package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


// query uses an integer value as input.
func intQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	integer, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), int32(integer))
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// query uses an integer value as input.
func intQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, int32) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	integer, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := wrapperFn(cfg, r, int32(integer))
	if errors.Is(err, errQueryRedirect) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}