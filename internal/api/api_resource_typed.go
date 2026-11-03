package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type TypedAPIResource struct {
	ID            int32   `json:"-"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
	Type          string  `json:"type"`
	URL           string  `json:"url"`
}

func (r TypedAPIResource) IsZero() bool {
	return r.Name == ""
}

func (r TypedAPIResource) GetID() int32 {
	return r.ID
}

func (r TypedAPIResource) GetURL() string {
	return r.URL
}

func (r TypedAPIResource) GetTestKey() string {
	nameStr := h.NameToString(r.Name, r.Version, nil)
	return fmt.Sprintf("%s - %s", r.Type, nameStr)
}

func (r TypedAPIResource) String() string {
	return fmt.Sprintf("typed api resource '%s', type: %s, url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.Type, r.URL)
}

func (r TypedAPIResource) GetAPIResource() APIResource {
	return r
}

func idToTypedAPIResource[T seedtypes.IsTyped, R any, A APIResource](cfg *Config, i handlerInput[T, R, A], id int32) TypedAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID)
	params := res.GetResParamsTyped()

	return TypedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		Type:          params.Type,
		URL:           CreateResourceURL(cfg, i.Endpoint, params.ID),
	}
}

func keyToTypedAPIResource[T seedtypes.IsTyped, R, K any, A APIResource](cfg *Config, i handlerInput[T, R, A], key K) TypedAPIResource {
	res, _ := seeding.GetResource(key, i.objLookup)
	return idToTypedAPIResource(cfg, i, res.GetID())
}

func keyPtrToTypedAPIResourcePtr[T seedtypes.IsTyped, R, K any, A APIResource](cfg *Config, i handlerInput[T, R, A], keyPtr *K) *TypedAPIResource {
	if keyPtr == nil {
		return nil
	}

	res := keyToTypedAPIResource(cfg, i, *keyPtr)
	return &res
}

func keysToTypedAPIResources[T seedtypes.IsTyped, R, K any, A APIResource](cfg *Config, i handlerInput[T, R, A], refs []K) []TypedAPIResource {
	resources := []TypedAPIResource{}

	for _, ref := range refs {
		resource := keyToTypedAPIResource(cfg, i, ref)
		resources = append(resources, resource)
	}

	return resources
}
