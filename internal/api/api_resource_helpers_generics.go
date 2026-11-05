package api

import (
	"net/http"
)

func (i handlerInput[T, R, A]) toResources(IDs []int32) []A {
	resources := make([]A, len(IDs))

	for idx, id := range IDs {
		resources[idx] = i.idToResFunc(i, id)
	}

	return resources
}

func (i handlerInput[T, R, A]) toResList(r *http.Request, IDs []int32) (APIResourceList[A], error) {
	resources := i.toResources(IDs)

	resourceList, err := newAPIResourceList(r, resources)
	if err != nil {
		return APIResourceList[A]{}, err
	}

	return resourceList, nil
}
