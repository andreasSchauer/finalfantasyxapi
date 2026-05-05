package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getResDbItemOne[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], filterItem seeding.Lookupable, dbQuery DbQueryIntOne) (A, error) {
	var zeroType A

	dbId, err := dbQuery(r.Context(), filterItem.GetID())
	if err != nil {
		return zeroType, newHTTPErrorDB(i.resourceType, filterItem, err)
	}

	res := i.idToResFunc(cfg, i, dbId)
	return res, nil
}

// get relationship resources of item. handlerInput = endpoint of fetched resources
func getResourcesDbItem[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], filterItem seeding.Lookupable, dbQuery DbQueryIntMany) ([]A, error) {
	dbIds, err := dbQuery(r.Context(), filterItem.GetID())
	if err != nil {
		return nil, newHTTPErrorDB(i.resourceType, filterItem, err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

// filter resources by item id. handlerInput = endpoint of fetched resources. lookup type = resourceType of id
func getResourcesDbID[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, lookupType string, dbQuery DbQueryIntMany) ([]A, error) {
	dbIds, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by %s id '%d'.", i.resourceType, lookupType, id), err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

func getResPtrDB[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item seeding.Lookupable, dbQuery DbQueryIntOne) (*A, error) {
	dbID, err := dbQuery(r.Context(), item.GetID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %s of %s.", i.resourceType, item), err)
	}

	res := i.idToResFunc(cfg, i, dbID)
	return &res, nil
}

func dbQueriesToApiResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, pResType string, dbQueryMap map[string]DbQueryIntMany) ([]A, error) {
	resLists := []filteredResList[A]{}

	for key := range dbQueryMap {
		resList := frl(getResourcesDbID(cfg, r, i, id, pResType, dbQueryMap[key]))
		resLists = append(resLists, resList)
	}

	return combineFilteredAPIResources(resLists)
}
