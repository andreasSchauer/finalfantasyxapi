package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToQuestAPIResource[T seedtypes.IsQuest, R any, A APIResource](i handlerInput[T, R, A], id int32) QuestAPIResource {
	questLookup, _ := lookup.GetResourceByID(id, i.objLookupID)
	params := questLookup.GetResParamsQuest()

	switch params.Type {
	case string(database.QuestTypeSidequest):
		sidequest, _ := lookup.GetResource(*params.Sidequest, lookup.L.Sidequests)

		return QuestAPIResource{
			ID:        params.ID,
			Sidequest: sidequest.Name,
			Subquest:  nil,
			URL:       ah.CreateResourceURL(i.Endpoint, params.ID),
		}

	case string(database.QuestTypeSubquest):
		subquest, _ := lookup.GetResource(*params.Subquest, lookup.L.Subquests)
		sidequest, _ := lookup.GetResourceByID(subquest.SidequestID, lookup.L.SidequestsID)

		return QuestAPIResource{
			ID:        params.ID,
			Sidequest: sidequest.Name,
			Subquest:  &subquest.Name,
			URL:       ah.CreateResourceURL(i.Endpoint, params.ID),
		}
	}

	return QuestAPIResource{}
}

func questToQuestAPIResource(quest seedtypes.Quest) QuestAPIResource {
	switch quest.Type {
	case database.QuestTypeSidequest:
		return nameToQuestAPIResource(E.Sidequests, quest.Name)

	case database.QuestTypeSubquest:
		return nameToQuestAPIResource(E.Subquests, quest.Name)
	}

	return QuestAPIResource{}
}

func nameToQuestAPIResource[T seedtypes.IsQuest, R any, A APIResource](i handlerInput[T, R, A], name string) QuestAPIResource {
	quest, _ := lookup.GetResource(name, i.objLookup)
	return idToQuestAPIResource(i, quest.GetID())
}

func namesToQuestAPIResources[T seedtypes.IsQuest, R any, A APIResource](i handlerInput[T, R, A], names []string) []QuestAPIResource {
	resources := []QuestAPIResource{}

	for _, name := range names {
		resource := nameToQuestAPIResource(i, name)
		resources = append(resources, resource)
	}

	return resources
}
