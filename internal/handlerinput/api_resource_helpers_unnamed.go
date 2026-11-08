package handlerinput

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func IDToUnnamedAPIResource[T seedtypes.IsUnnamed, R any, A apiresource.APIResource](i Main[T, R, A], id int32) apiresource.Unnamed {
	res, _ := lookup.GetResourceByID(id, i.ObjLookupID) // no error needed, because everything was verified through seedtypes
	params := res.GetResParamsUnnamed()

	return apiresource.Unnamed{
		ID:  params.ID,
		URL: ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}
