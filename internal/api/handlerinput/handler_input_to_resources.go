package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func (i Main[T, R, A]) ToAPIResources(IDs []int32) []A {
	resources := make([]A, len(IDs))

	for idx, id := range IDs {
		resources[idx] = i.IDToResFunc(i, id)
	}

	return resources
}

func (i Main[T, R, A]) ToAPIResList(r *http.Request, IDs []int32) (apiresource.APIResourceList[A], error) {
	resources := i.ToAPIResources(IDs)

	resourceList, err := NewAPIResourceList(r, resources)
	if err != nil {
		return apiresource.APIResourceList[A]{}, err
	}

	return resourceList, nil
}
