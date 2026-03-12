package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
