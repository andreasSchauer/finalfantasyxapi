package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocationdApiResourceList struct {
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

type LocationArea struct {
	Location    string `json:"location"`
	SubLocation string `json:"sublocation"`
	Area        string `json:"area"`
	Version     *int32 `json:"version,omitempty"`
}

func (la LocationArea) ToKeyFields() []any {
	return []any{
		la.Location,
		la.SubLocation,
		la.Area,
		h.DerefOrNil(la.Version),
	}
}

func (la LocationArea) Error() string {
	return fmt.Sprintf("location area with location: %s, sublocation: %s, area: %s, version: %v", la.Location, la.SubLocation, la.Area, h.DerefOrNil(la.Version))
}

func newLocationArea(location, sublocation, area string, version *int32) LocationArea {
	return LocationArea{
		Location:    location,
		SubLocation: sublocation,
		Area:        area,
		Version:     version,
	}
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


func (cfg *apiConfig) newLocationAPIResourceList(r *http.Request, resources []LocationAPIResource) (LocationdApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return LocationdApiResourceList{}, err
	}

	list := LocationdApiResourceList{
		ListParams: listParams,
		Results: 	shownResources,
	}

	return list, nil
}