package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type UnnamedApiResourceList struct {
	ListParams
	Results []UnnamedAPIResource `json:"results"`
}

func (l UnnamedApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l UnnamedApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type UnnamedAPIResource struct {
	ID  int32  `json:"id"`
	URL string `json:"url"`
}

func (r UnnamedAPIResource) IsZero() bool {
	return r.ID == 0
}

func (r UnnamedAPIResource) GetID() int32 {
	return r.ID
}

func (r UnnamedAPIResource) GetURL() string {
	return r.URL
}

func (r UnnamedAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r UnnamedAPIResource) Error() string {
	return fmt.Sprintf("unnamed api resource with url: %s", r.URL)
}

func (r UnnamedAPIResource) GetAPIResource() APIResource {
	return r
}

func (cfg *Config) newUnnamedAPIResource(endpoint string, id int32) UnnamedAPIResource {
	if id == 0 {
		return UnnamedAPIResource{}
	}

	return UnnamedAPIResource{
		ID:  id,
		URL: cfg.createResourceURL(endpoint, id),
	}
}

func createUnnamedAPIResources[T any](
	cfg *Config,
	items []T,
	endpoint string,
	mapper func(T) (id int32),
) []UnnamedAPIResource {
	resources := []UnnamedAPIResource{}

	for _, item := range items {
		id := mapper(item)
		resource := cfg.newUnnamedAPIResource(endpoint, id)

		if !resource.IsZero() {
			resources = append(resources, resource)
		}
	}

	return resources
}

func newUnnamedAPIResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []UnnamedAPIResource) (UnnamedApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, i, resources)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	list := UnnamedApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}


func idToUnnamedAPIResource[T h.IsUnnamed, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], id int32) UnnamedAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seeding
	params := res.GetResParamsUnnamed()

	return UnnamedAPIResource{
		ID:            params.ID,
		URL:           cfg.createResourceURL(i.endpoint, params.ID),
	}
}