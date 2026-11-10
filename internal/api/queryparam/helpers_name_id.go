package queryparam

import (
	"errors"
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

// ParseNameID validates name/id-queryParam and checks emptiness. returns the corresponding id.
func (q QueryParam) ParseNameID[P seedtypes.Lookupable](r *http.Request, pResType ts.ResTypeSingle, pLookup map[string]P) (int32, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return 0, err
	}

	return q.CheckNameID(query, pResType, pLookup)
}

// CheckNameID validates an id or single-segment-resource name, checks emptiness, and returns the corresponding id.
func (q QueryParam) CheckNameID[P seedtypes.Lookupable](query string, pResType ts.ResTypeSingle, pLookup map[string]P) (int32, error) {
	id, err := q.CheckID(query, len(pLookup))
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, errs.ErrNotAnID) {
		return 0, err
	}

	resource, err := ah.CheckUniqueName(query, pLookup)
	if err == nil {
		return resource.ID, nil
	}

	return 0, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown %s '%s' used for parameter '%s'", pResType, query, q.Name), err)
}
