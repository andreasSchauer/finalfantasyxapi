package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocationApiResourceList struct {
	ListParams
	Results []LocationAPIResource `json:"results"`
}

func (l LocationApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l LocationApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
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

func (r LocationAPIResource) GetID() int32 {
	return r.AreaID
}

func (r LocationAPIResource) GetURL() string {
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

func (r LocationAPIResource) GetAPIResource() APIResource {
	return r
}

func (cfg *Config) newLocationBasedAPIResource(area LocationArea) LocationAPIResource {
	areaLookup, _ := seeding.GetResource(area, cfg.l.Areas)

	return LocationAPIResource{
		AreaID:        areaLookup.ID,
		LocationArea:  area,
		Specification: areaLookup.Specification,
		URL:           cfg.createResourceURL(cfg.e.areas.endpoint, areaLookup.ID),
	}
}

func createLocationBasedAPIResources[T any](
	cfg *Config,
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

func newLocationAPIResourceList[T h.HasID, R any, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, L], resources []LocationAPIResource) (LocationApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, i, resources)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	list := LocationApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}



func idToLocationAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, LocationApiResourceList], id int32) LocationAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seeding
	return areaToLocationResource(cfg, i, res)
}

func locAreaToLocationAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, LocationApiResourceList], area seeding.LocationArea) LocationAPIResource {
	res, _ := seeding.GetResource(area, i.objLookup)
	return areaToLocationResource(cfg, i, res)
}


func areaToLocationResource(cfg *Config, i handlerInput[seeding.Area, Area, LocationApiResourceList], area seeding.Area) LocationAPIResource {
	params := area.GetResParamsLocation()

	return LocationAPIResource{
		AreaID:            	params.AreaID,
		LocationArea: LocationArea{
			Location: 		params.Location,
			Sublocation: 	params.Sublocation,
			Area: 			params.Area,
			Version: 		params.Version,
		},
		Specification: 		params.Specification,
		URL:           		cfg.createResourceURL(i.endpoint, params.AreaID),
	}
}

func idsToLocationAPIResources(cfg *Config, i handlerInput[seeding.Area, Area, LocationApiResourceList], IDs []int32) []LocationAPIResource {
	resources := []LocationAPIResource{}

	for _, id := range IDs {
		resource := idToLocationAPIResource(cfg, i, id)
		resources = append(resources, resource)
	}

	return resources
}