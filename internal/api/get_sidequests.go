package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getSidequest(r *http.Request, i handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource], id int32) (Sidequest, error) {
	sidequest, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sidequest{}, err
	}

	subquests, err := E.Subquests.dbGetRelations(r.Context(), sidequest, config.Cfg.DB.GetSidequestSubquestIDs)
	if err != nil {
		return Sidequest{}, err
	}

	response := Sidequest{
		ID:           sidequest.ID,
		Name:         sidequest.Name,
		UntypedQuest: idToQuestAPIResource(E.Quests, sidequest.Quest.ID),
		Availability: sidequest.Availability,
		IsRepeatable: sidequest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(sidequest.Completion, convertQuestCompletion),
		Subquests:    subquests,
	}

	return response, nil
}

func retrieveSidequests(r *http.Request, i handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{})
}
