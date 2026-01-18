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


func idToNamedAPIResource[T h.IsNamed, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], id int32) NamedAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seeding
	params := res.GetResParamsNamed()

	return NamedAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		URL:           cfg.createResourceURL(i.endpoint, params.ID),
	}
}

// essentially parses name or name/version like a handler and then converts the id to a NamedAPIResource
func nameToNamedAPIResource[T h.IsNamed, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], name string, version *int32) NamedAPIResource {
	var parseResp parseResponse
	switch version {
	case nil:
		parseResp, _ = checkUniqueName(name, i.objLookup)
	default:
		parseResp, _ = checkNameVersion(name, version, i.objLookup)
	}

	return idToNamedAPIResource(cfg, i, parseResp.ID)
}

func namesToNamedAPIResources[T h.IsNamed, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], names []string) []NamedAPIResource {
	resources := []NamedAPIResource{}

	for _, name := range names {
		resource := nameToNamedAPIResource(cfg, i, name, nil)
		resources = append(resources, resource)
	}

	return resources
}


// converts inputs to a resourceAmount of any kind by calling the given constructor func
func nameToResourceAmount[NA NameAmount, RA ResourceAmount, T h.IsNamed, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], item NA, fn func(NamedAPIResource, int32) RA) RA {
	resource := nameToNamedAPIResource(cfg, i, item.GetName(), item.GetVersion())
	return fn(resource, item.GetVal())
}


func namesToResourceAmounts[NA NameAmount, RA ResourceAmount, T h.IsNamed, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], items []NA, fn func(NamedAPIResource, int32) RA) []RA {
	results := []RA{}

	for _, item := range items {
		ra := nameToResourceAmount(cfg, i, item, fn)
		results = append(results, ra)
	}

	return results
}


func newNamedAPIResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []NamedAPIResource) (NamedApiResourceList, error) {
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


// only used for newNamedAPIResourceFromType, since the newer function won't work with that
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