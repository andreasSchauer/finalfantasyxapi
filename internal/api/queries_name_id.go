package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// query uses the name or id of another resource type to filter resources
func nameOrIdQuery[T, P h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName, pResType string, pLookup map[string]P, dbQuery DbQueryIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseNameIdQuery(r, queryParam, pResType, pLookup)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}
