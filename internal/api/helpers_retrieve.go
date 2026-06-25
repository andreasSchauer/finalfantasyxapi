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

func filterIDs[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], IDs []int32, filteredLists []filteredIdList) ([]int32, error) {
	filteredIDs := IDs

	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return nil, filtered.err
		}
		filteredIDs = getSharedIDs(filteredIDs, filtered.IDs)
	}

	if i.avlFunc != nil {
		var err error
		filteredIDs, err = i.avlFunc(cfg, r, filteredIDs)
		if err != nil {
			return nil, err
		}
	}

	flip, err := parseBooleanQuery(r, i.queryLookup[qpnFlip])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	if flip {
		filteredIDs = removeIDs(IDs, filteredIDs)
	}

	return filteredIDs, nil
}
