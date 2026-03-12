package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func verifyParamsAndGet[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32) (T, error) {
	var zeroType T

	err := verifyQueryParams(cfg, r, i, &id, nil)
	if err != nil {
		return zeroType, err
	}

	resource, err := seeding.GetResourceByID(id, i.objLookupID)
	if err != nil {
		return zeroType, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist.", i.resourceType, id), err)
	}

	return resource, nil
}

func verifyParamsAndRetrieve[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]int32, error) {
	err := verifyQueryParams(cfg, r, i, nil, nil)
	if err != nil {
		return nil, err
	}

	dbIDs, err := i.retrieveQuery(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss.", i.resourceType), err)
	}

	return dbIDs, nil
}