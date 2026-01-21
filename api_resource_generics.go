package main

import (
	"context"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func idsToAPIResources[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], IDs []int32) []A {
	resources := []A{}

	for _, id := range IDs {
		resource := i.idToResFunc(cfg, i, id)
		resources = append(resources, resource)
	}

	return resources
}

func idsToAPIResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], dbIDs []int32) (L, error) {
	var zeroType L

	resources := idsToAPIResources(cfg, i, dbIDs)
	
	resourceList, err := i.resToListFunc(cfg, r, resources)
	if err != nil {
		return zeroType, err
	}

	return resourceList, nil
}

// get relationship resources of item. handlerInput = endpoint of fetched resources
func getResourcesDB[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item seeding.LookupableID, dbQuery func(context.Context, int32) ([]int32, error)) ([]A, error) {
	dbIds, err := dbQuery(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %ss of %s.", i.resourceType, item), err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

// filter resources by item id. handlerInput = endpoint of resources
func filterResourcesDB[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, lookupType string, dbQuery func(context.Context, int32) ([]int32, error)) ([]A, error) {
	dbIds, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by %s id '%d.", i.resourceType, lookupType, id), err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}
