package api

import (
	"errors"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// validates name/id-queryParam and checks emptiness. returns the corresponding id.
func parseNameIdQuery[P h.HasID](r *http.Request, queryParam QueryType, pResType string, pLookup map[string]P) (int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	return checkQueryNameID(query, pResType, queryParam, pLookup)
}

// validates an id or single-segment-resource name, checks emptiness, and returns the corresponding id.
func checkQueryNameID[P h.HasID](query, pResType string, queryParam QueryType, pLookup map[string]P) (int32, error) {
	id, err := checkQueryID(query, queryParam, len(pLookup))
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, errNotAnID) {
		return 0, err
	}

	resource, err := checkUniqueName(query, pLookup)
	if err == nil {
		return resource.ID, nil
	}

	return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown %s '%s' used for parameter '%s'.", pResType, query, queryParam.Name), err)
}
