package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// query uses a list of names or ids as database input to filter for resources. alternatively, "none" can be used as input instead.
func nameIdListQueryNul[T, P h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName, pResType string, pLookup map[string]P, dbQuery DbQueryIntList) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	queryIDs, err := parseNameIdListQuery(r, queryParam, pResType, pLookup)
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