package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AbilityAPIResourceList struct {
	ListParams
	Results []AbilityAPIResource `json:"results"`
}

func (l AbilityAPIResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l AbilityAPIResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type AbilityAPIResource struct {
	ID            int32                `json:"-"`
	Name          string               `json:"name"`
	Version       *int32               `json:"version,omitempty"`
	Specification *string              `json:"specification,omitempty"`
	AbilityType   database.AbilityType `json:"ability_type"`
	URL           string               `json:"url"`
}

func (r AbilityAPIResource) IsZero() bool {
	return r.Name == ""
}

func (r AbilityAPIResource) GetID() int32 {
	return r.ID
}

func (r AbilityAPIResource) GetURL() string {
	return r.URL
}

func (r AbilityAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r AbilityAPIResource) Error() string {
	return fmt.Sprintf("ability api resource '%s', type: %s, url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.AbilityType, r.URL)
}

func (r AbilityAPIResource) GetAPIResource() APIResource {
	return r
}

func idToAbilityAPIResource(cfg *Config, i handlerInput[seeding.Ability, Ability, AbilityAPIResource, AbilityAPIResourceList], id int32) AbilityAPIResource {
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)
	return abilityToAbilityAPIResource(cfg, i, ability)
}

func refToAbilityAPIResource(cfg *Config, i handlerInput[seeding.Ability, Ability, AbilityAPIResource, AbilityAPIResourceList], ref seeding.AbilityReference) AbilityAPIResource {
	ability, _ := seeding.GetResource(ref, i.objLookup)
	return abilityToAbilityAPIResource(cfg, i, ability)
}

func abilityToAbilityAPIResource(cfg *Config, i handlerInput[seeding.Ability, Ability, AbilityAPIResource, AbilityAPIResourceList], ability seeding.Ability) AbilityAPIResource {
	params := ability.GetResParamsAbility()

	return AbilityAPIResource{
		ID:            params.ID,
		Name:          params.Name,
		Version:       params.Version,
		Specification: params.Specification,
		AbilityType:   params.AbilityType,
		URL:           createResourceURL(cfg, i.endpoint, params.ID),
	}
}

func refPtrToAbilityAPIResourcePtr(cfg *Config, i handlerInput[seeding.Ability, Ability, AbilityAPIResource, AbilityAPIResourceList], refPtr *seeding.AbilityReference) *AbilityAPIResource {
	if refPtr == nil {
		return nil
	}

	res := refToAbilityAPIResource(cfg, i, *refPtr)
	return &res
}

func refsToAbilityAPIResources(cfg *Config, i handlerInput[seeding.Ability, Ability, AbilityAPIResource, AbilityAPIResourceList], refs []seeding.AbilityReference) []AbilityAPIResource {
	resources := []AbilityAPIResource{}

	for _, ref := range refs {
		resource := refToAbilityAPIResource(cfg, i, ref)
		resources = append(resources, resource)
	}

	return resources
}

func newAbilityAPIResourceList(cfg *Config, r *http.Request, resources []AbilityAPIResource) (AbilityAPIResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return AbilityAPIResourceList{}, err
	}

	list := AbilityAPIResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
