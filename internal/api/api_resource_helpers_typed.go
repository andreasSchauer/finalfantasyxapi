package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToTypedAPIResource[T seedtypes.IsTyped, R any, A APIResource](i handlerInput[T, R, A], id int32) TypedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID)
	params := res.GetResParamsTyped()

	return TypedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		Type:          params.Type,
		URL:           ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}

func keyToTypedAPIResource[T seedtypes.IsTyped, R, K any, A APIResource](i handlerInput[T, R, A], key K) TypedAPIResource {
	res, _ := lookup.GetResource(key, i.objLookup)
	return idToTypedAPIResource(i, res.GetID())
}

func keyPtrToTypedAPIResourcePtr[T seedtypes.IsTyped, R, K any, A APIResource](i handlerInput[T, R, A], keyPtr *K) *TypedAPIResource {
	if keyPtr == nil {
		return nil
	}

	res := keyToTypedAPIResource(i, *keyPtr)
	return &res
}

func keysToTypedAPIResources[T seedtypes.IsTyped, R, K any, A APIResource](i handlerInput[T, R, A], refs []K) []TypedAPIResource {
	resources := []TypedAPIResource{}

	for _, ref := range refs {
		resource := keyToTypedAPIResource(i, ref)
		resources = append(resources, resource)
	}

	return resources
}
