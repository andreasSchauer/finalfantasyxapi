package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToAreaAPIResource(i handlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource], id int32) apiresource.AreaAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID)
	return areaToAreaAPIResource(i, res)
}

func idPtrToAreaAPIResPtr(i handlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource], idPtr *int32) *apiresource.AreaAPIResource {
	if idPtr == nil {
		return nil
	}

	res := idToAreaAPIResource(i, *idPtr)
	return &res
}

func locAreaToAreaAPIResource(i handlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource], area seedtypes.LocationArea) apiresource.AreaAPIResource {
	res, _ := lookup.GetResource(area, i.objLookup)
	return areaToAreaAPIResource(i, res)
}

func locAreasToAreaAPIResources(i handlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource], areas []seedtypes.LocationArea) []apiresource.AreaAPIResource {
	resources := make([]apiresource.AreaAPIResource, len(areas))

	for idx, area := range areas {
		resources[idx] = locAreaToAreaAPIResource(i, area)
	}

	return resources
}

// shared logic. not indended to be called directly
func areaToAreaAPIResource(i handlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource], area seedtypes.Area) apiresource.AreaAPIResource {
	params := area.GetResParamsLocation()

	return apiresource.AreaAPIResource{
		AreaID: params.AreaID,
		Location:    params.Location,
		Sublocation: params.Sublocation,
		Area:        params.Area,
		Version:     params.Version,
		Specification: params.Specification,
		URL:           ah.CreateResourceURL(i.Endpoint, params.AreaID),
	}
}
