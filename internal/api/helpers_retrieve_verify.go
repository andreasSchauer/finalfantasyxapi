package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func verifyParamsAndGet[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32) (T, error) {
	var zeroType T

	err := verifyQueryParams(cfg, r, i, &id, nil)
	if err != nil {
		return zeroType, err
	}

	resource, err := seeding.GetResourceByID(id, i.objLookupID)
	if err != nil {
		return zeroType, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist.", i.resTypeSing, id), err)
	}

	return resource, nil
}

func verifyParamsAndRetrieve[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]int32, error) {
	err := verifyQueryParams(cfg, r, i, nil, nil)
	if err != nil {
		return nil, err
	}

	dbIDs, err := i.retrieveQuery(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s.", i.resTypePlural), err)
	}

	return dbIDs, nil
}
