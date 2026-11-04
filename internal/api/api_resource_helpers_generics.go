package api

import (
	"net/http"
)

func (i handlerInput[T, R, A]) toResources(server *Server, IDs []int32) []A {
	resources := make([]A, len(IDs))

	for idx, id := range IDs {
		resources[idx] = i.idToResFunc(server, i, id)
	}

	return resources
}

func (i handlerInput[T, R, A]) toResList(server *Server, r *http.Request, IDs []int32) (APIResourceList[A], error) {
	resources := i.toResources(server, IDs)

	resourceList, err := newAPIResourceList(server, r, resources)
	if err != nil {
		return APIResourceList[A]{}, err
	}

	return resourceList, nil
}
