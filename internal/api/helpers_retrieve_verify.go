package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (i handlerInput[T, R, A]) queryParamsVerifyGet(r *http.Request, id int32) (T, error) {
	var zeroType T

	err := verifyQueryParams(r, i.Endpoint, i.queryLookup, &id, nil)
	if err != nil {
		return zeroType, err
	}

	resource, err := seeding.GetResourceByID(id, i.objLookupID)
	if err != nil {
		return zeroType, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist", i.resTypeSingle, id), err)
	}

	return resource, nil
}

func (i handlerInput[T, R, A]) queryParamsVerifyRetrieve(r *http.Request) ([]int32, error) {
	err := verifyQueryParams[int32](r, i.Endpoint, i.queryLookup, nil, nil)
	if err != nil {
		return nil, err
	}

	dbIDs, err := i.retrieveQuery(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s", i.resTypePlural), err)
	}

	return dbIDs, nil
}
