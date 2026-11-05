package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToUnnamedAPIResource[T seedtypes.IsUnnamed, R any, A APIResource](i handlerInput[T, R, A], id int32) UnnamedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seedtypes
	params := res.GetResParamsUnnamed()

	return UnnamedAPIResource{
		ID:  params.ID,
		URL: ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}
