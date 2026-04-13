package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSubquest(r *http.Request, i handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList], id int32) (Subquest, error) {
	subquest, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Subquest{}, err
	}

	response := Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		UntypedQuest:    idToQuestAPIResource(cfg, cfg.e.quests, subquest.Quest.ID),
		ParentSidequest: idToQuestAPIResource(cfg, cfg.e.sidequests, subquest.SidequestID),
		Availability:    enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, subquest.Availability, cfg.t.AvailabilityType),
		IsRepeatable:    subquest.IsRepeatable,
		Completion:      convertQuestCompletion(cfg, *subquest.Completion),
	}

	if response.ParentSidequest.Sidequest == "monster arena" {
		response.ArenaCreation = namePtrToNamedAPIResPtr(cfg, cfg.e.arenaCreations, &subquest.Name, nil)
	}

	return response, nil
}

func (cfg *Config) retrieveSubquests(r *http.Request, i handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList]) (QuestApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return QuestApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[QuestAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.AvailabilityType, resources, "availability", cfg.db.GetSubquestIDsByAvailability)),
		frl(boolQuery(cfg, r, i, resources, "repeatable", cfg.db.GetSubquestIDsByRepeatable)),
	})
}
