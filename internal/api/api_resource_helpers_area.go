package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToAreaAPIResource(i handlerInput[seedtypes.Area, Area, apiresource.AreaAPIResource], id int32) apiresource.AreaAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID)
	return areaToAreaAPIResource(i, res)
}

func idPtrToAreaAPIResPtr(i handlerInput[seedtypes.Area, Area, apiresource.AreaAPIResource], idPtr *int32) *apiresource.AreaAPIResource {
	if idPtr == nil {
		return nil
	}

	res := idToAreaAPIResource(i, *idPtr)
	return &res
}

func locAreaToAreaAPIResource(i handlerInput[seedtypes.Area, Area, apiresource.AreaAPIResource], area seedtypes.LocationArea) apiresource.AreaAPIResource {
	res, _ := lookup.GetResource(area, i.objLookup)
	return areaToAreaAPIResource(i, res)
}

func locAreasToAreaAPIResources(i handlerInput[seedtypes.Area, Area, apiresource.AreaAPIResource], areas []seedtypes.LocationArea) []apiresource.AreaAPIResource {
	resources := []apiresource.AreaAPIResource{}

	for _, area := range areas {
		res := locAreaToAreaAPIResource(i, area)
		resources = append(resources, res)
	}

	return resources
}

// shared logic. not indended to be called directly
func areaToAreaAPIResource(i handlerInput[seedtypes.Area, Area, apiresource.AreaAPIResource], area seedtypes.Area) apiresource.AreaAPIResource {
	params := area.GetResParamsLocation()

	return apiresource.AreaAPIResource{
		AreaID: params.AreaID,
		LocationArea: seedtypes.LocationArea{
			Location:    params.Location,
			Sublocation: params.Sublocation,
			Area:        params.Area,
			Version:     params.Version,
		},
		Specification: params.Specification,
		URL:           ah.CreateResourceURL(i.Endpoint, params.AreaID),
	}
}
