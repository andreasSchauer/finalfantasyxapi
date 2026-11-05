package api

import (
	"net/http"
)

func (i handlerInput[T, R, A]) toAPIResources(IDs []int32) []A {
	resources := make([]A, len(IDs))

	for idx, id := range IDs {
		resources[idx] = i.idToResFunc(i, id)
	}

	return resources
}

func (i handlerInput[T, R, A]) toAPIResList(r *http.Request, IDs []int32) (APIResourceList[A], error) {
	resources := i.toAPIResources(IDs)

	resourceList, err := newAPIResourceList(r, resources)
	if err != nil {
		return APIResourceList[A]{}, err
	}

	return resourceList, nil
}
