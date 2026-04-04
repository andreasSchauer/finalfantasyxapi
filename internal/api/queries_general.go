package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


// used for query filters that can't really be generalized. this one simply checks, if it's empty and then calls the wrapperFn
func basicQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, string, QueryType) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return inputRes, nil
	}

	dbIDs, err := wrapperFn(cfg, r, query, queryParam)
	if errors.Is(err, errQueryRedirect) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}
