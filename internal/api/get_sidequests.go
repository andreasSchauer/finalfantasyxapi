package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSidequest(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestApiResourceList], id int32) (Sidequest, error) {
	sidequest, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Sidequest{}, err
	}

	subquests, err := getResourcesDbItem(cfg, r, cfg.e.subquests, sidequest, cfg.db.GetSidequestSubquestIDs)
	if err != nil {
		return Sidequest{}, err
	}

	response := Sidequest{
		ID:         	sidequest.ID,
		Name:       	sidequest.Name,
		UntypedQuest: 	idToQuestAPIResource(cfg, cfg.e.quests, sidequest.Quest.ID),
		Availability:   newNamedAPIResourceFromEnum(cfg, cfg.e.availabilityType.endpoint, sidequest.Availability, cfg.t.AvailabilityType),
		Completion: 	convertObjPtr(cfg, sidequest.Completion, convertQuestCompletion),
		Subquests:  	subquests,
	}

	return response, nil
}

func (cfg *Config) retrieveSidequests(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestApiResourceList]) (QuestApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return QuestApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[QuestAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.AvailabilityType, resources, "availability", cfg.db.GetSidequestIDsByAvailability)),
	})
}
