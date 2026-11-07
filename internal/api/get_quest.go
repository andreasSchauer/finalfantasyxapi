package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getQuest(r *http.Request, i handlerinput.HandlerInput[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource], id int32) (apitypes.Quest, error) {
	quest, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Quest{}, err
	}

	response := apitypes.Quest{
		ID:           quest.ID,
		Name:         quest.Name,
		Type:         quest.Type,
		TypedQuest:   questToQuestAPIResource(quest),
		Availability: quest.Availability,
		IsRepeatable: quest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(quest.Completion, convertQuestCompletion),
	}

	return response, nil
}

func retrieveQuests(r *http.Request, i handlerinput.HandlerInput[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnum(r, enums.T.QuestType, ids, ts.QpnType, config.Cfg.DB.GetQuestIDsByType),
		i.QueryBool(r, ids, ts.QpnRepeatable, config.Cfg.DB.GetQuestIDsByRepeatable),
	})
}
