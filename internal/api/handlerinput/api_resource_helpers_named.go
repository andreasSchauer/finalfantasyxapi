package handlerinput

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func IDToNamedAPIResource[T seedtypes.IsNamed, R any, A apiresource.APIResource](i Main[T, R, A], id int32) apiresource.Named {
	res, _ := lookup.GetResourceByID(id, i.ObjLookupID) // no error needed, because everything was verified through seedtypes
	params := res.GetResParamsNamed()

	return apiresource.Named{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		URL:           ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}

func NameToNamedAPIResource[T seedtypes.IsNamed, R any, A apiresource.APIResource](i Main[T, R, A], name string, version *int32) apiresource.Named {
	var parseResp ah.ParseResponse
	switch version {
	case nil:
		parseResp, _ = ah.CheckUniqueName(name, i.ObjLookup)
	default:
		parseResp, _ = ah.CheckNameVersion(name, version, i.ObjLookup)
	}

	return IDToNamedAPIResource(i, parseResp.ID)
}

func NamePtrToNamedAPIResPtr[T seedtypes.IsNamed, R any, A apiresource.APIResource](i Main[T, R, A], namePtr *string, version *int32) *apiresource.Named {
	if namePtr == nil {
		return nil
	}

	res := NameToNamedAPIResource(i, *namePtr, version)
	return &res
}

func NamesToNamedAPIResources[T seedtypes.IsNamed, R any, A apiresource.APIResource](i Main[T, R, A], names []string) []apiresource.Named {
	resources := make([]apiresource.Named, len(names))

	for idx, name := range names {
		resources[idx] = NameToNamedAPIResource(i, name, nil)
	}

	return resources
}

func enumToNamedAPIResource(e enums.EnumResponse) apiresource.Named {
	return apiresource.Named{
		Name: string(e.Name),
		URL:  ah.CreateEnumURL(e.Name),
	}
}

func enumsToNamedAPIResources(enums []enums.EnumResponse) []apiresource.Named {
	resources := make([]apiresource.Named, 0, len(enums))

	for _, enum := range enums {
		resource := enumToNamedAPIResource(enum)
		resources = append(resources, resource)
	}

	return resources
}

func endpointToNamedAPIResource(e ts.EndpointName) apiresource.Named {
	return apiresource.Named{
		Name: string(e),
		URL:  ah.CreateListURL(e),
	}
}

func EndpointsToNamedAPIResources(endpoints []ts.EndpointName) []apiresource.Named {
	resources := make([]apiresource.Named, 0, len(endpoints))

	for _, ep := range endpoints {
		resource := endpointToNamedAPIResource(ep)
		resources = append(resources, resource)
	}

	return resources
}
