package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToUnnamedAPIResource[T seedtypes.IsUnnamed, R any, A apiresource.APIResource](i HandlerInput[T, R, A], id int32) apiresource.UnnamedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seedtypes
	params := res.GetResParamsUnnamed()

	return apiresource.UnnamedAPIResource{
		ID:  params.ID,
		URL: ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}
