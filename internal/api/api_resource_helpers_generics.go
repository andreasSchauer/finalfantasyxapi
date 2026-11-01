package api

import (
	"net/http"
)

func (i handlerInput[T, R, A]) toResources(cfg *Config, IDs []int32) []A {
	resources := make([]A, len(IDs))

	for idx, id := range IDs {
		resources[idx] = i.idToResFunc(cfg, i, id)
	}

	return resources
}

func (i handlerInput[T, R, A]) toResList(cfg *Config, r *http.Request, IDs []int32) (APIResourceList[A], error) {
	resources := i.toResources(cfg, IDs)

	resourceList, err := newAPIResourceList(cfg, r, resources)
	if err != nil {
		return APIResourceList[A]{}, err
	}

	return resourceList, nil
}
