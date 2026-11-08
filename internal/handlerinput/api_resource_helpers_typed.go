package handlerinput

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func IDToTypedAPIResource[T seedtypes.IsTyped, R any, A apiresource.APIResource](i Main[T, R, A], id int32) apiresource.Typed {
	res, _ := lookup.GetResourceByID(id, i.ObjLookupID)
	params := res.GetResParamsTyped()

	return apiresource.Typed{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		Type:          params.Type,
		URL:           ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}

func KeyToTypedAPIResource[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i Main[T, R, A], key K) apiresource.Typed {
	res, _ := lookup.GetResource(key, i.ObjLookup)
	return IDToTypedAPIResource(i, res.GetID())
}

func KeyPtrToTypedAPIResourcePtr[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i Main[T, R, A], keyPtr *K) *apiresource.Typed {
	if keyPtr == nil {
		return nil
	}

	res := KeyToTypedAPIResource(i, *keyPtr)
	return &res
}

func KeysToTypedAPIResources[T seedtypes.IsTyped, R, K any, A apiresource.APIResource](i Main[T, R, A], refs []K) []apiresource.Typed {
	resources := make([]apiresource.Typed, len(refs))

	for idx, ref := range refs {
		resources[idx] = KeyToTypedAPIResource(i, ref)
	}

	return resources
}
