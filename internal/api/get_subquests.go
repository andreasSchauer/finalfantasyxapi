package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getSubquest(r *http.Request, i handlerinput.HandlerInput[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource], id int32) (apitypes.Subquest, error) {
	subquest, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Subquest{}, err
	}

	response := apitypes.Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		UntypedQuest:    handlerinput.IDToQuestAPIResource(E.Quests, subquest.Quest.ID),
		ParentSidequest: handlerinput.IDToQuestAPIResource(E.Sidequests, subquest.SidequestID),
		Availability:    subquest.Availability,
		IsRepeatable:    subquest.IsRepeatable,
		Completion:      convertQuestCompletion(*subquest.Completion),
	}

	if response.ParentSidequest.Sidequest == "monster arena" {
		response.ArenaCreation = handlerinput.NamePtrToNamedAPIResPtr(E.ArenaCreations, &subquest.Name, nil)
	}

	return response, nil
}

func retrieveSubquests(r *http.Request, i handlerinput.HandlerInput[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryBool(r, ids, ts.QpnRepeatable, config.Cfg.DB.GetSubquestIDsByRepeatable),
	})
}
