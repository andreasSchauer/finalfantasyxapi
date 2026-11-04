package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getSidequest(r *http.Request, i handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource], id int32) (Sidequest, error) {
	sidequest, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sidequest{}, err
	}

	subquests, err := E.Subquests.dbGetRelations(cfg, r.Context(), sidequest, cfg.db.GetSidequestSubquestIDs)
	if err != nil {
		return Sidequest{}, err
	}

	response := Sidequest{
		ID:           sidequest.ID,
		Name:         sidequest.Name,
		UntypedQuest: idToQuestAPIResource(cfg, E.Quests, sidequest.Quest.ID),
		Availability: sidequest.Availability,
		IsRepeatable: sidequest.IsRepeatable,
		Completion:   convertObjPtr(cfg, sidequest.Completion, convertQuestCompletion),
		Subquests:    subquests,
	}

	return response, nil
}

func (cfg *Config) retrieveSidequests(r *http.Request, i handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{})
}
