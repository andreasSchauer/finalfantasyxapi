package handlerinput

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func (i HandlerInput[T, R, A]) QueryParamsVerifyGet(r *http.Request, id int32) (T, error) {
	var zeroType T

	err := VerifyQueryParams(r, i.Endpoint, i.queryLookup, &id, nil)
	if err != nil {
		return zeroType, err
	}

	resource, err := lookup.GetResourceByID(id, i.objLookupID)
	if err != nil {
		return zeroType, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist", i.resTypeSingle, id), err)
	}

	return resource, nil
}

func (i HandlerInput[T, R, A]) QueryParamsVerifyRetrieve(r *http.Request) ([]int32, error) {
	err := VerifyQueryParams[int32](r, i.Endpoint, i.queryLookup, nil, nil)
	if err != nil {
		return nil, err
	}

	dbIDs, err := i.retrieveQuery(r.Context())
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s", i.resTypePlural), err)
	}

	return dbIDs, nil
}
