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
	return fmt.Sprintf("location based api resource: %s, url: %s", r.LocationArea, r.URL)
}

func (r LocationAPIResource) GetAPIResource() APIResource {
	return r
}


func idToLocationAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList], id int32) LocationAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID)
	return areaToLocationResource(cfg, i, res)
}

// useful for id-less locationArea slices retrieved from lookup
func locAreaToLocationAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList], area seeding.LocationArea) LocationAPIResource {
	res, _ := seeding.GetResource(area, i.objLookup)
	return areaToLocationResource(cfg, i, res)
}


// shared logic. not indended to be called directly
func areaToLocationResource(cfg *Config, i handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList], area seeding.Area) LocationAPIResource {
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
		URL:           		createResourceURL(cfg, i.endpoint, params.AreaID),
	}
}


func newLocationAPIResourceList(cfg *Config, r *http.Request, resources []LocationAPIResource) (LocationApiResourceList, error) {
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