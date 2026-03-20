package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type QuestApiResourceList struct {
	ListParams
	Results		[]QuestAPIResource	`json:"results"`
}

func (l QuestApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l QuestApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type QuestAPIResource struct {
	ID 			int32	`json:"-"`
	Sidequest	string	`json:"sidequest"`
	Subquest	string	`json:"subquest"`
	URL			string	`json:"url"`
}

func (r QuestAPIResource) IsZero() bool {
	return r.Subquest == ""
}

func (r QuestAPIResource) GetID() int32 {
	return r.ID
}

func (r QuestAPIResource) GetURL() string {
	return r.URL
}

func (r QuestAPIResource) ToKeyFields() []any {
	return []any{
		r.URL,
	}
}

func (r QuestAPIResource) Error() string {
	return fmt.Sprintf("quest api resource '%s', url: %s", r.Subquest, r.URL)
}

func (r QuestAPIResource) GetAPIResource() APIResource {
	return r
}

func idToQuestAPIResource(cfg *Config, i handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList], id int32) QuestAPIResource {
	subquest, _ := seeding.GetResourceByID(id, i.objLookupID)
	sidequest, _ := seeding.GetResourceByID(subquest.SidequestID, cfg.l.SidequestsID)

	return QuestAPIResource{
		ID: 		subquest.ID,
		Sidequest: 	sidequest.Name,
		Subquest:	subquest.Name,
		URL: 		createResourceURL(cfg, i.endpoint, subquest.ID),
	}
}

func nameToQuestAPIResource(cfg *Config, i handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList], name string) QuestAPIResource {
	subquest,_ := seeding.GetResource(name, i.objLookup)
	return idToQuestAPIResource(cfg, i, subquest.ID)
}

func namesToQuestAPIResources(cfg *Config, i handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList], names []string) []QuestAPIResource {
	resources := []QuestAPIResource{}

	for _, name := range names {
		resource := nameToQuestAPIResource(cfg, i, name)
		resources = append(resources, resource)
	}

	return resources
}

func newQuestAPIResourceList(cfg *Config, r *http.Request, resources []QuestAPIResource) (QuestApiResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return QuestApiResourceList{}, err
	}

	list := QuestApiResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}