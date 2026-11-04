package api

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type NamedAPIResource struct {
	ID            int32   `json:"-"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r NamedAPIResource) IsZero() bool {
	return r.Name == ""
}

func (r NamedAPIResource) GetID() int32 {
	return r.ID
}

func (r NamedAPIResource) GetURL() string {
	return r.URL
}

func (r NamedAPIResource) GetTestKey() string {
	return h.NameToString(r.Name, r.Version, nil)
}

func (r NamedAPIResource) String() string {
	return fmt.Sprintf("named api resource '%s', url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.URL)
}

func (r NamedAPIResource) GetAPIResource() APIResource {
	return r
}

func (r NamedAPIResource) NameIs(name string, version *int32) bool {
	return r.Name == name && h.PtrsEqual(r.Version, version, h.DerefInt32Ptr)
}

func idToNamedAPIResource[T seedtypes.IsNamed, R any, A APIResource](server *Server, i handlerInput[T, R, A], id int32) NamedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seedtypes
	params := res.GetResParamsNamed()

	return NamedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		URL:           CreateResourceURL(server, i.Endpoint, params.ID),
	}
}

// essentially parses name or name/version like a handler and then converts the id to a NamedAPIResource
func nameToNamedAPIResource[T seedtypes.IsNamed, R any, A APIResource](server *Server, i handlerInput[T, R, A], name string, version *int32) NamedAPIResource {
	var parseResp parseResponse
	switch version {
	case nil:
		parseResp, _ = checkUniqueName(name, i.objLookup)
	default:
		parseResp, _ = checkNameVersion(name, version, i.objLookup)
	}

	return idToNamedAPIResource(server, i, parseResp.ID)
}

func namePtrToNamedAPIResPtr[T seedtypes.IsNamed, R any, A APIResource](server *Server, i handlerInput[T, R, A], namePtr *string, version *int32) *NamedAPIResource {
	if namePtr == nil {
		return nil
	}

	res := nameToNamedAPIResource(server, i, *namePtr, version)
	return &res
}

func namesToNamedAPIResources[T seedtypes.IsNamed, R any, A APIResource](server *Server, i handlerInput[T, R, A], names []string) []NamedAPIResource {
	resources := make([]NamedAPIResource, 0, len(names)) // put this into more high frequency places

	for _, name := range names {
		resource := nameToNamedAPIResource(server, i, name, nil)
		resources = append(resources, resource)
	}

	return resources
}

func enumToNamedAPIResource(server *Server, e enums.EnumResponse) NamedAPIResource {
	return NamedAPIResource{
		Name: string(e.Name),
		URL:  createEnumURL(server, e.Name),
	}
}

func enumsToNamedAPIResources(server *Server, enums []enums.EnumResponse) []NamedAPIResource {
	resources := make([]NamedAPIResource, 0, len(enums))

	for _, enum := range enums {
		resource := enumToNamedAPIResource(server, enum)
		resources = append(resources, resource)
	}

	return resources
}

func endpointToNamedAPIResource(server *Server, e ts.EndpointName) NamedAPIResource {
	return NamedAPIResource{
		Name: string(e),
		URL:  createListURL(server, e),
	}
}

func endpointsToNamedAPIResources(server *Server, endpoints []ts.EndpointName) []NamedAPIResource {
	resources := make([]NamedAPIResource, 0, len(endpoints))

	for _, ep := range endpoints {
		resource := endpointToNamedAPIResource(server, ep)
		resources = append(resources, resource)
	}

	return resources
}

func refsToNamedAPIResources(server *Server, refs []seedtypes.AbilityReference) []NamedAPIResource {
	abilities := []NamedAPIResource{}

	for _, ref := range refs {
		abilityRes := refToNamedAPIResource(server, ref)
		abilities = append(abilities, abilityRes)
	}

	return abilities
}

func refToNamedAPIResource(server *Server, ref seedtypes.AbilityReference) NamedAPIResource {
	var res NamedAPIResource

	switch database.AbilityType(ref.AbilityType) {
	case database.AbilityTypeMiscAbility:
		res = nameToNamedAPIResource(server, E.MiscAbilities, ref.Name, ref.Version)

	case database.AbilityTypePlayerAbility:
		res = nameToNamedAPIResource(server, E.PlayerAbilities, ref.Name, ref.Version)

	case database.AbilityTypeEnemyAbility:
		res = nameToNamedAPIResource(server, E.EnemyAbilities, ref.Name, ref.Version)

	case database.AbilityTypeOverdriveAbility:
		res = nameToNamedAPIResource(server, E.OverdriveAbilities, ref.Name, ref.Version)

	case database.AbilityTypeItemAbility:
		res = nameToNamedAPIResource(server, E.ItemAbilities, ref.Name, ref.Version)

	case database.AbilityTypeTriggerCommand:
		res = nameToNamedAPIResource(server, E.TriggerCommands, ref.Name, ref.Version)

	}
	return res
}
