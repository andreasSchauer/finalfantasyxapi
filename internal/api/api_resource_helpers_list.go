package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

func newAPIResourceList[A apiresource.APIResource](r *http.Request, resources []A) (apiresource.APIResourceList[A], error) {
	listParams, shownResources, err := CreatePaginatedList(r, resources)
	if err != nil {
		return apiresource.APIResourceList[A]{}, err
	}

	list := apiresource.APIResourceList[A]{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
