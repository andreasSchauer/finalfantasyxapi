package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// get relationship resources of item. handlerInput = endpoint of fetched resources
func getResourcesDbItem[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], filterItem seeding.LookupableID, dbQuery func(context.Context, int32) ([]int32, error)) ([]A, error) {
	dbIds, err := dbQuery(r.Context(), filterItem.GetID())
	if err != nil {
		return nil, newHTTPErrorDB(i.resourceType, filterItem, err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

// filter resources by item id. handlerInput = endpoint of fetched resources. lookup type = resourceType of id
func getResourcesDbID[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, lookupType string, dbQuery func(context.Context, int32) ([]int32, error)) ([]A, error) {
	dbIds, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by %s id '%d'.", i.resourceType, lookupType, id), err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

func getResPtrDB[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item seeding.LookupableID, dbQuery func(context.Context, int32) (int32, error)) (*A, error) {
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