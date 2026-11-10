package handlerinput

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func IDToAreaAPIResource(i Main[seedtypes.Area, apitypes.Area, apiresource.Area], id int32) apiresource.Area {
	res, _ := lookup.GetResourceByID(id, i.ObjLookupID)
	return areaToAreaAPIResource(i, res)
}

func IDPtrToAreaAPIResPtr(i Main[seedtypes.Area, apitypes.Area, apiresource.Area], idPtr *int32) *apiresource.Area {
	if idPtr == nil {
		return nil
	}

	res := IDToAreaAPIResource(i, *idPtr)
	return &res
}

func LocAreaToAreaAPIResource(i Main[seedtypes.Area, apitypes.Area, apiresource.Area], area seedtypes.LocationArea) apiresource.Area {
	res, _ := lookup.GetResource(area, i.ObjLookup)
	return areaToAreaAPIResource(i, res)
}

func LocAreasToAreaAPIResources(i Main[seedtypes.Area, apitypes.Area, apiresource.Area], areas []seedtypes.LocationArea) []apiresource.Area {
	resources := make([]apiresource.Area, len(areas))

	for idx, area := range areas {
		resources[idx] = LocAreaToAreaAPIResource(i, area)
	}

	return resources
}

// shared logic. not indended to be called directly
func areaToAreaAPIResource(i Main[seedtypes.Area, apitypes.Area, apiresource.Area], area seedtypes.Area) apiresource.Area {
	params := area.GetResParamsLocation()

	return apiresource.Area{
		AreaID:        params.AreaID,
		Location:      params.Location,
		Sublocation:   params.Sublocation,
		Area:          params.Area,
		Version:       params.Version,
		Specification: params.Specification,
		URL:           ah.CreateResourceURL(i.Endpoint, params.AreaID),
	}
}
