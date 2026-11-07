package handlerinput

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func IDToTypedAPIResource[T seedtypes.IsTyped, R any, A apiresource.APIResource](i HandlerInput[T, R, A], id int32) apiresource.TypedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.ObjLookupID)
	params := res.GetResParamsTyped()

	return apiresource.TypedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		Type:          params.Type,
		URL:           ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}

func KeyToTypedAPIResource[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i HandlerInput[T, R, A], key K) apiresource.TypedAPIResource {
	res, _ := lookup.GetResource(key, i.ObjLookup)
	return IDToTypedAPIResource(i, res.GetID())
}

func KeyPtrToTypedAPIResourcePtr[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i HandlerInput[T, R, A], keyPtr *K) *apiresource.TypedAPIResource {
	if keyPtr == nil {
		return nil
	}

	res := KeyToTypedAPIResource(i, *keyPtr)
	return &res
}

func KeysToTypedAPIResources[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i HandlerInput[T, R, A], refs []K) []apiresource.TypedAPIResource {
	resources := make([]apiresource.TypedAPIResource, len(refs))

	for idx, ref := range refs {
		resources[idx] = KeyToTypedAPIResource(i, ref)
	}

	return resources
}
