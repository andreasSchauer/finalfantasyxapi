package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getQuest(r *http.Request, i handlerinput.Main[seedtypes.Quest, apitypes.Quest, apiresource.Quest], id int32) (apitypes.Quest, error) {
	quest, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Quest{}, err
	}

	response := apitypes.Quest{
		ID:           quest.ID,
		Name:         quest.Name,
		Type:         quest.Type,
		TypedQuest:   apiconv.QuestToQuestAPIResource(quest),
		Availability: quest.Availability,
		IsRepeatable: quest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(quest.Completion, apiconv.ConvertQuestCompletion),
	}

	return response, nil
}

func retrieveQuests(r *http.Request, i handlerinput.Main[seedtypes.Quest, apitypes.Quest, apiresource.Quest]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnum(r, enums.T.QuestType, ids, ts.QpnType, config.Cfg.DB.GetQuestIDsByType),
		i.QueryBool(r, ids, ts.QpnRepeatable, config.Cfg.DB.GetQuestIDsByRepeatable),
	})
}
