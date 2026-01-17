package main

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

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

func newNamedAPIResourceFromType[T, N any](cfg *Config, endpoint, key string, et EnumType[T, N]) (NamedAPIResource, error) {
	enumType, err := GetTypedAPIResource(key, et)
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, err.Error(), fmt.Errorf("%s: %v", endpoint, err))
	}

	resource := cfg.newNamedAPIResourceSimple(endpoint, enumType.ID, enumType.Name)

	return resource, nil
}


// Searches a TypedAPIResource based on its value or an idString (mostly from queries)
func GetTypedAPIResource[T, N any](key string, et EnumType[T, N]) (TypedAPIResource, error) {
	id, err := strconv.Atoi(key)
	if err == nil {
		for _, res := range et.lookup {
			if int32(id) == res.ID {
				return res, nil
			}
		}
	}

	res, found := et.lookup[key]
	if !found {
		return TypedAPIResource{}, fmt.Errorf("value '%s' is not valid in enum '%s'.", key, et.name)
	}

	return res, nil
}

func newTypedAPIResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []TypedAPIResource) (TypedApiResourceList, error) {
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
