package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getQuest(r *http.Request, i handlerInput[seedtypes.Quest, Quest, QuestAPIResource], id int32) (Quest, error) {
	quest, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Quest{}, err
	}

	response := Quest{
		ID:           quest.ID,
		Name:         quest.Name,
		Type:         quest.Type,
		TypedQuest:   questToQuestAPIResource(quest),
		Availability: quest.Availability,
		IsRepeatable: quest.IsRepeatable,
		Completion:   convertObjPtr(quest.Completion, convertQuestCompletion),
	}

	return response, nil
}

func retrieveQuests(r *http.Request, i handlerInput[seedtypes.Quest, Quest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnum(r, enums.T.QuestType, ids, ts.QpnType, config.Cfg.DB.GetQuestIDsByType),
		i.queryBool(r, ids, ts.QpnRepeatable, config.Cfg.DB.GetQuestIDsByRepeatable),
	})
}
