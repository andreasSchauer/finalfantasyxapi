package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getSubquest(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, apiresource.QuestAPIResource], id int32) (Subquest, error) {
	subquest, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Subquest{}, err
	}

	response := Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		UntypedQuest:    idToQuestAPIResource(E.Quests, subquest.Quest.ID),
		ParentSidequest: idToQuestAPIResource(E.Sidequests, subquest.SidequestID),
		Availability:    subquest.Availability,
		IsRepeatable:    subquest.IsRepeatable,
		Completion:      convertQuestCompletion(*subquest.Completion),
	}

	if response.ParentSidequest.Sidequest == "monster arena" {
		response.ArenaCreation = namePtrToNamedAPIResPtr(E.ArenaCreations, &subquest.Name, nil)
	}

	return response, nil
}

func retrieveSubquests(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, apiresource.QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnRepeatable, config.Cfg.DB.GetSubquestIDsByRepeatable),
	})
}
