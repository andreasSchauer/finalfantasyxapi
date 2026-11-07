package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToTypedAPIResource[T seedtypes.IsTyped, R any, A apiresource.APIResource](i HandlerInput[T, R, A], id int32) apiresource.TypedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID)
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

func keyToTypedAPIResource[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i HandlerInput[T, R, A], key K) apiresource.TypedAPIResource {
	res, _ := lookup.GetResource(key, i.objLookup)
	return idToTypedAPIResource(i, res.GetID())
}

func keyPtrToTypedAPIResourcePtr[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i HandlerInput[T, R, A], keyPtr *K) *apiresource.TypedAPIResource {
	if keyPtr == nil {
		return nil
	}

	res := keyToTypedAPIResource(i, *keyPtr)
	return &res
}

func keysToTypedAPIResources[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i HandlerInput[T, R, A], refs []K) []apiresource.TypedAPIResource {
	resources := make([]apiresource.TypedAPIResource, len(refs))

	for idx, ref := range refs {
		resources[idx] = keyToTypedAPIResource(i, ref)
	}

	return resources
}
