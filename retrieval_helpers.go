package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func verifyParamsAndGet[T h.HasID, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], id int32) (T, error) {
	var zeroType T

	err := verifyQueryParams(r, i, &id)
	if err != nil {
		return zeroType, err
	}

	resource, err := seeding.GetResourceByID(id, i.objLookupID)
	if err != nil {
		return zeroType, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist.", i.resourceType, id), err)
	}

	return resource, nil
}


func getMultipleAPIResources[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], name string) (L, error) {
	var zeroType L

	dbIDs, err := i.getMultipleQuery(r.Context(), name)
	if err != nil {
		return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'.", i.resourceType, name), err)
	}

	return idsToAPIResourceList(cfg, r, i, dbIDs)
}


func retrieveAPIResources[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]A, error) {
	dbIDs, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return idsToAPIResources(cfg, i, dbIDs), nil
}


func verifyParamsAndRetrieve[T h.HasID, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L]) ([]int32, error) {
	err := verifyQueryParams(r, i, nil)
	if err != nil {
		return nil, err
	}

	dbIDs, err := i.retrieveQuery(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss.", i.resourceType), err)
	}

	return dbIDs, nil
}


func filterAPIResources[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []A, filteredLists []filteredResList[A]) (L, error) {
	var zeroType L
	
	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return zeroType, filtered.err
		}
		resources = getSharedResources(resources, filtered.resources)
	}

	resourceList, err := i.resToListFunc(cfg, r, i, resources)
	if err != nil {
		return zeroType, err
	}

	return resourceList, nil
}


