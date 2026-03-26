package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type TypedAPIResourceList struct {
	ListParams
	Results []TypedAPIResource `json:"results"`
}

func (l TypedAPIResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l TypedAPIResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

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

func (r TypedAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r TypedAPIResource) GetKey() string {
	nameStr := h.NameToString(r.Name, r.Version, nil)
	return fmt.Sprintf("%s - %s", r.Type, nameStr)
}

func (r TypedAPIResource) Error() string {
	return fmt.Sprintf("typed api resource '%s', type: %s, url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.Type, r.URL)
}

func (r TypedAPIResource) GetAPIResource() APIResource {
	return r
}

func idToTypedAPIResource[T h.IsTyped, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], id int32) TypedAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID)
	params := res.GetResParamsTyped()

	return TypedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		Type:          params.Type,
		URL:           createResourceURL(cfg, i.endpoint, params.ID),
	}
}

func keyToTypedAPIResource[T h.IsTyped, R, K any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], key K) TypedAPIResource {
	res, _ := seeding.GetResource(key, i.objLookup)
	return idToTypedAPIResource(cfg, i, res.GetID())
}

func keyPtrToTypedAPIResourcePtr[T h.IsTyped, R, K any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], keyPtr *K) *TypedAPIResource {
	if keyPtr == nil {
		return nil
	}

	res := keyToTypedAPIResource(cfg, i, *keyPtr)
	return &res
}

func keysToTypedAPIResources[T h.IsTyped, R, K any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], refs []K) []TypedAPIResource {
	resources := []TypedAPIResource{}

	for _, ref := range refs {
		resource := keyToTypedAPIResource(cfg, i, ref)
		resources = append(resources, resource)
	}

	return resources
}



func newTypedAPIResourceList(cfg *Config, r *http.Request, resources []TypedAPIResource) (TypedAPIResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return TypedAPIResourceList{}, err
	}

	list := TypedAPIResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
