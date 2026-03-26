package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type QuestApiResourceList struct {
	ListParams
	Results []QuestAPIResource `json:"results"`
}

func (l QuestApiResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l QuestApiResourceList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type QuestAPIResource struct {
	ID        int32   `json:"-"`
	Sidequest string  `json:"sidequest"`
	Subquest  *string `json:"subquest,omitempty"`
	URL       string  `json:"url"`
}

func (r QuestAPIResource) IsZero() bool {
	return r.Sidequest == "" && r.Subquest == nil
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

func (r QuestAPIResource) GetKey() string {
	var subquestStr string

	if r.Subquest != nil {
		subquestStr = fmt.Sprintf(" - %s", *r.Subquest)
	}
	return r.Sidequest + subquestStr
}

func (r QuestAPIResource) Error() string {
	return fmt.Sprintf("quest api resource '%s', url: %s", h.PtrToString(r.Subquest), r.URL)
}

func (r QuestAPIResource) GetAPIResource() APIResource {
	return r
}

func idToQuestAPIResource[T h.IsQuest, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], id int32) QuestAPIResource {
	questLookup, _ := seeding.GetResourceByID(id, i.objLookupID)
	params := questLookup.GetResParamsQuest()

	switch params.Type {
	case string(database.QuestTypeSidequest):
		sidequest, _ := seeding.GetResource(*params.Sidequest, cfg.l.Sidequests)

		return QuestAPIResource{
			ID: 		params.ID,
			Sidequest: 	sidequest.Name,
			Subquest: 	nil,
			URL: 		createResourceURL(cfg, i.endpoint, params.ID),
		}

	case string(database.QuestTypeSubquest):
		subquest, _ := seeding.GetResource(*params.Subquest, cfg.l.Subquests)
		sidequest, _ := seeding.GetResourceByID(subquest.SidequestID, cfg.l.SidequestsID)

		return QuestAPIResource{
			ID: 		params.ID,
			Sidequest: 	sidequest.Name,
			Subquest: 	&subquest.Name,
			URL: 		createResourceURL(cfg,i.endpoint, params.ID),
		}
	}

	return QuestAPIResource{}
}

func questToQuestAPIResource(cfg *Config, quest seeding.Quest) QuestAPIResource {
	switch quest.Type {
	case database.QuestTypeSidequest:
		return nameToQuestAPIResource(cfg, cfg.e.sidequests, quest.Name)

	case database.QuestTypeSubquest:
		return nameToQuestAPIResource(cfg, cfg.e.subquests, quest.Name)
	}

	return QuestAPIResource{}
}

func nameToQuestAPIResource[T h.IsQuest, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], name string) QuestAPIResource {
	quest, _ := seeding.GetResource(name, i.objLookup)
	return idToQuestAPIResource(cfg, i, quest.GetID())
}

func namesToQuestAPIResources[T h.IsQuest, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], names []string) []QuestAPIResource {
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
