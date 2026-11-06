package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToQuestAPIResource[T seedtypes.IsQuest, R any, A apiresource.APIResource](i handlerInput[T, R, A], id int32) apiresource.QuestAPIResource {
	questLookup, _ := lookup.GetResourceByID(id, i.objLookupID)
	params := questLookup.GetResParamsQuest()

	switch params.Type {
	case string(database.QuestTypeSidequest):
		sidequest, _ := lookup.GetResource(*params.Sidequest, lookup.L.Sidequests)

		return apiresource.QuestAPIResource{
			ID:        params.ID,
			Sidequest: sidequest.Name,
			Subquest:  nil,
			URL:       ah.CreateResourceURL(i.Endpoint, params.ID),
		}

	case string(database.QuestTypeSubquest):
		subquest, _ := lookup.GetResource(*params.Subquest, lookup.L.Subquests)
		sidequest, _ := lookup.GetResourceByID(subquest.SidequestID, lookup.L.SidequestsID)

		return apiresource.QuestAPIResource{
			ID:        params.ID,
			Sidequest: sidequest.Name,
			Subquest:  &subquest.Name,
			URL:       ah.CreateResourceURL(i.Endpoint, params.ID),
		}
	}

	return apiresource.QuestAPIResource{}
}

func questToQuestAPIResource(quest seedtypes.Quest) apiresource.QuestAPIResource {
	switch quest.Type {
	case database.QuestTypeSidequest:
		return nameToQuestAPIResource(E.Sidequests, quest.Name)

	case database.QuestTypeSubquest:
		return nameToQuestAPIResource(E.Subquests, quest.Name)
	}

	return apiresource.QuestAPIResource{}
}

func nameToQuestAPIResource[T seedtypes.IsQuest, R any, A apiresource.APIResource](i handlerInput[T, R, A], name string) apiresource.QuestAPIResource {
	quest, _ := lookup.GetResource(name, i.objLookup)
	return idToQuestAPIResource(i, quest.GetID())
}

func namesToQuestAPIResources[T seedtypes.IsQuest, R any, A apiresource.APIResource](i handlerInput[T, R, A], names []string) []apiresource.QuestAPIResource {
	resources := make([]apiresource.QuestAPIResource, len(names))

	for idx, name := range names {
		resources[idx] = nameToQuestAPIResource(i, name)
	}

	return resources
}
