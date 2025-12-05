package main

import (
	"fmt"
	"net/http"
)


type UnnamedApiResourceList struct {
	ListParams
	Results []UnnamedAPIResource `json:"results"`
}

type UnnamedAPIResource struct {
	ID            int32   `json:"id"`
	URL           string  `json:"url"`
}

func (r UnnamedAPIResource) IsZero() bool {
	return r.ID == 0
}

func (r UnnamedAPIResource) getID() int32 {
	return r.ID
}

func (r UnnamedAPIResource) getURL() string {
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

func (r UnnamedAPIResource) getAPIResource() IsAPIResource {
	return r
}


func (cfg *apiConfig) newUnnamedAPIResource(endpoint string, id int32) UnnamedAPIResource {
	if id == 0 {
		return UnnamedAPIResource{}
	}

	return UnnamedAPIResource{
		ID:            id,
		URL:           cfg.createURL(endpoint, id),
	}
}


func createUnnamedAPIResources[T any](
	cfg *apiConfig,
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


func (cfg *apiConfig) newUnnamedAPIResourceList(r *http.Request, resources []UnnamedAPIResource) (UnnamedApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	list := UnnamedApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
