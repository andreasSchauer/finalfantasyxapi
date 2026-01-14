package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type NamedApiResourceList struct {
	ListParams
	Results []NamedAPIResource `json:"results"`
}

func (l NamedApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l NamedApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

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

func (r NamedAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r NamedAPIResource) Error() string {
	return fmt.Sprintf("named api resource %s, version: %v, url: %s", r.Name, h.DerefOrNil(r.Version), r.URL)
}

func (r NamedAPIResource) GetAPIResource() APIResource {
	return r
}

// I can't tell yet, whether this function will be useful in the future.
// Good, if both id and name are already known
func (cfg *Config) newNamedAPIResource(endpoint string, id int32, name string, version *int32, spec *string) NamedAPIResource {
	if name == "" {
		return NamedAPIResource{}
	}

	return NamedAPIResource{
		ID:            id,
		Name:          name,
		Version:       version,
		Specification: spec,
		URL:           cfg.createResourceURL(endpoint, id),
	}
}

// I can't tell yet, whether this function will be useful in the future.
// Good, if both id and name are already known
func (cfg *Config) newNamedAPIResourceSimple(endpoint string, id int32, name string) NamedAPIResource {
	if name == "" {
		return NamedAPIResource{}
	}

	return NamedAPIResource{
		ID:   id,
		Name: name,
		URL:  cfg.createResourceURL(endpoint, id),
	}
}

// will be replaced by idsToNamedAPIResources
func createNamedAPIResources[T any](
	cfg *Config,
	items []T,
	endpoint string,
	mapper func(T) (id int32, name string, version *int32, spec *string),
) []NamedAPIResource {
	resources := []NamedAPIResource{}

	for _, item := range items {
		id, name, version, spec := mapper(item)
		resource := cfg.newNamedAPIResource(endpoint, id, name, version, spec)

		if !resource.IsZero() {
			resources = append(resources, resource)
		}
	}

	return resources
}

// will be replaced by idsToNamedAPIResources
func createNamedAPIResourcesSimple[T any](
	cfg *Config,
	items []T,
	endpoint string,
	mapper func(T) (id int32, name string),
) []NamedAPIResource {
	resources := []NamedAPIResource{}

	for _, item := range items {
		id, name := mapper(item)
		resource := cfg.newNamedAPIResourceSimple(endpoint, id, name)

		if !resource.IsZero() {
			resources = append(resources, resource)
		}
	}

	return resources
}

func newNamedAPIResourceList[T h.HasID, R any, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, L], resources []NamedAPIResource) (NamedApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, i, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	list := NamedApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}

// id and lookup based stuff

// good for db returns and if other funcs need to create a NamedAPIResource
func idToNamedAPIResource[T h.IsNamed](cfg *Config, endpoint string, id int32, lookup map[int32]T) NamedAPIResource {
	res, _ := seeding.GetResourceByID(id, lookup) // no error needed, because everything was verified through seeding

	params := res.GetResParamsNamed()

	return NamedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		URL:           cfg.createResourceURL(endpoint, params.ID),
	}
}

// essentially parses name or name/version like a handler and then converts the id to a NamedAPIResource
func nameToNamedAPIResource[T h.IsNamed](cfg *Config, endpoint, name string, version *int32, lookup map[string]T, idLookup map[int32]T) NamedAPIResource {
	if name == "" {
		return NamedAPIResource{}
	}

	var parseResp parseResponse
	switch version {
	case nil:
		parseResp, _ = checkUniqueName(name, lookup)
	default:
		parseResp, _ = checkNameVersion(name, version, lookup)
	}

	return idToNamedAPIResource(cfg, endpoint, parseResp.ID, idLookup)
}

// converts inputs to a resourceAmount of any kind by calling the given constructor func
// still need a method of type assertion for itemAmount that is done before calling this function
func nameToResourceAmount[R ResourceAmount, T h.IsNamed](cfg *Config, endpoint, name string, version *int32, amount int32, lookup map[string]T, lookupID map[int32]T, fn func(NamedAPIResource, int32) R) R {
	resource := nameToNamedAPIResource(cfg, endpoint, name, version, lookup, lookupID)
	return fn(resource, amount)
}

// takes a slice of ids (e.g. from a db list query) and creates a []NamedAPIResource
func idsToNamedAPIResources[T h.IsNamed](cfg *Config, endpoint string, IDs []int32, lookup map[int32]T) []NamedAPIResource {
	resources := []NamedAPIResource{}

	for _, id := range IDs {
		resource := idToNamedAPIResource(cfg, endpoint, id, lookup)
		resources = append(resources, resource)
	}

	return resources
}
