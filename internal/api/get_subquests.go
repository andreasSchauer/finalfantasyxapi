package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getSubquest(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource], id int32) (Subquest, error) {
	subquest, err := i.queryParamsVerifyGet(r, id)
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

func (cfg *Config) retrieveSubquests(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryBool(r, ids, typestrings.QpnRepeatable, cfg.db.GetSubquestIDsByRepeatable),
	})
}
