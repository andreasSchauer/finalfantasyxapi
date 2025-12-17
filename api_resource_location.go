package main

import (
	"fmt"
	"net/http"

	
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocationApiResourceList struct {
	ListParams
	Results []LocationAPIResource `json:"results"`
}

type LocationAPIResource struct {
	AreaID int32 `json:"-"`
	LocationArea
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r LocationAPIResource) IsZero() bool {
	return r.Area == ""
}

func (r LocationAPIResource) getID() int32 {
	return r.AreaID
}

func (r LocationAPIResource) getURL() string {
	return r.URL
}

func (r LocationAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r LocationAPIResource) Error() string {
	return fmt.Sprintf("location based api resource: %s, url: %s", r.LocationArea.Error(), r.URL)
}

func (r LocationAPIResource) getAPIResource() IsAPIResource {
	return r
}

func (cfg *apiConfig) newLocationBasedAPIResource(area LocationArea) LocationAPIResource {
	areaLookup, _ := seeding.GetResource(area, cfg.l.Areas)

	return LocationAPIResource{
		AreaID:        areaLookup.ID,
		LocationArea:  area,
		Specification: areaLookup.Specification,
		URL:           cfg.createURL("areas", areaLookup.ID),
	}
}

func createLocationBasedAPIResources[T any](
	cfg *apiConfig,
	items []T,
	mapper func(T) (location, sublocation, area string, version *int32),
) []LocationAPIResource {
	resources := []LocationAPIResource{}

	for _, item := range items {
		location, sublocation, area, version := mapper(item)
		locationArea := newLocationArea(location, sublocation, area, version)
		resource := cfg.newLocationBasedAPIResource(locationArea)

		if !resource.IsZero() {
			resources = append(resources, resource)
		}
	}

	return resources
}

func (cfg *apiConfig) newLocationAPIResourceList(r *http.Request, resources []LocationAPIResource) (LocationApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	list := LocationApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
