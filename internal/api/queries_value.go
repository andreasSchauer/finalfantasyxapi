package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)



func valueQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryStringMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	if replParamsPresent(r, queryParam, i.queryLookup) {
		return inputRes, nil
	}

	value, err := parseValueQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), value)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}