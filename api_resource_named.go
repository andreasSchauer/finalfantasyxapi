package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type NamedApiResourceList struct {
	ListParams
	Results []NamedAPIResource `json:"results"`
}

func (l NamedApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l NamedApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type NamedAPIResource struct {
	ID            int32   `json:"-"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r NamedAPIResource) IsZero() bool {
	return r.Name == ""
}

func (r NamedAPIResource) getID() int32 {
	return r.ID
}

func (r NamedAPIResource) getURL() string {
	return r.URL
}

func (r NamedAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r NamedAPIResource) Error() string {
	return fmt.Sprintf("named api resource %s, version: %v, url: %s", r.Name, h.DerefOrNil(r.Version), r.URL)
}

func (r NamedAPIResource) getAPIResource() IsAPIResource {
	return r
}

func (cfg *Config) newNamedAPIResource(endpoint string, id int32, name string, version *int32, spec *string) NamedAPIResource {
	if name == "" {
		return NamedAPIResource{}
	}

	return NamedAPIResource{
		ID:            id,
		Name:          name,
		Version:       version,
		Specification: spec,
		URL:           cfg.createURL(endpoint, id),
	}
}

func (cfg *Config) newNamedAPIResourceSimple(endpoint string, id int32, name string) NamedAPIResource {
	if name == "" {
		return NamedAPIResource{}
	}

	return NamedAPIResource{
		ID:   id,
		Name: name,
		URL:  cfg.createURL(endpoint, id),
	}
}

func createNamedAPIResources[T any](
	cfg *Config,
	items []T,
	endpoint string,
	mapper func(T) (id int32, name string, version *int32, spec *string),
) []NamedAPIResource {
	resources := []NamedAPIResource{}

	for _, item := range items {
		id, name, version, spec := mapper(item)
		resource := cfg.newNamedAPIResource(endpoint, id, name, version, spec)

		if !resource.IsZero() {
			resources = append(resources, resource)
		}
	}

	return resources
}

func createNamedAPIResourcesSimple[T any](
	cfg *Config,
	items []T,
	endpoint string,
	mapper func(T) (id int32, name string),
) []NamedAPIResource {
	resources := []NamedAPIResource{}

	for _, item := range items {
		id, name := mapper(item)
		resource := cfg.newNamedAPIResourceSimple(endpoint, id, name)

		if !resource.IsZero() {
			resources = append(resources, resource)
		}
	}

	return resources
}

func (cfg *Config) newNamedAPIResourceList(r *http.Request, resources []NamedAPIResource) (NamedApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	list := NamedApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
