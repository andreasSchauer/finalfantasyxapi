package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaApiResourceList struct {
	ListParams
	Results []AreaAPIResource `json:"results"`
}

func (l AreaApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l AreaApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type AreaAPIResource struct {
	AreaID int32 `json:"-"`
	LocationArea
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r AreaAPIResource) IsZero() bool {
	return r.Area == ""
}

func (r AreaAPIResource) GetID() int32 {
	return r.AreaID
}

func (r AreaAPIResource) GetURL() string {
	return r.URL
}

func (r AreaAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r AreaAPIResource) Error() string {
	return fmt.Sprintf("area api resource: %s, url: %s", r.LocationArea, r.URL)
}

func (r AreaAPIResource) GetAPIResource() APIResource {
	return r
}

func idToAreaAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], id int32) AreaAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID)
	return areaToAreaAPIResource(cfg, i, res)
}

// useful for id-less locationArea slices retrieved from lookup
func locAreaToAreaAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], area seeding.LocationArea) AreaAPIResource {
	res, _ := seeding.GetResource(area, i.objLookup)
	return areaToAreaAPIResource(cfg, i, res)
}

func locAreasToAreaAPIResources(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], areas []seeding.LocationArea) []AreaAPIResource {
	resources := []AreaAPIResource{}

	for _, area := range areas {
		res := locAreaToAreaAPIResource(cfg, i, area)
		resources = append(resources, res)
	}

	return resources
}

// shared logic. not indended to be called directly
func areaToAreaAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], area seeding.Area) AreaAPIResource {
	params := area.GetResParamsLocation()

	return AreaAPIResource{
		AreaID: params.AreaID,
		LocationArea: LocationArea{
			Location:    params.Location,
			Sublocation: params.Sublocation,
			Area:        params.Area,
			Version:     params.Version,
		},
		Specification: params.Specification,
		URL:           createResourceURL(cfg, i.endpoint, params.AreaID),
	}
}

func newAreaAPIResourceList(cfg *Config, r *http.Request, resources []AreaAPIResource) (AreaApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return AreaApiResourceList{}, err
	}

	list := AreaApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
