package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaAPIResourceList struct {
	ListParams
	Results []AreaAPIResource `json:"results"`
}

func (l AreaAPIResourceList) GetListParams() ListParams {
	return l.ListParams
}

func (l AreaAPIResourceList) GetResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type AreaAPIResource struct {
	AreaID int32 `json:"-"`
	seeding.LocationArea
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

func (r AreaAPIResource) GetTestKey() string {
	return h.NameToString(r.Area, r.Version, nil)
}

func (r AreaAPIResource) String() string {
	return fmt.Sprintf("area api resource: %s, url: %s", r.LocationArea, r.URL)
}

func (r AreaAPIResource) GetAPIResource() APIResource {
	return r
}

func idToAreaAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaAPIResourceList], id int32) AreaAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID)
	return areaToAreaAPIResource(cfg, i, res)
}

func idPtrToAreaAPIResPtr(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaAPIResourceList], idPtr *int32) *AreaAPIResource {
	if idPtr == nil {
		return nil
	}

	res := idToAreaAPIResource(cfg, i, *idPtr)
	return &res
}

func locAreaToAreaAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaAPIResourceList], area seeding.LocationArea) AreaAPIResource {
	res, _ := seeding.GetResource(area, i.objLookup)
	return areaToAreaAPIResource(cfg, i, res)
}

func locAreasToAreaAPIResources(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaAPIResourceList], areas []seeding.LocationArea) []AreaAPIResource {
	resources := []AreaAPIResource{}

	for _, area := range areas {
		res := locAreaToAreaAPIResource(cfg, i, area)
		resources = append(resources, res)
	}

	return resources
}

// shared logic. not indended to be called directly
func areaToAreaAPIResource(cfg *Config, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaAPIResourceList], area seeding.Area) AreaAPIResource {
	params := area.GetResParamsLocation()

	return AreaAPIResource{
		AreaID: params.AreaID,
		LocationArea: seeding.LocationArea{
			Location:    params.Location,
			Sublocation: params.Sublocation,
			Area:        params.Area,
			Version:     params.Version,
		},
		Specification: params.Specification,
		URL:           CreateResourceURL(cfg, i.Endpoint, params.AreaID),
	}
}

func newAreaAPIResourceList(cfg *Config, r *http.Request, resources []AreaAPIResource) (AreaAPIResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return AreaAPIResourceList{}, err
	}

	list := AreaAPIResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
