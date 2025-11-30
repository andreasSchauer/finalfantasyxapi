package main

import (
	"net/http"
)


type NamedApiResourceList struct {
	ListParams
	Results		[]NamedAPIResource	`json:"results"`
}


type NamedAPIResource struct {
	Name			string		`json:"name"`
	Version			*int32		`json:"version,omitempty"`
	Specification	*string		`json:"specification,omitempty"`
	URL				string		`json:"url"`
}


func (cfg *apiConfig) newNamedAPIResourceList(r *http.Request, resources []NamedAPIResource) (NamedApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	list := NamedApiResourceList{
		ListParams: listParams,
		Results: 	shownResources,
	}

	return list, nil
}


func (cfg *apiConfig) newNamedAPIResource(endpoint string, id int32, name string, version *int32, spec *string) NamedAPIResource {
	if name == "" {
		return NamedAPIResource{}
	}

	return NamedAPIResource{
		Name: 			name,
		Version: 		version,
		Specification: 	spec,
		URL: 			cfg.createURL(endpoint, id),
	}
}


func createNamedAPIResources[T any](
	cfg *apiConfig,
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


func (cfg *apiConfig) newNamedAPIResourceSimple(endpoint string, id int32, name string) NamedAPIResource {
	if name == "" {
		return NamedAPIResource{}
	}

	return NamedAPIResource{
		Name: 			name,
		URL: 			cfg.createURL(endpoint, id),
	}
}

func (r NamedAPIResource) IsZero() bool {
	return r.Name == ""
}

func createNamedAPIResourcesSimple[T any](
	cfg *apiConfig,
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