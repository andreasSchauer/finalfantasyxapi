package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// validates name/id-queryParam and checks emptiness. returns the corresponding id.
func (q QueryParam) parseNameID[P seedtypes.Lookupable](r *http.Request, pResType typestrings.ResTypeSingle, pLookup map[string]P) (int32, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return 0, err
	}

	return q.checkNameID(query, pResType, pLookup)
}

// validates an id or single-segment-resource name, checks emptiness, and returns the corresponding id.
func (q QueryParam) checkNameID[P seedtypes.Lookupable](query string, pResType typestrings.ResTypeSingle, pLookup map[string]P) (int32, error) {
	id, err := q.checkID(query, len(pLookup))
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

	return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown %s '%s' used for parameter '%s'", pResType, query, q.Name), err)
}
