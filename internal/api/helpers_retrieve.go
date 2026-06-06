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
		return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'.", i.resourceType, name), err)
	}

	return idsToAPIResourceList(cfg, r, i, dbIDs)
}

func retrieveAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]A, error) {
	dbIDs, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return idsToAPIResources(cfg, i, dbIDs), nil
}

func filterAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []A, filteredLists []filteredResList[A]) (L, error) {
	var zeroType L
	filteredRes := resources

	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return zeroType, filtered.err
		}
		filteredRes = getSharedResources(filteredRes, filtered.resources)
	}

	if i.avlFunc != nil {
		var err error
		filteredRes, err = i.avlFunc(cfg, r, filteredRes)
		if err != nil {
			return zeroType, err
		}
	}

	flip, err := parseBooleanQuery(r, i.queryLookup["flip"])
	if errExceptEmptyQuery(err) {
		return zeroType, err
	}

	if flip {
		filteredRes = removeResources(resources, filteredRes)
	}

	resourceList, err := i.resToListFunc(cfg, r, filteredRes)
	if err != nil {
		return zeroType, err
	}

	return resourceList, nil
}



func resToIDs[A APIResource](resources []A) []int32 {
	ids := []int32{}

	for _, res := range resources {
		ids = append(ids, res.GetID())
	}

	return ids
}
