package main

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type TypedApiResourceList struct {
	ListParams
	Results []TypedAPIResource `json:"results"`
}

func (l TypedApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l TypedApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type TypedAPIResource struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}

func (r TypedAPIResource) IsZero() bool {
	return r.ID == 0
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

func (r TypedAPIResource) Error() string {
	return fmt.Sprintf("Typed api resource with url: %s", r.URL)
}

func (r TypedAPIResource) GetAPIResource() APIResource {
	return r
}

func (cfg *Config) newNamedAPIResourceFromType(endpoint, key string, lookup map[string]TypedAPIResource) (NamedAPIResource, error) {
	enumType, err := GetEnumType(key, lookup)
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get enum '%s' for '%s'.", key, endpoint), fmt.Errorf("%s: %v", endpoint, err))
	}

	resource := cfg.newNamedAPIResourceSimple(endpoint, enumType.ID, enumType.Name)

	return resource, nil
}

func newTypedAPIResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], lookup map[string]TypedAPIResource) (TypedApiResourceList, error) {
	resources := cfg.createTypeResourceSlice(i.endpoint, lookup)

	listParams, shownResources, err := createPaginatedList(cfg, r, i, resources)
	if err != nil {
		return TypedApiResourceList{}, err
	}

	list := TypedApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}

func (cfg *Config) createTypeResourceSlice(endpoint string, lookup map[string]TypedAPIResource) []TypedAPIResource {
	resources := []TypedAPIResource{}

	for _, resource := range lookup {
		resource.URL = cfg.createResourceURL(endpoint, resource.ID)
		resources = append(resources, resource)
	}

	slices.SortStableFunc(resources, sortAPIResources)

	return resources
}

func typeSliceToMap(enumTypes []TypedAPIResource) map[string]TypedAPIResource {
	typeMap := make(map[string]TypedAPIResource)

	for i, enumType := range enumTypes {
		typeMap[enumType.Name] = TypedAPIResource{
			ID:          int32(i + 1),
			Name:        enumType.Name,
			Description: enumType.Description,
		}
	}

	return typeMap
}
