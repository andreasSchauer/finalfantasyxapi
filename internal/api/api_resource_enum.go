package api

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EnumApiResourceList struct {
	ListParams
	Results []EnumAPIResource `json:"results"`
}

func (l EnumApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l EnumApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type EnumAPIResource struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}

func (r EnumAPIResource) IsZero() bool {
	return r.ID == 0
}

func (r EnumAPIResource) GetID() int32 {
	return r.ID
}

func (r EnumAPIResource) GetURL() string {
	return r.URL
}

func (r EnumAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r EnumAPIResource) Error() string {
	return fmt.Sprintf("enum api resource with url: %s", r.URL)
}

func (r EnumAPIResource) GetAPIResource() APIResource {
	return r
}

func newNamedAPIResourceFromEnum[E, N any](cfg *Config, endpoint, key string, et EnumType[E, N]) NamedAPIResource {
	enumRes, _ := GetEnumAPIResource(key, et)

	resource := newNamedAPIResourceSimple(cfg, endpoint, enumRes.ID, enumRes.Name)

	return resource
}

// Searches an EnumAPIResource based on its value or an idString (mostly from queries)
func GetEnumAPIResource[E, N any](key string, et EnumType[E, N]) (EnumAPIResource, error) {
	id, err := strconv.Atoi(key)
	if err == nil {
		if id > len(et.lookup) || id <= 0 {
			return EnumAPIResource{}, errIdNotFound
		}
		for _, res := range et.lookup {
			if int32(id) == res.ID {
				return res, nil
			}
		}
	}

	res, found := et.lookup[key]
	if found {
		return res, nil
	}

	keyWithSpaces := h.GetNameWithSpaces(key, "-")
	res, found = et.lookup[keyWithSpaces]
	if found {
		return res, nil
	}

	return EnumAPIResource{}, errNoResource
}

func newEnumAPIResourceList(cfg *Config, r *http.Request, resources []EnumAPIResource) (EnumApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return EnumApiResourceList{}, err
	}

	list := EnumApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}

func createTypeResourceSlice(cfg *Config, endpoint string, lookup map[string]EnumAPIResource) []EnumAPIResource {
	resources := []EnumAPIResource{}

	for _, resource := range lookup {
		resource.URL = createResourceURL(cfg, endpoint, resource.ID)
		resources = append(resources, resource)
	}

	slices.SortStableFunc(resources, sortAPIResources)

	return resources
}

func typeSliceToMap(enumTypes []EnumAPIResource) map[string]EnumAPIResource {
	typeMap := make(map[string]EnumAPIResource)

	for i, enumType := range enumTypes {
		typeMap[enumType.Name] = EnumAPIResource{
			ID:          int32(i + 1),
			Name:        enumType.Name,
			Description: enumType.Description,
		}
	}

	return typeMap
}
