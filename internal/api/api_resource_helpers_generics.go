package api

import (
	"net/http"
)

func (i handlerInput[T, R, A]) toResources(srv *Server, IDs []int32) []A {
	resources := make([]A, len(IDs))

	for idx, id := range IDs {
		resources[idx] = i.idToResFunc(srv, i, id)
	}

	return resources
}

func (i handlerInput[T, R, A]) toResList(srv *Server, r *http.Request, IDs []int32) (APIResourceList[A], error) {
	resources := i.toResources(srv, IDs)

	resourceList, err := newAPIResourceList(srv, r, resources)
	if err != nil {
		return APIResourceList[A]{}, err
	}

	return resourceList, nil
}
