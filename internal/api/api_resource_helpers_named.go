package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func idToNamedAPIResource[T seedtypes.IsNamed, R any, A apiresource.APIResource](i handlerInput[T, R, A], id int32) apiresource.NamedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seedtypes
	params := res.GetResParamsNamed()

	return apiresource.NamedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		URL:           ah.CreateResourceURL(i.Endpoint, params.ID),
	}
}

// essentially parses name or name/version like a handler and then converts the id to a apiresource.NamedAPIResource
func nameToNamedAPIResource[T seedtypes.IsNamed, R any, A apiresource.APIResource](i handlerInput[T, R, A], name string, version *int32) apiresource.NamedAPIResource {
	var parseResp ah.ParseResponse
	switch version {
	case nil:
		parseResp, _ = ah.CheckUniqueName(name, i.objLookup)
	default:
		parseResp, _ = ah.CheckNameVersion(name, version, i.objLookup)
	}

	return idToNamedAPIResource(i, parseResp.ID)
}

func namePtrToNamedAPIResPtr[T seedtypes.IsNamed, R any, A apiresource.APIResource](i handlerInput[T, R, A], namePtr *string, version *int32) *apiresource.NamedAPIResource {
	if namePtr == nil {
		return nil
	}

	res := nameToNamedAPIResource(i, *namePtr, version)
	return &res
}

func namesToNamedAPIResources[T seedtypes.IsNamed, R any, A apiresource.APIResource](i handlerInput[T, R, A], names []string) []apiresource.NamedAPIResource {
	resources := make([]apiresource.NamedAPIResource, len(names))

	for idx, name := range names {
		resources[idx] = nameToNamedAPIResource(i, name, nil)
	}

	return resources
}

func enumToNamedAPIResource(e enums.EnumResponse) apiresource.NamedAPIResource {
	return apiresource.NamedAPIResource{
		Name: string(e.Name),
		URL:  ah.CreateEnumURL(e.Name),
	}
}

func enumsToNamedAPIResources(enums []enums.EnumResponse) []apiresource.NamedAPIResource {
	resources := make([]apiresource.NamedAPIResource, 0, len(enums))

	for _, enum := range enums {
		resource := enumToNamedAPIResource(enum)
		resources = append(resources, resource)
	}

	return resources
}

func endpointToNamedAPIResource(e ts.EndpointName) apiresource.NamedAPIResource {
	return apiresource.NamedAPIResource{
		Name: string(e),
		URL:  ah.CreateListURL(e),
	}
}

func endpointsToNamedAPIResources(endpoints []ts.EndpointName) []apiresource.NamedAPIResource {
	resources := make([]apiresource.NamedAPIResource, 0, len(endpoints))

	for _, ep := range endpoints {
		resource := endpointToNamedAPIResource(ep)
		resources = append(resources, resource)
	}

	return resources
}

func refsToNamedAPIResources(refs []seedtypes.AbilityReference) []apiresource.NamedAPIResource {
	abilities := []apiresource.NamedAPIResource{}

	for _, ref := range refs {
		abilityRes := refToNamedAPIResource(ref)
		abilities = append(abilities, abilityRes)
	}

	return abilities
}

func refToNamedAPIResource(ref seedtypes.AbilityReference) apiresource.NamedAPIResource {
	var res apiresource.NamedAPIResource

	switch database.AbilityType(ref.AbilityType) {
	case database.AbilityTypeMiscAbility:
		res = nameToNamedAPIResource(E.MiscAbilities, ref.Name, ref.Version)

	case database.AbilityTypePlayerAbility:
		res = nameToNamedAPIResource(E.PlayerAbilities, ref.Name, ref.Version)

	case database.AbilityTypeEnemyAbility:
		res = nameToNamedAPIResource(E.EnemyAbilities, ref.Name, ref.Version)

	case database.AbilityTypeOverdriveAbility:
		res = nameToNamedAPIResource(E.OverdriveAbilities, ref.Name, ref.Version)

	case database.AbilityTypeItemAbility:
		res = nameToNamedAPIResource(E.ItemAbilities, ref.Name, ref.Version)

	case database.AbilityTypeTriggerCommand:
		res = nameToNamedAPIResource(E.TriggerCommands, ref.Name, ref.Version)

	}
	return res
}
