package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type AreaAPIResource struct {
	AreaID int32 `json:"-"`
	seedtypes.LocationArea
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

func idToAreaAPIResource(srv *Server, i handlerInput[seedtypes.Area, Area, AreaAPIResource], id int32) AreaAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID)
	return areaToAreaAPIResource(srv, i, res)
}

func idPtrToAreaAPIResPtr(srv *Server, i handlerInput[seedtypes.Area, Area, AreaAPIResource], idPtr *int32) *AreaAPIResource {
	if idPtr == nil {
		return nil
	}

	res := idToAreaAPIResource(srv, i, *idPtr)
	return &res
}

func locAreaToAreaAPIResource(srv *Server, i handlerInput[seedtypes.Area, Area, AreaAPIResource], area seedtypes.LocationArea) AreaAPIResource {
	res, _ := lookup.GetResource(area, i.objLookup)
	return areaToAreaAPIResource(srv, i, res)
}

func locAreasToAreaAPIResources(srv *Server, i handlerInput[seedtypes.Area, Area, AreaAPIResource], areas []seedtypes.LocationArea) []AreaAPIResource {
	resources := []AreaAPIResource{}

	for _, area := range areas {
		res := locAreaToAreaAPIResource(srv, i, area)
		resources = append(resources, res)
	}

	return resources
}

// shared logic. not indended to be called directly
func areaToAreaAPIResource(srv *Server, i handlerInput[seedtypes.Area, Area, AreaAPIResource], area seedtypes.Area) AreaAPIResource {
	params := area.GetResParamsLocation()

	return AreaAPIResource{
		AreaID: params.AreaID,
		LocationArea: seedtypes.LocationArea{
			Location:    params.Location,
			Sublocation: params.Sublocation,
			Area:        params.Area,
			Version:     params.Version,
		},
		Specification: params.Specification,
		URL:           CreateResourceURL(srv, i.Endpoint, params.AreaID),
	}
}
