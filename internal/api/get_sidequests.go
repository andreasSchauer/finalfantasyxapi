package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSidequest(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestAPIResourceList], id int32) (Sidequest, error) {
	sidequest, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Sidequest{}, err
	}

	subquests, err := getResourcesDBItem(cfg, r.Context(), cfg.E.Subquests, sidequest, cfg.db.GetSidequestSubquestIDs)
	if err != nil {
		return Sidequest{}, err
	}

	response := Sidequest{
		ID:           sidequest.ID,
		Name:         sidequest.Name,
		UntypedQuest: idToQuestAPIResource(cfg, cfg.E.Quests, sidequest.Quest.ID),
		Availability: sidequest.Availability,
		IsRepeatable: sidequest.IsRepeatable,
		Completion:   convertObjPtr(cfg, sidequest.Completion, convertQuestCompletion),
		Subquests:    subquests,
	}

	return response, nil
}

func (cfg *Config) retrieveSidequests(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{})
}
