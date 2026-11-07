package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getSidequest(r *http.Request, i handlerinput.HandlerInput[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource], id int32) (apitypes.Sidequest, error) {
	sidequest, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Sidequest{}, err
	}

	subquests, err := E.Subquests.DBGetRelations(r.Context(), sidequest, config.Cfg.DB.GetSidequestSubquestIDs)
	if err != nil {
		return apitypes.Sidequest{}, err
	}

	response := apitypes.Sidequest{
		ID:           sidequest.ID,
		Name:         sidequest.Name,
		UntypedQuest: handlerinput.IDToQuestAPIResource(E.Quests, sidequest.Quest.ID),
		Availability: sidequest.Availability,
		IsRepeatable: sidequest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(sidequest.Completion, convertQuestCompletion),
		Subquests:    subquests,
	}

	return response, nil
}

func retrieveSidequests(r *http.Request, i handlerinput.HandlerInput[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{})
}
