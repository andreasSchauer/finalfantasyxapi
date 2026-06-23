package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMultipleAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], name string) (L, error) {
	var zeroType L

	dbIDs, err := i.getMultipleQuery(r.Context(), name)
	if err != nil {
		return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'.", i.resTypeSing, name), err)
	}

	return idsToAPIResourceList(cfg, r, i, dbIDs)
}

func retrieveAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]A, error) {
	dbIDs, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

func filterIDs[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], IDs []int32, filteredLists []filteredIdList) (L, error) {
	var zeroType L
	filteredIDs := IDs

	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return zeroType, filtered.err
		}
		filteredIDs = getSharedIDs(filteredIDs, filtered.IDs)
	}

	if i.avlFunc != nil {
		var err error
		filteredIDs, err = i.avlFunc(cfg, r, filteredIDs)
		if err != nil {
			return zeroType, err
		}
	}

	flip, err := parseBooleanQuery(r, i.queryLookup[qpnFlip])
	if errExceptEmptyQuery(err) {
		return zeroType, err
	}

	if flip {
		filteredIDs = removeIDs(IDs, filteredIDs)
	}

	resources := idsToAPIResources(cfg, i, filteredIDs)

	resourceList, err := i.resToListFunc(cfg, r, resources)
	if err != nil {
		return zeroType, err
	}

	return resourceList, nil
}
