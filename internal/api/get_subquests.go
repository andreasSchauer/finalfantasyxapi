package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSubquest(r *http.Request, i handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList], id int32) (Subquest, error) {
	subquest, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Subquest{}, err
	}

	response := Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		UntypedQuest:    idToQuestAPIResource(cfg, cfg.E.Quests, subquest.Quest.ID),
		ParentSidequest: idToQuestAPIResource(cfg, cfg.E.Sidequests, subquest.SidequestID),
		Availability:    subquest.Availability,
		IsRepeatable:    subquest.IsRepeatable,
		Completion:      convertQuestCompletion(cfg, *subquest.Completion),
	}

	if response.ParentSidequest.Sidequest == "monster arena" {
		response.ArenaCreation = namePtrToNamedAPIResPtr(cfg, cfg.E.ArenaCreations, &subquest.Name, nil)
	}

	return response, nil
}

func (cfg *Config) retrieveSubquests(r *http.Request, i handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		boolQuery(r, i, ids, QpnRepeatable, cfg.db.GetSubquestIDsByRepeatable),
	})
}
