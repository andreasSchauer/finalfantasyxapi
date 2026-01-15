package main

import (
	"context"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func retrieveNamedAPIResources[T h.IsNamed, R any, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, L]) ([]NamedAPIResource, error) {
	dbIDs, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return idsToNamedAPIResources(cfg, i, dbIDs), nil
}

func retrieveLocationAPIResources(cfg *Config, r *http.Request, i handlerInput[seeding.Area, Area, LocationApiResourceList]) ([]LocationAPIResource, error) {
	dbIDs, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return idsToLocationAPIResources(cfg, i, dbIDs), nil
}

func verifyParamsAndGet[T h.HasID, R any, L APIResourceList](r *http.Request, i handlerInput[T, R, L], id int32) (T, error) {
	var zeroType T

	err := verifyQueryParams(r, i, nil)
	if err != nil {
		return zeroType, err
	}

	resource, err := seeding.GetResourceByID(id, i.objLookupID)
	if err != nil {
		return zeroType, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist.", i.resourceType, id), err)
	}

	return resource, nil
}

func verifyParamsAndRetrieve[T h.HasID, R any, L APIResourceList](r *http.Request, i handlerInput[T, R, L]) ([]int32, error) {
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



func filterNamedAPIResources[T h.IsNamed, R any, L APIResourceList](cfg *Config, r *http.Request,i handlerInput[T, R, L], resources []NamedAPIResource, filteredLists []filteredResList[NamedAPIResource]) (NamedApiResourceList, error) {
	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return NamedApiResourceList{}, filtered.err
		}
		resources = getSharedResources(resources, filtered.resources)
	}

	resourceList, err := newNamedAPIResourceList(cfg, r, i, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return resourceList, nil
}


func filterLocationAPIResources(cfg *Config, r *http.Request,i handlerInput[seeding.Area, Area, LocationApiResourceList], resources []LocationAPIResource, filteredLists []filteredResList[LocationAPIResource]) (LocationApiResourceList, error) {
	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return LocationApiResourceList{}, filtered.err
		}
		resources = getSharedResources(resources, filtered.resources)
	}

	resourceList, err := newLocationAPIResourceList(cfg, r, i, resources)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	return resourceList, nil
}


// used to convert dbQuery Outputs into resources. Should maybe be in their own file
func getNamedResources[T h.IsNamed, R any, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, L], item seeding.LookupableID, dbQuery func (context.Context, int32) ([]int32, error)) ([]NamedAPIResource, error) {
	dbIds, err := dbQuery(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %ss of %s", i.resourceType, item), err)
	}

	resources := idsToNamedAPIResources(cfg, i, dbIds)
	return resources, nil
}


func getUnnamedResources[T h.IsUnnamed, R any, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, L], item seeding.LookupableID, dbQuery func (context.Context, int32) ([]int32, error)) ([]UnnamedAPIResource, error) {
	dbIds, err := dbQuery(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %ss of %s", i.resourceType, item), err)
	}

	resources := idsToUnnamedAPIResources(cfg, i, dbIds)
	return resources, nil
}